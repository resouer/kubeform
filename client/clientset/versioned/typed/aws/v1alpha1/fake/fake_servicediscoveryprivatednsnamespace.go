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

// FakeServiceDiscoveryPrivateDNSNamespaces implements ServiceDiscoveryPrivateDNSNamespaceInterface
type FakeServiceDiscoveryPrivateDNSNamespaces struct {
	Fake *FakeAwsV1alpha1
	ns   string
}

var servicediscoveryprivatednsnamespacesResource = schema.GroupVersionResource{Group: "aws.kubeform.com", Version: "v1alpha1", Resource: "servicediscoveryprivatednsnamespaces"}

var servicediscoveryprivatednsnamespacesKind = schema.GroupVersionKind{Group: "aws.kubeform.com", Version: "v1alpha1", Kind: "ServiceDiscoveryPrivateDNSNamespace"}

// Get takes name of the serviceDiscoveryPrivateDNSNamespace, and returns the corresponding serviceDiscoveryPrivateDNSNamespace object, and an error if there is any.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicediscoveryprivatednsnamespacesResource, c.ns, name), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace), err
}

// List takes label and field selectors, and returns the list of ServiceDiscoveryPrivateDNSNamespaces that match those selectors.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) List(opts v1.ListOptions) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicediscoveryprivatednsnamespacesResource, servicediscoveryprivatednsnamespacesKind, c.ns, opts), &v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList{ListMeta: obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceDiscoveryPrivateDNSNamespaces.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicediscoveryprivatednsnamespacesResource, c.ns, opts))

}

// Create takes the representation of a serviceDiscoveryPrivateDNSNamespace and creates it.  Returns the server's representation of the serviceDiscoveryPrivateDNSNamespace, and an error, if there is any.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Create(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicediscoveryprivatednsnamespacesResource, c.ns, serviceDiscoveryPrivateDNSNamespace), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace), err
}

// Update takes the representation of a serviceDiscoveryPrivateDNSNamespace and updates it. Returns the server's representation of the serviceDiscoveryPrivateDNSNamespace, and an error, if there is any.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Update(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicediscoveryprivatednsnamespacesResource, c.ns, serviceDiscoveryPrivateDNSNamespace), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) UpdateStatus(serviceDiscoveryPrivateDNSNamespace *v1alpha1.ServiceDiscoveryPrivateDNSNamespace) (*v1alpha1.ServiceDiscoveryPrivateDNSNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicediscoveryprivatednsnamespacesResource, "status", c.ns, serviceDiscoveryPrivateDNSNamespace), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace), err
}

// Delete takes name of the serviceDiscoveryPrivateDNSNamespace and deletes it. Returns an error if one occurs.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicediscoveryprivatednsnamespacesResource, c.ns, name), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicediscoveryprivatednsnamespacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceDiscoveryPrivateDNSNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched serviceDiscoveryPrivateDNSNamespace.
func (c *FakeServiceDiscoveryPrivateDNSNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceDiscoveryPrivateDNSNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicediscoveryprivatednsnamespacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceDiscoveryPrivateDNSNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceDiscoveryPrivateDNSNamespace), err
}
