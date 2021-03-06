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

// FakeMariadbFirewallRules implements MariadbFirewallRuleInterface
type FakeMariadbFirewallRules struct {
	Fake *FakeAzurermV1alpha1
	ns   string
}

var mariadbfirewallrulesResource = schema.GroupVersionResource{Group: "azurerm.kubeform.com", Version: "v1alpha1", Resource: "mariadbfirewallrules"}

var mariadbfirewallrulesKind = schema.GroupVersionKind{Group: "azurerm.kubeform.com", Version: "v1alpha1", Kind: "MariadbFirewallRule"}

// Get takes name of the mariadbFirewallRule, and returns the corresponding mariadbFirewallRule object, and an error if there is any.
func (c *FakeMariadbFirewallRules) Get(name string, options v1.GetOptions) (result *v1alpha1.MariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mariadbfirewallrulesResource, c.ns, name), &v1alpha1.MariadbFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbFirewallRule), err
}

// List takes label and field selectors, and returns the list of MariadbFirewallRules that match those selectors.
func (c *FakeMariadbFirewallRules) List(opts v1.ListOptions) (result *v1alpha1.MariadbFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mariadbfirewallrulesResource, mariadbfirewallrulesKind, c.ns, opts), &v1alpha1.MariadbFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MariadbFirewallRuleList{ListMeta: obj.(*v1alpha1.MariadbFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.MariadbFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mariadbFirewallRules.
func (c *FakeMariadbFirewallRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mariadbfirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a mariadbFirewallRule and creates it.  Returns the server's representation of the mariadbFirewallRule, and an error, if there is any.
func (c *FakeMariadbFirewallRules) Create(mariadbFirewallRule *v1alpha1.MariadbFirewallRule) (result *v1alpha1.MariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mariadbfirewallrulesResource, c.ns, mariadbFirewallRule), &v1alpha1.MariadbFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbFirewallRule), err
}

// Update takes the representation of a mariadbFirewallRule and updates it. Returns the server's representation of the mariadbFirewallRule, and an error, if there is any.
func (c *FakeMariadbFirewallRules) Update(mariadbFirewallRule *v1alpha1.MariadbFirewallRule) (result *v1alpha1.MariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mariadbfirewallrulesResource, c.ns, mariadbFirewallRule), &v1alpha1.MariadbFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMariadbFirewallRules) UpdateStatus(mariadbFirewallRule *v1alpha1.MariadbFirewallRule) (*v1alpha1.MariadbFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mariadbfirewallrulesResource, "status", c.ns, mariadbFirewallRule), &v1alpha1.MariadbFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbFirewallRule), err
}

// Delete takes name of the mariadbFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeMariadbFirewallRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mariadbfirewallrulesResource, c.ns, name), &v1alpha1.MariadbFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMariadbFirewallRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mariadbfirewallrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MariadbFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched mariadbFirewallRule.
func (c *FakeMariadbFirewallRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MariadbFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mariadbfirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MariadbFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MariadbFirewallRule), err
}
