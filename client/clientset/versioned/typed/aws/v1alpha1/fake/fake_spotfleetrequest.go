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

// FakeSpotFleetRequests implements SpotFleetRequestInterface
type FakeSpotFleetRequests struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var spotfleetrequestsResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "spotfleetrequests"}

var spotfleetrequestsKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "SpotFleetRequest"}

// Get takes name of the spotFleetRequest, and returns the corresponding spotFleetRequest object, and an error if there is any.
func (c *FakeSpotFleetRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.SpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spotfleetrequestsResource, c.ns, name), &v1alpha1.SpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotFleetRequest), err
}

// List takes label and field selectors, and returns the list of SpotFleetRequests that match those selectors.
func (c *FakeSpotFleetRequests) List(opts v1.ListOptions) (result *v1alpha1.SpotFleetRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spotfleetrequestsResource, spotfleetrequestsKind, c.ns, opts), &v1alpha1.SpotFleetRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpotFleetRequestList{ListMeta: obj.(*v1alpha1.SpotFleetRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpotFleetRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spotFleetRequests.
func (c *FakeSpotFleetRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spotfleetrequestsResource, c.ns, opts))

}

// Create takes the representation of a spotFleetRequest and creates it.  Returns the server's representation of the spotFleetRequest, and an error, if there is any.
func (c *FakeSpotFleetRequests) Create(spotFleetRequest *v1alpha1.SpotFleetRequest) (result *v1alpha1.SpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spotfleetrequestsResource, c.ns, spotFleetRequest), &v1alpha1.SpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotFleetRequest), err
}

// Update takes the representation of a spotFleetRequest and updates it. Returns the server's representation of the spotFleetRequest, and an error, if there is any.
func (c *FakeSpotFleetRequests) Update(spotFleetRequest *v1alpha1.SpotFleetRequest) (result *v1alpha1.SpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spotfleetrequestsResource, c.ns, spotFleetRequest), &v1alpha1.SpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotFleetRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSpotFleetRequests) UpdateStatus(spotFleetRequest *v1alpha1.SpotFleetRequest) (*v1alpha1.SpotFleetRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(spotfleetrequestsResource, "status", c.ns, spotFleetRequest), &v1alpha1.SpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotFleetRequest), err
}

// Delete takes name of the spotFleetRequest and deletes it. Returns an error if one occurs.
func (c *FakeSpotFleetRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(spotfleetrequestsResource, c.ns, name), &v1alpha1.SpotFleetRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpotFleetRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spotfleetrequestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpotFleetRequestList{})
	return err
}

// Patch applies the patch and returns the patched spotFleetRequest.
func (c *FakeSpotFleetRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SpotFleetRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spotfleetrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SpotFleetRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SpotFleetRequest), err
}
