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

// FakeOpsworksMysqlLayers implements OpsworksMysqlLayerInterface
type FakeOpsworksMysqlLayers struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var opsworksmysqllayersResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "opsworksmysqllayers"}

var opsworksmysqllayersKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "OpsworksMysqlLayer"}

// Get takes name of the opsworksMysqlLayer, and returns the corresponding opsworksMysqlLayer object, and an error if there is any.
func (c *FakeOpsworksMysqlLayers) Get(name string, options v1.GetOptions) (result *v1alpha1.OpsworksMysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(opsworksmysqllayersResource, c.ns, name), &v1alpha1.OpsworksMysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksMysqlLayer), err
}

// List takes label and field selectors, and returns the list of OpsworksMysqlLayers that match those selectors.
func (c *FakeOpsworksMysqlLayers) List(opts v1.ListOptions) (result *v1alpha1.OpsworksMysqlLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(opsworksmysqllayersResource, opsworksmysqllayersKind, c.ns, opts), &v1alpha1.OpsworksMysqlLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpsworksMysqlLayerList{ListMeta: obj.(*v1alpha1.OpsworksMysqlLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpsworksMysqlLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested opsworksMysqlLayers.
func (c *FakeOpsworksMysqlLayers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(opsworksmysqllayersResource, c.ns, opts))

}

// Create takes the representation of a opsworksMysqlLayer and creates it.  Returns the server's representation of the opsworksMysqlLayer, and an error, if there is any.
func (c *FakeOpsworksMysqlLayers) Create(opsworksMysqlLayer *v1alpha1.OpsworksMysqlLayer) (result *v1alpha1.OpsworksMysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(opsworksmysqllayersResource, c.ns, opsworksMysqlLayer), &v1alpha1.OpsworksMysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksMysqlLayer), err
}

// Update takes the representation of a opsworksMysqlLayer and updates it. Returns the server's representation of the opsworksMysqlLayer, and an error, if there is any.
func (c *FakeOpsworksMysqlLayers) Update(opsworksMysqlLayer *v1alpha1.OpsworksMysqlLayer) (result *v1alpha1.OpsworksMysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(opsworksmysqllayersResource, c.ns, opsworksMysqlLayer), &v1alpha1.OpsworksMysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksMysqlLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpsworksMysqlLayers) UpdateStatus(opsworksMysqlLayer *v1alpha1.OpsworksMysqlLayer) (*v1alpha1.OpsworksMysqlLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(opsworksmysqllayersResource, "status", c.ns, opsworksMysqlLayer), &v1alpha1.OpsworksMysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksMysqlLayer), err
}

// Delete takes name of the opsworksMysqlLayer and deletes it. Returns an error if one occurs.
func (c *FakeOpsworksMysqlLayers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(opsworksmysqllayersResource, c.ns, name), &v1alpha1.OpsworksMysqlLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpsworksMysqlLayers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(opsworksmysqllayersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpsworksMysqlLayerList{})
	return err
}

// Patch applies the patch and returns the patched opsworksMysqlLayer.
func (c *FakeOpsworksMysqlLayers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.OpsworksMysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(opsworksmysqllayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpsworksMysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpsworksMysqlLayer), err
}
