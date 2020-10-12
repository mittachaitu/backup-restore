// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	backuprestorev1 "github.com/mayadata.io/kubera-backup-restore/pkg/apis/backuprestore/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKuberaBackups implements KuberaBackupInterface
type FakeKuberaBackups struct {
	Fake *FakeBackuprestoreV1
	ns   string
}

var kuberabackupsResource = schema.GroupVersionResource{Group: "backuprestore.kubera.io", Version: "v1", Resource: "kuberabackups"}

var kuberabackupsKind = schema.GroupVersionKind{Group: "backuprestore.kubera.io", Version: "v1", Kind: "KuberaBackup"}

// Get takes name of the kuberaBackup, and returns the corresponding kuberaBackup object, and an error if there is any.
func (c *FakeKuberaBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *backuprestorev1.KuberaBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kuberabackupsResource, c.ns, name), &backuprestorev1.KuberaBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backuprestorev1.KuberaBackup), err
}

// List takes label and field selectors, and returns the list of KuberaBackups that match those selectors.
func (c *FakeKuberaBackups) List(ctx context.Context, opts v1.ListOptions) (result *backuprestorev1.KuberaBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kuberabackupsResource, kuberabackupsKind, c.ns, opts), &backuprestorev1.KuberaBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &backuprestorev1.KuberaBackupList{ListMeta: obj.(*backuprestorev1.KuberaBackupList).ListMeta}
	for _, item := range obj.(*backuprestorev1.KuberaBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kuberaBackups.
func (c *FakeKuberaBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kuberabackupsResource, c.ns, opts))

}

// Create takes the representation of a kuberaBackup and creates it.  Returns the server's representation of the kuberaBackup, and an error, if there is any.
func (c *FakeKuberaBackups) Create(ctx context.Context, kuberaBackup *backuprestorev1.KuberaBackup, opts v1.CreateOptions) (result *backuprestorev1.KuberaBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kuberabackupsResource, c.ns, kuberaBackup), &backuprestorev1.KuberaBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backuprestorev1.KuberaBackup), err
}

// Update takes the representation of a kuberaBackup and updates it. Returns the server's representation of the kuberaBackup, and an error, if there is any.
func (c *FakeKuberaBackups) Update(ctx context.Context, kuberaBackup *backuprestorev1.KuberaBackup, opts v1.UpdateOptions) (result *backuprestorev1.KuberaBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kuberabackupsResource, c.ns, kuberaBackup), &backuprestorev1.KuberaBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backuprestorev1.KuberaBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKuberaBackups) UpdateStatus(ctx context.Context, kuberaBackup *backuprestorev1.KuberaBackup, opts v1.UpdateOptions) (*backuprestorev1.KuberaBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kuberabackupsResource, "status", c.ns, kuberaBackup), &backuprestorev1.KuberaBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backuprestorev1.KuberaBackup), err
}

// Delete takes name of the kuberaBackup and deletes it. Returns an error if one occurs.
func (c *FakeKuberaBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kuberabackupsResource, c.ns, name), &backuprestorev1.KuberaBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKuberaBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kuberabackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &backuprestorev1.KuberaBackupList{})
	return err
}

// Patch applies the patch and returns the patched kuberaBackup.
func (c *FakeKuberaBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *backuprestorev1.KuberaBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kuberabackupsResource, c.ns, name, pt, data, subresources...), &backuprestorev1.KuberaBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*backuprestorev1.KuberaBackup), err
}