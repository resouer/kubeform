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

// FakeVpcDHCPOptionsAssociations implements VpcDHCPOptionsAssociationInterface
type FakeVpcDHCPOptionsAssociations struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var vpcdhcpoptionsassociationsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "vpcdhcpoptionsassociations"}

var vpcdhcpoptionsassociationsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "VpcDHCPOptionsAssociation"}

// Get takes name of the vpcDHCPOptionsAssociation, and returns the corresponding vpcDHCPOptionsAssociation object, and an error if there is any.
func (c *FakeVpcDHCPOptionsAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.VpcDHCPOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcdhcpoptionsassociationsResource, c.ns, name), &v1alpha1.VpcDHCPOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), err
}

// List takes label and field selectors, and returns the list of VpcDHCPOptionsAssociations that match those selectors.
func (c *FakeVpcDHCPOptionsAssociations) List(opts v1.ListOptions) (result *v1alpha1.VpcDHCPOptionsAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcdhcpoptionsassociationsResource, vpcdhcpoptionsassociationsKind, c.ns, opts), &v1alpha1.VpcDHCPOptionsAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcDHCPOptionsAssociationList{ListMeta: obj.(*v1alpha1.VpcDHCPOptionsAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcDHCPOptionsAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcDHCPOptionsAssociations.
func (c *FakeVpcDHCPOptionsAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcdhcpoptionsassociationsResource, c.ns, opts))

}

// Create takes the representation of a vpcDHCPOptionsAssociation and creates it.  Returns the server's representation of the vpcDHCPOptionsAssociation, and an error, if there is any.
func (c *FakeVpcDHCPOptionsAssociations) Create(vpcDHCPOptionsAssociation *v1alpha1.VpcDHCPOptionsAssociation) (result *v1alpha1.VpcDHCPOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcdhcpoptionsassociationsResource, c.ns, vpcDHCPOptionsAssociation), &v1alpha1.VpcDHCPOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), err
}

// Update takes the representation of a vpcDHCPOptionsAssociation and updates it. Returns the server's representation of the vpcDHCPOptionsAssociation, and an error, if there is any.
func (c *FakeVpcDHCPOptionsAssociations) Update(vpcDHCPOptionsAssociation *v1alpha1.VpcDHCPOptionsAssociation) (result *v1alpha1.VpcDHCPOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcdhcpoptionsassociationsResource, c.ns, vpcDHCPOptionsAssociation), &v1alpha1.VpcDHCPOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcDHCPOptionsAssociations) UpdateStatus(vpcDHCPOptionsAssociation *v1alpha1.VpcDHCPOptionsAssociation) (*v1alpha1.VpcDHCPOptionsAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcdhcpoptionsassociationsResource, "status", c.ns, vpcDHCPOptionsAssociation), &v1alpha1.VpcDHCPOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), err
}

// Delete takes name of the vpcDHCPOptionsAssociation and deletes it. Returns an error if one occurs.
func (c *FakeVpcDHCPOptionsAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcdhcpoptionsassociationsResource, c.ns, name), &v1alpha1.VpcDHCPOptionsAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcDHCPOptionsAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcdhcpoptionsassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcDHCPOptionsAssociationList{})
	return err
}

// Patch applies the patch and returns the patched vpcDHCPOptionsAssociation.
func (c *FakeVpcDHCPOptionsAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VpcDHCPOptionsAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcdhcpoptionsassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcDHCPOptionsAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcDHCPOptionsAssociation), err
}
