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

// FakeConfigAggregateAuthorizations implements ConfigAggregateAuthorizationInterface
type FakeConfigAggregateAuthorizations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var configaggregateauthorizationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "configaggregateauthorizations"}

var configaggregateauthorizationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ConfigAggregateAuthorization"}

// Get takes name of the configAggregateAuthorization, and returns the corresponding configAggregateAuthorization object, and an error if there is any.
func (c *FakeConfigAggregateAuthorizations) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configaggregateauthorizationsResource, c.ns, name), &v1alpha1.ConfigAggregateAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), err
}

// List takes label and field selectors, and returns the list of ConfigAggregateAuthorizations that match those selectors.
func (c *FakeConfigAggregateAuthorizations) List(opts v1.ListOptions) (result *v1alpha1.ConfigAggregateAuthorizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configaggregateauthorizationsResource, configaggregateauthorizationsKind, c.ns, opts), &v1alpha1.ConfigAggregateAuthorizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigAggregateAuthorizationList{ListMeta: obj.(*v1alpha1.ConfigAggregateAuthorizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigAggregateAuthorizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configAggregateAuthorizations.
func (c *FakeConfigAggregateAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configaggregateauthorizationsResource, c.ns, opts))

}

// Create takes the representation of a configAggregateAuthorization and creates it.  Returns the server's representation of the configAggregateAuthorization, and an error, if there is any.
func (c *FakeConfigAggregateAuthorizations) Create(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configaggregateauthorizationsResource, c.ns, configAggregateAuthorization), &v1alpha1.ConfigAggregateAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), err
}

// Update takes the representation of a configAggregateAuthorization and updates it. Returns the server's representation of the configAggregateAuthorization, and an error, if there is any.
func (c *FakeConfigAggregateAuthorizations) Update(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configaggregateauthorizationsResource, c.ns, configAggregateAuthorization), &v1alpha1.ConfigAggregateAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigAggregateAuthorizations) UpdateStatus(configAggregateAuthorization *v1alpha1.ConfigAggregateAuthorization) (*v1alpha1.ConfigAggregateAuthorization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configaggregateauthorizationsResource, "status", c.ns, configAggregateAuthorization), &v1alpha1.ConfigAggregateAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), err
}

// Delete takes name of the configAggregateAuthorization and deletes it. Returns an error if one occurs.
func (c *FakeConfigAggregateAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configaggregateauthorizationsResource, c.ns, name), &v1alpha1.ConfigAggregateAuthorization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigAggregateAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configaggregateauthorizationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigAggregateAuthorizationList{})
	return err
}

// Patch applies the patch and returns the patched configAggregateAuthorization.
func (c *FakeConfigAggregateAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigAggregateAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configaggregateauthorizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigAggregateAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigAggregateAuthorization), err
}
