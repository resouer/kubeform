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

// FakeIamOpenidConnectProviders implements IamOpenidConnectProviderInterface
type FakeIamOpenidConnectProviders struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var iamopenidconnectprovidersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "iamopenidconnectproviders"}

var iamopenidconnectprovidersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "IamOpenidConnectProvider"}

// Get takes name of the iamOpenidConnectProvider, and returns the corresponding iamOpenidConnectProvider object, and an error if there is any.
func (c *FakeIamOpenidConnectProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.IamOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iamopenidconnectprovidersResource, c.ns, name), &v1alpha1.IamOpenidConnectProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), err
}

// List takes label and field selectors, and returns the list of IamOpenidConnectProviders that match those selectors.
func (c *FakeIamOpenidConnectProviders) List(opts v1.ListOptions) (result *v1alpha1.IamOpenidConnectProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iamopenidconnectprovidersResource, iamopenidconnectprovidersKind, c.ns, opts), &v1alpha1.IamOpenidConnectProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IamOpenidConnectProviderList{ListMeta: obj.(*v1alpha1.IamOpenidConnectProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.IamOpenidConnectProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iamOpenidConnectProviders.
func (c *FakeIamOpenidConnectProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iamopenidconnectprovidersResource, c.ns, opts))

}

// Create takes the representation of a iamOpenidConnectProvider and creates it.  Returns the server's representation of the iamOpenidConnectProvider, and an error, if there is any.
func (c *FakeIamOpenidConnectProviders) Create(iamOpenidConnectProvider *v1alpha1.IamOpenidConnectProvider) (result *v1alpha1.IamOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iamopenidconnectprovidersResource, c.ns, iamOpenidConnectProvider), &v1alpha1.IamOpenidConnectProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), err
}

// Update takes the representation of a iamOpenidConnectProvider and updates it. Returns the server's representation of the iamOpenidConnectProvider, and an error, if there is any.
func (c *FakeIamOpenidConnectProviders) Update(iamOpenidConnectProvider *v1alpha1.IamOpenidConnectProvider) (result *v1alpha1.IamOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iamopenidconnectprovidersResource, c.ns, iamOpenidConnectProvider), &v1alpha1.IamOpenidConnectProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIamOpenidConnectProviders) UpdateStatus(iamOpenidConnectProvider *v1alpha1.IamOpenidConnectProvider) (*v1alpha1.IamOpenidConnectProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iamopenidconnectprovidersResource, "status", c.ns, iamOpenidConnectProvider), &v1alpha1.IamOpenidConnectProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), err
}

// Delete takes name of the iamOpenidConnectProvider and deletes it. Returns an error if one occurs.
func (c *FakeIamOpenidConnectProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iamopenidconnectprovidersResource, c.ns, name), &v1alpha1.IamOpenidConnectProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIamOpenidConnectProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iamopenidconnectprovidersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IamOpenidConnectProviderList{})
	return err
}

// Patch applies the patch and returns the patched iamOpenidConnectProvider.
func (c *FakeIamOpenidConnectProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamOpenidConnectProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iamopenidconnectprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.IamOpenidConnectProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IamOpenidConnectProvider), err
}
