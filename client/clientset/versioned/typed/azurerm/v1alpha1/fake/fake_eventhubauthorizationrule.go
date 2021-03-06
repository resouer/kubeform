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
	v1alpha1 "kubeform.dev/kubeform/apis/azurerm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventhubAuthorizationRules implements EventhubAuthorizationRuleInterface
type FakeEventhubAuthorizationRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var eventhubauthorizationrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "eventhubauthorizationrules"}

var eventhubauthorizationrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "EventhubAuthorizationRule"}

// Get takes name of the eventhubAuthorizationRule, and returns the corresponding eventhubAuthorizationRule object, and an error if there is any.
func (c *FakeEventhubAuthorizationRules) Get(name string, options v1.GetOptions) (result *v1alpha1.EventhubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventhubauthorizationrulesResource, c.ns, name), &v1alpha1.EventhubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubAuthorizationRule), err
}

// List takes label and field selectors, and returns the list of EventhubAuthorizationRules that match those selectors.
func (c *FakeEventhubAuthorizationRules) List(opts v1.ListOptions) (result *v1alpha1.EventhubAuthorizationRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventhubauthorizationrulesResource, eventhubauthorizationrulesKind, c.ns, opts), &v1alpha1.EventhubAuthorizationRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventhubAuthorizationRuleList{ListMeta: obj.(*v1alpha1.EventhubAuthorizationRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventhubAuthorizationRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventhubAuthorizationRules.
func (c *FakeEventhubAuthorizationRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventhubauthorizationrulesResource, c.ns, opts))

}

// Create takes the representation of a eventhubAuthorizationRule and creates it.  Returns the server's representation of the eventhubAuthorizationRule, and an error, if there is any.
func (c *FakeEventhubAuthorizationRules) Create(eventhubAuthorizationRule *v1alpha1.EventhubAuthorizationRule) (result *v1alpha1.EventhubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventhubauthorizationrulesResource, c.ns, eventhubAuthorizationRule), &v1alpha1.EventhubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubAuthorizationRule), err
}

// Update takes the representation of a eventhubAuthorizationRule and updates it. Returns the server's representation of the eventhubAuthorizationRule, and an error, if there is any.
func (c *FakeEventhubAuthorizationRules) Update(eventhubAuthorizationRule *v1alpha1.EventhubAuthorizationRule) (result *v1alpha1.EventhubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventhubauthorizationrulesResource, c.ns, eventhubAuthorizationRule), &v1alpha1.EventhubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubAuthorizationRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventhubAuthorizationRules) UpdateStatus(eventhubAuthorizationRule *v1alpha1.EventhubAuthorizationRule) (*v1alpha1.EventhubAuthorizationRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventhubauthorizationrulesResource, "status", c.ns, eventhubAuthorizationRule), &v1alpha1.EventhubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubAuthorizationRule), err
}

// Delete takes name of the eventhubAuthorizationRule and deletes it. Returns an error if one occurs.
func (c *FakeEventhubAuthorizationRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventhubauthorizationrulesResource, c.ns, name), &v1alpha1.EventhubAuthorizationRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventhubAuthorizationRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventhubauthorizationrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventhubAuthorizationRuleList{})
	return err
}

// Patch applies the patch and returns the patched eventhubAuthorizationRule.
func (c *FakeEventhubAuthorizationRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventhubAuthorizationRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventhubauthorizationrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventhubAuthorizationRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventhubAuthorizationRule), err
}
