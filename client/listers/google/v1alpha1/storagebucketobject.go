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
	v1alpha1 "kubeform.dev/kubeform/apis/google/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageBucketObjectLister helps list StorageBucketObjects.
type StorageBucketObjectLister interface {
	// List lists all StorageBucketObjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBucketObject, err error)
	// StorageBucketObjects returns an object that can list and get StorageBucketObjects.
	StorageBucketObjects(namespace string) StorageBucketObjectNamespaceLister
	StorageBucketObjectListerExpansion
}

// storageBucketObjectLister implements the StorageBucketObjectLister interface.
type storageBucketObjectLister struct {
	indexer cache.Indexer
}

// NewStorageBucketObjectLister returns a new StorageBucketObjectLister.
func NewStorageBucketObjectLister(indexer cache.Indexer) StorageBucketObjectLister {
	return &storageBucketObjectLister{indexer: indexer}
}

// List lists all StorageBucketObjects in the indexer.
func (s *storageBucketObjectLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBucketObject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBucketObject))
	})
	return ret, err
}

// StorageBucketObjects returns an object that can list and get StorageBucketObjects.
func (s *storageBucketObjectLister) StorageBucketObjects(namespace string) StorageBucketObjectNamespaceLister {
	return storageBucketObjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageBucketObjectNamespaceLister helps list and get StorageBucketObjects.
type StorageBucketObjectNamespaceLister interface {
	// List lists all StorageBucketObjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.StorageBucketObject, err error)
	// Get retrieves the StorageBucketObject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.StorageBucketObject, error)
	StorageBucketObjectNamespaceListerExpansion
}

// storageBucketObjectNamespaceLister implements the StorageBucketObjectNamespaceLister
// interface.
type storageBucketObjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageBucketObjects in the indexer for a given namespace.
func (s storageBucketObjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageBucketObject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageBucketObject))
	})
	return ret, err
}

// Get retrieves the StorageBucketObject from the indexer for a given namespace and name.
func (s storageBucketObjectNamespaceLister) Get(name string) (*v1alpha1.StorageBucketObject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagebucketobject"), name)
	}
	return obj.(*v1alpha1.StorageBucketObject), nil
}
