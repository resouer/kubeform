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

// GuarddutyIpsetsGetter has a method to return a GuarddutyIpsetInterface.
// A group's client should implement this interface.
type GuarddutyIpsetsGetter interface {
	GuarddutyIpsets(namespace string) GuarddutyIpsetInterface
}

// GuarddutyIpsetInterface has methods to work with GuarddutyIpset resources.
type GuarddutyIpsetInterface interface {
	Create(*v1alpha1.GuarddutyIpset) (*v1alpha1.GuarddutyIpset, error)
	Update(*v1alpha1.GuarddutyIpset) (*v1alpha1.GuarddutyIpset, error)
	UpdateStatus(*v1alpha1.GuarddutyIpset) (*v1alpha1.GuarddutyIpset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GuarddutyIpset, error)
	List(opts v1.ListOptions) (*v1alpha1.GuarddutyIpsetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyIpset, err error)
	GuarddutyIpsetExpansion
}

// guarddutyIpsets implements GuarddutyIpsetInterface
type guarddutyIpsets struct {
	client rest.Interface
	ns     string
}

// newGuarddutyIpsets returns a GuarddutyIpsets
func newGuarddutyIpsets(c *AwsV1alpha1Client, namespace string) *guarddutyIpsets {
	return &guarddutyIpsets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the guarddutyIpset, and returns the corresponding guarddutyIpset object, and an error if there is any.
func (c *guarddutyIpsets) Get(name string, options v1.GetOptions) (result *v1alpha1.GuarddutyIpset, err error) {
	result = &v1alpha1.GuarddutyIpset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GuarddutyIpsets that match those selectors.
func (c *guarddutyIpsets) List(opts v1.ListOptions) (result *v1alpha1.GuarddutyIpsetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GuarddutyIpsetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested guarddutyIpsets.
func (c *guarddutyIpsets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a guarddutyIpset and creates it.  Returns the server's representation of the guarddutyIpset, and an error, if there is any.
func (c *guarddutyIpsets) Create(guarddutyIpset *v1alpha1.GuarddutyIpset) (result *v1alpha1.GuarddutyIpset, err error) {
	result = &v1alpha1.GuarddutyIpset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		Body(guarddutyIpset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a guarddutyIpset and updates it. Returns the server's representation of the guarddutyIpset, and an error, if there is any.
func (c *guarddutyIpsets) Update(guarddutyIpset *v1alpha1.GuarddutyIpset) (result *v1alpha1.GuarddutyIpset, err error) {
	result = &v1alpha1.GuarddutyIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		Name(guarddutyIpset.Name).
		Body(guarddutyIpset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *guarddutyIpsets) UpdateStatus(guarddutyIpset *v1alpha1.GuarddutyIpset) (result *v1alpha1.GuarddutyIpset, err error) {
	result = &v1alpha1.GuarddutyIpset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		Name(guarddutyIpset.Name).
		SubResource("status").
		Body(guarddutyIpset).
		Do().
		Into(result)
	return
}

// Delete takes name of the guarddutyIpset and deletes it. Returns an error if one occurs.
func (c *guarddutyIpsets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *guarddutyIpsets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("guarddutyipsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched guarddutyIpset.
func (c *guarddutyIpsets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GuarddutyIpset, err error) {
	result = &v1alpha1.GuarddutyIpset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("guarddutyipsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
