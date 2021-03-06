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

package v1alpha1

import (
	"time"

	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"
	scheme "kubeform.dev/kubeform/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IamGroupMembershipsGetter has a method to return a IamGroupMembershipInterface.
// A group's client should implement this interface.
type IamGroupMembershipsGetter interface {
	IamGroupMemberships(namespace string) IamGroupMembershipInterface
}

// IamGroupMembershipInterface has methods to work with IamGroupMembership resources.
type IamGroupMembershipInterface interface {
	Create(*v1alpha1.IamGroupMembership) (*v1alpha1.IamGroupMembership, error)
	Update(*v1alpha1.IamGroupMembership) (*v1alpha1.IamGroupMembership, error)
	UpdateStatus(*v1alpha1.IamGroupMembership) (*v1alpha1.IamGroupMembership, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.IamGroupMembership, error)
	List(opts v1.ListOptions) (*v1alpha1.IamGroupMembershipList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupMembership, err error)
	IamGroupMembershipExpansion
}

// iamGroupMemberships implements IamGroupMembershipInterface
type iamGroupMemberships struct {
	client rest.Interface
	ns     string
}

// newIamGroupMemberships returns a IamGroupMemberships
func newIamGroupMemberships(c *AwsV1alpha1Client, namespace string) *iamGroupMemberships {
	return &iamGroupMemberships{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iamGroupMembership, and returns the corresponding iamGroupMembership object, and an error if there is any.
func (c *iamGroupMemberships) Get(name string, options v1.GetOptions) (result *v1alpha1.IamGroupMembership, err error) {
	result = &v1alpha1.IamGroupMembership{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IamGroupMemberships that match those selectors.
func (c *iamGroupMemberships) List(opts v1.ListOptions) (result *v1alpha1.IamGroupMembershipList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IamGroupMembershipList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iamGroupMemberships.
func (c *iamGroupMemberships) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iamGroupMembership and creates it.  Returns the server's representation of the iamGroupMembership, and an error, if there is any.
func (c *iamGroupMemberships) Create(iamGroupMembership *v1alpha1.IamGroupMembership) (result *v1alpha1.IamGroupMembership, err error) {
	result = &v1alpha1.IamGroupMembership{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		Body(iamGroupMembership).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iamGroupMembership and updates it. Returns the server's representation of the iamGroupMembership, and an error, if there is any.
func (c *iamGroupMemberships) Update(iamGroupMembership *v1alpha1.IamGroupMembership) (result *v1alpha1.IamGroupMembership, err error) {
	result = &v1alpha1.IamGroupMembership{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		Name(iamGroupMembership.Name).
		Body(iamGroupMembership).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *iamGroupMemberships) UpdateStatus(iamGroupMembership *v1alpha1.IamGroupMembership) (result *v1alpha1.IamGroupMembership, err error) {
	result = &v1alpha1.IamGroupMembership{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		Name(iamGroupMembership.Name).
		SubResource("status").
		Body(iamGroupMembership).
		Do().
		Into(result)
	return
}

// Delete takes name of the iamGroupMembership and deletes it. Returns an error if one occurs.
func (c *iamGroupMemberships) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iamGroupMemberships) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iamGroupMembership.
func (c *iamGroupMemberships) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IamGroupMembership, err error) {
	result = &v1alpha1.IamGroupMembership{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iamgroupmemberships").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
