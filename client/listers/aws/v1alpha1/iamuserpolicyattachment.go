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

// IamUserPolicyAttachmentLister helps list IamUserPolicyAttachments.
type IamUserPolicyAttachmentLister interface {
	// List lists all IamUserPolicyAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicyAttachment, err error)
	// IamUserPolicyAttachments returns an object that can list and get IamUserPolicyAttachments.
	IamUserPolicyAttachments(namespace string) IamUserPolicyAttachmentNamespaceLister
	IamUserPolicyAttachmentListerExpansion
}

// iamUserPolicyAttachmentLister implements the IamUserPolicyAttachmentLister interface.
type iamUserPolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewIamUserPolicyAttachmentLister returns a new IamUserPolicyAttachmentLister.
func NewIamUserPolicyAttachmentLister(indexer cache.Indexer) IamUserPolicyAttachmentLister {
	return &iamUserPolicyAttachmentLister{indexer: indexer}
}

// List lists all IamUserPolicyAttachments in the indexer.
func (s *iamUserPolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserPolicyAttachment))
	})
	return ret, err
}

// IamUserPolicyAttachments returns an object that can list and get IamUserPolicyAttachments.
func (s *iamUserPolicyAttachmentLister) IamUserPolicyAttachments(namespace string) IamUserPolicyAttachmentNamespaceLister {
	return iamUserPolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IamUserPolicyAttachmentNamespaceLister helps list and get IamUserPolicyAttachments.
type IamUserPolicyAttachmentNamespaceLister interface {
	// List lists all IamUserPolicyAttachments in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicyAttachment, err error)
	// Get retrieves the IamUserPolicyAttachment from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.IamUserPolicyAttachment, error)
	IamUserPolicyAttachmentNamespaceListerExpansion
}

// iamUserPolicyAttachmentNamespaceLister implements the IamUserPolicyAttachmentNamespaceLister
// interface.
type iamUserPolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IamUserPolicyAttachments in the indexer for a given namespace.
func (s iamUserPolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IamUserPolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IamUserPolicyAttachment))
	})
	return ret, err
}

// Get retrieves the IamUserPolicyAttachment from the indexer for a given namespace and name.
func (s iamUserPolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.IamUserPolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("iamuserpolicyattachment"), name)
	}
	return obj.(*v1alpha1.IamUserPolicyAttachment), nil
}
