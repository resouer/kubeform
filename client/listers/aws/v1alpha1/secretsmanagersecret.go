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

// SecretsmanagerSecretLister helps list SecretsmanagerSecrets.
type SecretsmanagerSecretLister interface {
	// List lists all SecretsmanagerSecrets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecretsmanagerSecret, err error)
	// SecretsmanagerSecrets returns an object that can list and get SecretsmanagerSecrets.
	SecretsmanagerSecrets(namespace string) SecretsmanagerSecretNamespaceLister
	SecretsmanagerSecretListerExpansion
}

// secretsmanagerSecretLister implements the SecretsmanagerSecretLister interface.
type secretsmanagerSecretLister struct {
	indexer cache.Indexer
}

// NewSecretsmanagerSecretLister returns a new SecretsmanagerSecretLister.
func NewSecretsmanagerSecretLister(indexer cache.Indexer) SecretsmanagerSecretLister {
	return &secretsmanagerSecretLister{indexer: indexer}
}

// List lists all SecretsmanagerSecrets in the indexer.
func (s *secretsmanagerSecretLister) List(selector labels.Selector) (ret []*v1alpha1.SecretsmanagerSecret, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretsmanagerSecret))
	})
	return ret, err
}

// SecretsmanagerSecrets returns an object that can list and get SecretsmanagerSecrets.
func (s *secretsmanagerSecretLister) SecretsmanagerSecrets(namespace string) SecretsmanagerSecretNamespaceLister {
	return secretsmanagerSecretNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretsmanagerSecretNamespaceLister helps list and get SecretsmanagerSecrets.
type SecretsmanagerSecretNamespaceLister interface {
	// List lists all SecretsmanagerSecrets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecretsmanagerSecret, err error)
	// Get retrieves the SecretsmanagerSecret from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecretsmanagerSecret, error)
	SecretsmanagerSecretNamespaceListerExpansion
}

// secretsmanagerSecretNamespaceLister implements the SecretsmanagerSecretNamespaceLister
// interface.
type secretsmanagerSecretNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretsmanagerSecrets in the indexer for a given namespace.
func (s secretsmanagerSecretNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretsmanagerSecret, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretsmanagerSecret))
	})
	return ret, err
}

// Get retrieves the SecretsmanagerSecret from the indexer for a given namespace and name.
func (s secretsmanagerSecretNamespaceLister) Get(name string) (*v1alpha1.SecretsmanagerSecret, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretsmanagersecret"), name)
	}
	return obj.(*v1alpha1.SecretsmanagerSecret), nil
}
