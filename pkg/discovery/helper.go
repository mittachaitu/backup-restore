package discovery

import (
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/restmapper"
	klog "k8s.io/klog/v2"
)

// Helper contains functions that one needs to interact with API server
type Helper interface {
	// GetResources returns the current set of resources retrived
	// from api server
	GetResources() []*metav1.APIResourceList
	// GetNamespaceScopedAPIResourceList returns list of resources
	// which are namespaced scope
	GetNamespaceScopedAPIResourceList() []*metav1.APIResourceList
	// ServerVersionInfo returns the k8s version info
	ServerVersionInfo() *version.Info
	// ResourceFor takes a partially-resolved resource and
	// returns fully qualified GroupVersionResource and APIResource
	ResourcesFor(schema.GroupVersionResource) (schema.GroupVersionResource, metav1.APIResource, error)
	// KindFor taked a partially-resolved kind and
	// returns fully qualified GroupVersionKind and APIResource
	KindFor(schema.GroupVersionKind) (schema.GroupVersionResource, metav1.APIResource, error)
}

type helper struct {
	discoveryClient          discovery.DiscoveryInterface
	resources                []*metav1.APIResourceList
	namespaceScopedResources []*metav1.APIResourceList
	resourcesMap             map[schema.GroupVersionResource]metav1.APIResource
	kindMap                  map[schema.GroupVersionKind]metav1.APIResource
	versionInfo              *version.Info
	// Required for shortcut mappings
	restMapper meta.RESTMapper
}

// NewHelper returns new instance of helper which can be used to query about
// resources
func NewHelper(discoveryClient discovery.DiscoveryInterface) (Helper, error) {
	h := &helper{
		discoveryClient: discoveryClient,
	}
	if err := h.LoadResources(); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *helper) LoadResources() error {

	// loadServerGroupAndResourcces returns the supported groups and resources for
	// all groups and versions
	_, serverResources, err := loadServerGroupsAndResources(h.discoveryClient)
	if err != nil {
		return err
	}

	groupResources, err := restmapper.GetAPIGroupResources(h.discoveryClient)
	if err != nil {
		return err
	}

	h.restMapper = restmapper.NewDiscoveryRESTMapper(groupResources)

	// Filter resources based on their permissions
	h.resources = discovery.FilteredBy(
		discovery.ResourcePredicateFunc(filterByVerbs), serverResources)

	h.namespaceScopedResources = discovery.FilteredBy(
		discovery.ResourcePredicateFunc(filterByNamespaceScoped), serverResources)

	h.resourcesMap = make(map[schema.GroupVersionResource]metav1.APIResource)
	h.kindMap = make(map[schema.GroupVersionKind]metav1.APIResource)

	for _, resourceGroup := range h.resources {
		gv, err := schema.ParseGroupVersion(resourceGroup.GroupVersion)
		if err != nil {
			return errors.Wrapf(err, "unable to parse GroupVersion %s", resourceGroup.GroupVersion)
		}
		for _, resource := range resourceGroup.APIResources {
			gvr := gv.WithResource(resource.Name)
			gvk := gv.WithKind(resource.Kind)
			h.resourcesMap[gvr] = resource
			h.kindMap[gvk] = resource
		}
	}

	serverVersion, err := h.discoveryClient.ServerVersion()
	if err != nil {
		return err
	}
	h.versionInfo = serverVersion

	return nil
}

func loadServerGroupsAndResources(discoveryClient discovery.DiscoveryInterface) ([]*metav1.APIGroup, []*metav1.APIResourceList, error) {
	serverGroups, serverResources, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		if discoveryErr, ok := err.(*discovery.ErrGroupDiscoveryFailed); ok {
			for groupVersion, err := range discoveryErr.Groups {
				klog.Warningf("Failed to discover group: %v error: %v", groupVersion, err)
			}
			return serverGroups, serverResources, nil
		}
	}
	return serverGroups, serverResources, nil
}

func (h *helper) GetNamespaceScopedAPIResourceList() []*metav1.APIResourceList {
	return h.namespaceScopedResources
}

func (h *helper) GetResources() []*metav1.APIResourceList {
	return h.resources
}

func (h *helper) ServerVersionInfo() *version.Info {
	return h.versionInfo
}

func (h *helper) ResourcesFor(gvr schema.GroupVersionResource) (schema.GroupVersionResource, metav1.APIResource, error) {
	fqGVR, err := h.restMapper.ResourceFor(gvr)
	if err != nil {
		return schema.GroupVersionResource{}, metav1.APIResource{}, err
	}
	if val, ok := h.resourcesMap[fqGVR]; ok {
		return fqGVR, val, nil
	}
	return schema.GroupVersionResource{}, metav1.APIResource{},
		errors.Errorf("APIResorce not found for GroupVersionResource %v", fqGVR)
}

func (h *helper) KindFor(gvk schema.GroupVersionKind) (schema.GroupVersionResource, metav1.APIResource, error) {
	if resource, ok := h.kindMap[gvk]; ok {
		return schema.GroupVersionResource{
			Group:    resource.Group,
			Version:  resource.Version,
			Resource: resource.Name,
		}, resource, nil
	}

	mapper, err := h.restMapper.RESTMapping(schema.GroupKind{Group: gvk.Group, Kind: gvk.Kind}, gvk.Version)
	if err != nil {
		return schema.GroupVersionResource{}, metav1.APIResource{}, err
	}

	if resource, ok := h.kindMap[mapper.GroupVersionKind]; ok {
		return schema.GroupVersionResource{
			Group:    resource.Group,
			Version:  resource.Version,
			Resource: resource.Name,
		}, resource, nil
	}
	return schema.GroupVersionResource{}, metav1.APIResource{},
		errors.Errorf("APIResorce not found for GroupVersionResource %v", mapper.GroupVersionKind)
}

func filterByVerbs(groupVersion string, resource *metav1.APIResource) bool {
	return discovery.SupportsAllVerbs{Verbs: []string{"get", "list", "create", "delete"}}.Match(groupVersion, resource)
}

func filterByNamespaceScoped(groupVersion string, resource *metav1.APIResource) bool {
	return resource.Namespaced
}
