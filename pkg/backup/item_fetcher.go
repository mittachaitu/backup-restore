package backup

import (
	"context"

	"github.com/mayadata.io/kubera-backup-restore/pkg/client"
	discovery "github.com/mayadata.io/kubera-backup-restore/pkg/discovery"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
)

type server struct {
	discoveryClient discovery.Helper
	dynamicFactory  client.DynamicFactory
}

func NewServer(discoveryClient discovery.Helper, dynamicFactory client.DynamicFactory) *server {
	return &server{
		discoveryClient: discoveryClient,
		dynamicFactory:  dynamicFactory,
	}
}

func (s *server) GeResourcesFromNamespace(namespace string) ([]*unstructured.UnstructuredList, error) {
	namespaceObjects := []*unstructured.UnstructuredList{}

	for _, apiResourceList := range s.discoveryClient.GetNamespaceScopedAPIResourceList() {
		gv, err := schema.ParseGroupVersion(apiResourceList.GroupVersion)
		if err != nil {
			// This case will never happen because we are fetching from api server
			klog.Errorf("Failed to parse gv %v error: %v", apiResourceList.GroupVersion, err)
			continue
		}
		for _, apiResource := range apiResourceList.APIResources {
			dClient, _ := s.dynamicFactory.ClientForGroupVersionResource(gv, apiResource, namespace, context.TODO())
			objList, err := dClient.List(metav1.ListOptions{})
			if err != nil {
				klog.Errorf("Failed to list resources of %s error: %v", apiResource.Name, err)
				continue
			}
			if len(objList.Items) != 0 {
				namespaceObjects = append(namespaceObjects, objList)
			}
		}
	}
	return namespaceObjects, nil
}

func (s *server) GetResourcesFromGVR(
	gvr schema.GroupVersionResource, namespace string) (*unstructured.UnstructuredList, error) {
	fqGVR, apiResource, err := s.discoveryClient.ResourcesFor(gvr)
	if err != nil {
		return nil, err
	}

	dClient, _ := s.dynamicFactory.ClientForGroupVersionResource(fqGVR.GroupVersion(), apiResource, namespace, context.TODO())
	objList, err := dClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return objList, nil
}

func (s *server) GetResourcesFromGVK(
	gvk schema.GroupVersionKind, namespace string) (*unstructured.UnstructuredList, error) {
	fqGVR, apiResource, err := s.discoveryClient.KindFor(gvk)
	if err != nil {
		return nil, err
	}

	dClient, _ := s.dynamicFactory.ClientForGroupVersionResource(fqGVR.GroupVersion(), apiResource, namespace, context.TODO())
	objList, err := dClient.List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return objList, nil
}
