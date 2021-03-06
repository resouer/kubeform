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

// FakeBinaryAuthorizationPolicies implements BinaryAuthorizationPolicyInterface
type FakeBinaryAuthorizationPolicies struct {
	Fake *FakeGoogleV1alpha1
	ns   string
}

var binaryauthorizationpoliciesResource = schema.GroupVersionResource{Group: "google.kubeform.com", Version: "v1alpha1", Resource: "binaryauthorizationpolicies"}

var binaryauthorizationpoliciesKind = schema.GroupVersionKind{Group: "google.kubeform.com", Version: "v1alpha1", Kind: "BinaryAuthorizationPolicy"}

// Get takes name of the binaryAuthorizationPolicy, and returns the corresponding binaryAuthorizationPolicy object, and an error if there is any.
func (c *FakeBinaryAuthorizationPolicies) Get(name string, options v1.GetOptions) (result *v1alpha1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(binaryauthorizationpoliciesResource, c.ns, name), &v1alpha1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of BinaryAuthorizationPolicies that match those selectors.
func (c *FakeBinaryAuthorizationPolicies) List(opts v1.ListOptions) (result *v1alpha1.BinaryAuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(binaryauthorizationpoliciesResource, binaryauthorizationpoliciesKind, c.ns, opts), &v1alpha1.BinaryAuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BinaryAuthorizationPolicyList{ListMeta: obj.(*v1alpha1.BinaryAuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.BinaryAuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested binaryAuthorizationPolicies.
func (c *FakeBinaryAuthorizationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(binaryauthorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a binaryAuthorizationPolicy and creates it.  Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationPolicies) Create(binaryAuthorizationPolicy *v1alpha1.BinaryAuthorizationPolicy) (result *v1alpha1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(binaryauthorizationpoliciesResource, c.ns, binaryAuthorizationPolicy), &v1alpha1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationPolicy), err
}

// Update takes the representation of a binaryAuthorizationPolicy and updates it. Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationPolicies) Update(binaryAuthorizationPolicy *v1alpha1.BinaryAuthorizationPolicy) (result *v1alpha1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(binaryauthorizationpoliciesResource, c.ns, binaryAuthorizationPolicy), &v1alpha1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBinaryAuthorizationPolicies) UpdateStatus(binaryAuthorizationPolicy *v1alpha1.BinaryAuthorizationPolicy) (*v1alpha1.BinaryAuthorizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(binaryauthorizationpoliciesResource, "status", c.ns, binaryAuthorizationPolicy), &v1alpha1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationPolicy), err
}

// Delete takes name of the binaryAuthorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeBinaryAuthorizationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(binaryauthorizationpoliciesResource, c.ns, name), &v1alpha1.BinaryAuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBinaryAuthorizationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(binaryauthorizationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BinaryAuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched binaryAuthorizationPolicy.
func (c *FakeBinaryAuthorizationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(binaryauthorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BinaryAuthorizationPolicy), err
}
