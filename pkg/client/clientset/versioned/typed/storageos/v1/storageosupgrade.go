/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/storageos/cluster-operator/pkg/apis/storageos/v1"
	scheme "github.com/storageos/cluster-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StorageOSUpgradesGetter has a method to return a StorageOSUpgradeInterface.
// A group's client should implement this interface.
type StorageOSUpgradesGetter interface {
	StorageOSUpgrades(namespace string) StorageOSUpgradeInterface
}

// StorageOSUpgradeInterface has methods to work with StorageOSUpgrade resources.
type StorageOSUpgradeInterface interface {
	Create(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.CreateOptions) (*v1.StorageOSUpgrade, error)
	Update(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.UpdateOptions) (*v1.StorageOSUpgrade, error)
	UpdateStatus(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.UpdateOptions) (*v1.StorageOSUpgrade, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.StorageOSUpgrade, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.StorageOSUpgradeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StorageOSUpgrade, err error)
	StorageOSUpgradeExpansion
}

// storageOSUpgrades implements StorageOSUpgradeInterface
type storageOSUpgrades struct {
	client rest.Interface
	ns     string
}

// newStorageOSUpgrades returns a StorageOSUpgrades
func newStorageOSUpgrades(c *StorageosV1Client, namespace string) *storageOSUpgrades {
	return &storageOSUpgrades{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the storageOSUpgrade, and returns the corresponding storageOSUpgrade object, and an error if there is any.
func (c *storageOSUpgrades) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.StorageOSUpgrade, err error) {
	result = &v1.StorageOSUpgrade{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageosupgrades").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of StorageOSUpgrades that match those selectors.
func (c *storageOSUpgrades) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StorageOSUpgradeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.StorageOSUpgradeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("storageosupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested storageOSUpgrades.
func (c *storageOSUpgrades) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("storageosupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a storageOSUpgrade and creates it.  Returns the server's representation of the storageOSUpgrade, and an error, if there is any.
func (c *storageOSUpgrades) Create(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.CreateOptions) (result *v1.StorageOSUpgrade, err error) {
	result = &v1.StorageOSUpgrade{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("storageosupgrades").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageOSUpgrade).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a storageOSUpgrade and updates it. Returns the server's representation of the storageOSUpgrade, and an error, if there is any.
func (c *storageOSUpgrades) Update(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.UpdateOptions) (result *v1.StorageOSUpgrade, err error) {
	result = &v1.StorageOSUpgrade{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageosupgrades").
		Name(storageOSUpgrade.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageOSUpgrade).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *storageOSUpgrades) UpdateStatus(ctx context.Context, storageOSUpgrade *v1.StorageOSUpgrade, opts metav1.UpdateOptions) (result *v1.StorageOSUpgrade, err error) {
	result = &v1.StorageOSUpgrade{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("storageosupgrades").
		Name(storageOSUpgrade.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(storageOSUpgrade).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the storageOSUpgrade and deletes it. Returns an error if one occurs.
func (c *storageOSUpgrades) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageosupgrades").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *storageOSUpgrades) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("storageosupgrades").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched storageOSUpgrade.
func (c *storageOSUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.StorageOSUpgrade, err error) {
	result = &v1.StorageOSUpgrade{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("storageosupgrades").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
