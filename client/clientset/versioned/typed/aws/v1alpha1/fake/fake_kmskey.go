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
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKmsKeys implements KmsKeyInterface
type FakeKmsKeys struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var kmskeysResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "kmskeys"}

var kmskeysKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "KmsKey"}

// Get takes name of the kmsKey, and returns the corresponding kmsKey object, and an error if there is any.
func (c *FakeKmsKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.KmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kmskeysResource, c.ns, name), &v1alpha1.KmsKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKey), err
}

// List takes label and field selectors, and returns the list of KmsKeys that match those selectors.
func (c *FakeKmsKeys) List(opts v1.ListOptions) (result *v1alpha1.KmsKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kmskeysResource, kmskeysKind, c.ns, opts), &v1alpha1.KmsKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KmsKeyList{ListMeta: obj.(*v1alpha1.KmsKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.KmsKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kmsKeys.
func (c *FakeKmsKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kmskeysResource, c.ns, opts))

}

// Create takes the representation of a kmsKey and creates it.  Returns the server's representation of the kmsKey, and an error, if there is any.
func (c *FakeKmsKeys) Create(kmsKey *v1alpha1.KmsKey) (result *v1alpha1.KmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kmskeysResource, c.ns, kmsKey), &v1alpha1.KmsKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKey), err
}

// Update takes the representation of a kmsKey and updates it. Returns the server's representation of the kmsKey, and an error, if there is any.
func (c *FakeKmsKeys) Update(kmsKey *v1alpha1.KmsKey) (result *v1alpha1.KmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kmskeysResource, c.ns, kmsKey), &v1alpha1.KmsKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKmsKeys) UpdateStatus(kmsKey *v1alpha1.KmsKey) (*v1alpha1.KmsKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kmskeysResource, "status", c.ns, kmsKey), &v1alpha1.KmsKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKey), err
}

// Delete takes name of the kmsKey and deletes it. Returns an error if one occurs.
func (c *FakeKmsKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kmskeysResource, c.ns, name), &v1alpha1.KmsKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKmsKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kmskeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.KmsKeyList{})
	return err
}

// Patch applies the patch and returns the patched kmsKey.
func (c *FakeKmsKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KmsKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kmskeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.KmsKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KmsKey), err
}
