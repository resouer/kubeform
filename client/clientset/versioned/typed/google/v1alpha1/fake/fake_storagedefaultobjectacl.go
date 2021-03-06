/*
Copyright The Kubeform Authors.

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

package fake

import (
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStorageDefaultObjectACLs implements StorageDefaultObjectACLInterface
type FakeStorageDefaultObjectACLs struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var storagedefaultobjectaclsResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "storagedefaultobjectacls"}

var storagedefaultobjectaclsKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "StorageDefaultObjectACL"}

// Get takes name of the storageDefaultObjectACL, and returns the corresponding storageDefaultObjectACL object, and an error if there is any.
func (c *FakeStorageDefaultObjectACLs) Get(name string, options v1.GetOptions) (result *v1alpha1.StorageDefaultObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(storagedefaultobjectaclsResource, c.ns, name), &v1alpha1.StorageDefaultObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), err
}

// List takes label and field selectors, and returns the list of StorageDefaultObjectACLs that match those selectors.
func (c *FakeStorageDefaultObjectACLs) List(opts v1.ListOptions) (result *v1alpha1.StorageDefaultObjectACLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(storagedefaultobjectaclsResource, storagedefaultobjectaclsKind, c.ns, opts), &v1alpha1.StorageDefaultObjectACLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StorageDefaultObjectACLList{ListMeta: obj.(*v1alpha1.StorageDefaultObjectACLList).ListMeta}
	for _, item := range obj.(*v1alpha1.StorageDefaultObjectACLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested storageDefaultObjectACLs.
func (c *FakeStorageDefaultObjectACLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(storagedefaultobjectaclsResource, c.ns, opts))

}

// Create takes the representation of a storageDefaultObjectACL and creates it.  Returns the server's representation of the storageDefaultObjectACL, and an error, if there is any.
func (c *FakeStorageDefaultObjectACLs) Create(storageDefaultObjectACL *v1alpha1.StorageDefaultObjectACL) (result *v1alpha1.StorageDefaultObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(storagedefaultobjectaclsResource, c.ns, storageDefaultObjectACL), &v1alpha1.StorageDefaultObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), err
}

// Update takes the representation of a storageDefaultObjectACL and updates it. Returns the server's representation of the storageDefaultObjectACL, and an error, if there is any.
func (c *FakeStorageDefaultObjectACLs) Update(storageDefaultObjectACL *v1alpha1.StorageDefaultObjectACL) (result *v1alpha1.StorageDefaultObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(storagedefaultobjectaclsResource, c.ns, storageDefaultObjectACL), &v1alpha1.StorageDefaultObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStorageDefaultObjectACLs) UpdateStatus(storageDefaultObjectACL *v1alpha1.StorageDefaultObjectACL) (*v1alpha1.StorageDefaultObjectACL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(storagedefaultobjectaclsResource, "status", c.ns, storageDefaultObjectACL), &v1alpha1.StorageDefaultObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), err
}

// Delete takes name of the storageDefaultObjectACL and deletes it. Returns an error if one occurs.
func (c *FakeStorageDefaultObjectACLs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(storagedefaultobjectaclsResource, c.ns, name), &v1alpha1.StorageDefaultObjectACL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStorageDefaultObjectACLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(storagedefaultobjectaclsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StorageDefaultObjectACLList{})
	return err
}

// Patch applies the patch and returns the patched storageDefaultObjectACL.
func (c *FakeStorageDefaultObjectACLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.StorageDefaultObjectACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(storagedefaultobjectaclsResource, c.ns, name, pt, data, subresources...), &v1alpha1.StorageDefaultObjectACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.StorageDefaultObjectACL), err
}
