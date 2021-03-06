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

// EcsTaskDefinitionLister helps list EcsTaskDefinitions.
type EcsTaskDefinitionLister interface {
	// List lists all EcsTaskDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EcsTaskDefinition, err error)
	// EcsTaskDefinitions returns an object that can list and get EcsTaskDefinitions.
	EcsTaskDefinitions(namespace string) EcsTaskDefinitionNamespaceLister
	EcsTaskDefinitionListerExpansion
}

// ecsTaskDefinitionLister implements the EcsTaskDefinitionLister interface.
type ecsTaskDefinitionLister struct {
	indexer cache.Indexer
}

// NewEcsTaskDefinitionLister returns a new EcsTaskDefinitionLister.
func NewEcsTaskDefinitionLister(indexer cache.Indexer) EcsTaskDefinitionLister {
	return &ecsTaskDefinitionLister{indexer: indexer}
}

// List lists all EcsTaskDefinitions in the indexer.
func (s *ecsTaskDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.EcsTaskDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcsTaskDefinition))
	})
	return ret, err
}

// EcsTaskDefinitions returns an object that can list and get EcsTaskDefinitions.
func (s *ecsTaskDefinitionLister) EcsTaskDefinitions(namespace string) EcsTaskDefinitionNamespaceLister {
	return ecsTaskDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EcsTaskDefinitionNamespaceLister helps list and get EcsTaskDefinitions.
type EcsTaskDefinitionNamespaceLister interface {
	// List lists all EcsTaskDefinitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EcsTaskDefinition, err error)
	// Get retrieves the EcsTaskDefinition from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EcsTaskDefinition, error)
	EcsTaskDefinitionNamespaceListerExpansion
}

// ecsTaskDefinitionNamespaceLister implements the EcsTaskDefinitionNamespaceLister
// interface.
type ecsTaskDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EcsTaskDefinitions in the indexer for a given namespace.
func (s ecsTaskDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EcsTaskDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EcsTaskDefinition))
	})
	return ret, err
}

// Get retrieves the EcsTaskDefinition from the indexer for a given namespace and name.
func (s ecsTaskDefinitionNamespaceLister) Get(name string) (*v1alpha1.EcsTaskDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ecstaskdefinition"), name)
	}
	return obj.(*v1alpha1.EcsTaskDefinition), nil
}
