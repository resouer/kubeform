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

// FakeOrganizationsPolicyAttachments implements OrganizationsPolicyAttachmentInterface
type FakeOrganizationsPolicyAttachments struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var organizationspolicyattachmentsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "organizationspolicyattachments"}

var organizationspolicyattachmentsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OrganizationsPolicyAttachment"}

// Get takes name of the organizationsPolicyAttachment, and returns the corresponding organizationsPolicyAttachment object, and an error if there is any.
func (c *FakeOrganizationsPolicyAttachments) Get(name string, options v1.GetOptions) (result *v1alpha1.OrganizationsPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationspolicyattachmentsResource, c.ns, name), &v1alpha1.OrganizationsPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicyAttachment), err
}

// List takes label and field selectors, and returns the list of OrganizationsPolicyAttachments that match those selectors.
func (c *FakeOrganizationsPolicyAttachments) List(opts v1.ListOptions) (result *v1alpha1.OrganizationsPolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationspolicyattachmentsResource, organizationspolicyattachmentsKind, c.ns, opts), &v1alpha1.OrganizationsPolicyAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationsPolicyAttachmentList{ListMeta: obj.(*v1alpha1.OrganizationsPolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationsPolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationsPolicyAttachments.
func (c *FakeOrganizationsPolicyAttachments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationspolicyattachmentsResource, c.ns, opts))

}

// Create takes the representation of a organizationsPolicyAttachment and creates it.  Returns the server's representation of the organizationsPolicyAttachment, and an error, if there is any.
func (c *FakeOrganizationsPolicyAttachments) Create(organizationsPolicyAttachment *v1alpha1.OrganizationsPolicyAttachment) (result *v1alpha1.OrganizationsPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationspolicyattachmentsResource, c.ns, organizationsPolicyAttachment), &v1alpha1.OrganizationsPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicyAttachment), err
}

// Update takes the representation of a organizationsPolicyAttachment and updates it. Returns the server's representation of the organizationsPolicyAttachment, and an error, if there is any.
func (c *FakeOrganizationsPolicyAttachments) Update(organizationsPolicyAttachment *v1alpha1.OrganizationsPolicyAttachment) (result *v1alpha1.OrganizationsPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationspolicyattachmentsResource, c.ns, organizationsPolicyAttachment), &v1alpha1.OrganizationsPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationsPolicyAttachments) UpdateStatus(organizationsPolicyAttachment *v1alpha1.OrganizationsPolicyAttachment) (*v1alpha1.OrganizationsPolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationspolicyattachmentsResource, "status", c.ns, organizationsPolicyAttachment), &v1alpha1.OrganizationsPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicyAttachment), err
}

// Delete takes name of the organizationsPolicyAttachment and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationsPolicyAttachments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationspolicyattachmentsResource, c.ns, name), &v1alpha1.OrganizationsPolicyAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationsPolicyAttachments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationspolicyattachmentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationsPolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched organizationsPolicyAttachment.
func (c *FakeOrganizationsPolicyAttachments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OrganizationsPolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationspolicyattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationsPolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationsPolicyAttachment), err
}
