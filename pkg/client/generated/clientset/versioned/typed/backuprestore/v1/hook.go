// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	scheme "github.com/mayadata.io/kubera-backup-restore/pkg/client/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HooksGetter has a method to return a HookInterface.
// A group's client should implement this interface.
type HooksGetter interface {
	Hooks(namespace string) HookInterface
}

// HookInterface has methods to work with Hook resources.
type HookInterface interface {
	Create(ctx context.Context, hook *v1.Hook, opts metav1.CreateOptions) (*v1.Hook, error)
	Update(ctx context.Context, hook *v1.Hook, opts metav1.UpdateOptions) (*v1.Hook, error)
	UpdateStatus(ctx context.Context, hook *v1.Hook, opts metav1.UpdateOptions) (*v1.Hook, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Hook, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HookList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Hook, err error)
	HookExpansion
}

// hooks implements HookInterface
type hooks struct {
	client rest.Interface
	ns     string
}

// newHooks returns a Hooks
func newHooks(c *BackuprestoreV1Client, namespace string) *hooks {
	return &hooks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hook, and returns the corresponding hook object, and an error if there is any.
func (c *hooks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Hook, err error) {
	result = &v1.Hook{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hooks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Hooks that match those selectors.
func (c *hooks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HookList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HookList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hooks.
func (c *hooks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hook and creates it.  Returns the server's representation of the hook, and an error, if there is any.
func (c *hooks) Create(ctx context.Context, hook *v1.Hook, opts metav1.CreateOptions) (result *v1.Hook, err error) {
	result = &v1.Hook{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hooks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hook).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hook and updates it. Returns the server's representation of the hook, and an error, if there is any.
func (c *hooks) Update(ctx context.Context, hook *v1.Hook, opts metav1.UpdateOptions) (result *v1.Hook, err error) {
	result = &v1.Hook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hooks").
		Name(hook.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hook).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hooks) UpdateStatus(ctx context.Context, hook *v1.Hook, opts metav1.UpdateOptions) (result *v1.Hook, err error) {
	result = &v1.Hook{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hooks").
		Name(hook.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hook).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hook and deletes it. Returns an error if one occurs.
func (c *hooks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hooks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hooks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hooks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hook.
func (c *hooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Hook, err error) {
	result = &v1.Hook{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hooks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
