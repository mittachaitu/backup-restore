package client

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
)

// DynamicFactory contains methods for retrieving dynamic clients for GroupVersionResources and
// GroupVersionKinds
type DynamicFactory interface {
	// ClientForGroupVersionResource returns a Dynamic client for the given group/version
	// and resource for the given namespace.
	ClientForGroupVersionResource(gv schema.GroupVersion, resource metav1.APIResource, namespace string, context context.Context) (Dynamic, error)
}

// dynamicFactory implements DynamicFactory.
type dynamicFactory struct {
	dynamicClient dynamic.Interface
}

// NewDynamicFactory returns a new ClientPool-based dynamic factory.
func NewDynamicFactory(dynamicClient dynamic.Interface) DynamicFactory {
	return &dynamicFactory{dynamicClient: dynamicClient}
}

func (f *dynamicFactory) ClientForGroupVersionResource(gv schema.GroupVersion, resource metav1.APIResource, namespace string, ctx context.Context) (Dynamic, error) {
	return &dynamicResourceClient{
		resourceClient: f.dynamicClient.Resource(gv.WithResource(resource.Name)).Namespace(namespace),
		ctx:            ctx,
	}, nil
}

type dynamicResourceClient struct {
	resourceClient dynamic.ResourceInterface
	ctx            context.Context
}

// Dynamic contains client method for kubera to perform backup and restore
type Dynamic interface {
	// Create creates an object
	Create(obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
	// Update updates an object
	Update(obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
	// UpdateStatus updates an object status
	UpdateStatus(obj *unstructured.Unstructured) (*unstructured.Unstructured, error)
	// Fet fetch an object from etcd
	Get(name string, options metav1.GetOptions) (*unstructured.Unstructured, error)
	// List list an object from etcd
	List(opts metav1.ListOptions) (*unstructured.UnstructuredList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, data []byte) (*unstructured.Unstructured, error)
}

func (d *dynamicResourceClient) Create(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	return d.resourceClient.Create(d.ctx, obj, metav1.CreateOptions{})
}

func (d *dynamicResourceClient) Update(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	return d.resourceClient.Update(d.ctx, obj, metav1.UpdateOptions{})
}

func (d *dynamicResourceClient) UpdateStatus(obj *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	return d.resourceClient.UpdateStatus(d.ctx, obj, metav1.UpdateOptions{})
}

func (d *dynamicResourceClient) Get(name string, options metav1.GetOptions) (*unstructured.Unstructured, error) {
	return d.resourceClient.Get(d.ctx, name, options)
}

func (d *dynamicResourceClient) List(options metav1.ListOptions) (*unstructured.UnstructuredList, error) {
	return d.resourceClient.List(d.ctx, options)
}

func (d *dynamicResourceClient) Watch(options metav1.ListOptions) (watch.Interface, error) {
	return d.resourceClient.Watch(d.ctx, options)
}

func (d *dynamicResourceClient) Patch(name string, data []byte) (*unstructured.Unstructured, error) {
	return d.resourceClient.Patch(d.ctx, name, types.MergePatchType, data, metav1.PatchOptions{})
}
