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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/kubeform/apis/aws/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// Ec2TransitGatewayVpcAttachmentLister helps list Ec2TransitGatewayVpcAttachments.
type Ec2TransitGatewayVpcAttachmentLister interface {
	// List lists all Ec2TransitGatewayVpcAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGatewayVpcAttachment, err error)
	// Ec2TransitGatewayVpcAttachments returns an object that can list and get Ec2TransitGatewayVpcAttachments.
	Ec2TransitGatewayVpcAttachments(namespace string) Ec2TransitGatewayVpcAttachmentNamespaceLister
	Ec2TransitGatewayVpcAttachmentListerExpansion
}

// ec2TransitGatewayVpcAttachmentLister implements the Ec2TransitGatewayVpcAttachmentLister interface.
type ec2TransitGatewayVpcAttachmentLister struct {
	indexer cache.Indexer
}

// NewEc2TransitGatewayVpcAttachmentLister returns a new Ec2TransitGatewayVpcAttachmentLister.
func NewEc2TransitGatewayVpcAttachmentLister(indexer cache.Indexer) Ec2TransitGatewayVpcAttachmentLister {
	return &ec2TransitGatewayVpcAttachmentLister{indexer: indexer}
}

// List lists all Ec2TransitGatewayVpcAttachments in the indexer.
func (s *ec2TransitGatewayVpcAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGatewayVpcAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2TransitGatewayVpcAttachment))
	})
	return ret, err
}

// Ec2TransitGatewayVpcAttachments returns an object that can list and get Ec2TransitGatewayVpcAttachments.
func (s *ec2TransitGatewayVpcAttachmentLister) Ec2TransitGatewayVpcAttachments(namespace string) Ec2TransitGatewayVpcAttachmentNamespaceLister {
	return ec2TransitGatewayVpcAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// Ec2TransitGatewayVpcAttachmentNamespaceLister helps list and get Ec2TransitGatewayVpcAttachments.
type Ec2TransitGatewayVpcAttachmentNamespaceLister interface {
	// List lists all Ec2TransitGatewayVpcAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGatewayVpcAttachment, err error)
	// Get retrieves the Ec2TransitGatewayVpcAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Ec2TransitGatewayVpcAttachment, error)
	Ec2TransitGatewayVpcAttachmentNamespaceListerExpansion
}

// ec2TransitGatewayVpcAttachmentNamespaceLister implements the Ec2TransitGatewayVpcAttachmentNamespaceLister
// interface.
type ec2TransitGatewayVpcAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ec2TransitGatewayVpcAttachments in the indexer for a given namespace.
func (s ec2TransitGatewayVpcAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ec2TransitGatewayVpcAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ec2TransitGatewayVpcAttachment))
	})
	return ret, err
}

// Get retrieves the Ec2TransitGatewayVpcAttachment from the indexer for a given namespace and name.
func (s ec2TransitGatewayVpcAttachmentNamespaceLister) Get(name string) (*v1alpha1.Ec2TransitGatewayVpcAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ec2transitgatewayvpcattachment"), name)
	}
	return obj.(*v1alpha1.Ec2TransitGatewayVpcAttachment), nil
}
