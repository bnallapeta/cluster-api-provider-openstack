/*
Copyright 2024 The Kubernetes Authors.

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

package v1alpha7

import (
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
	apiv1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
)

// OpenStackClusterTemplateLister helps list OpenStackClusterTemplates.
// All objects returned here must be treated as read-only.
type OpenStackClusterTemplateLister interface {
	// List lists all OpenStackClusterTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiv1alpha7.OpenStackClusterTemplate, err error)
	// OpenStackClusterTemplates returns an object that can list and get OpenStackClusterTemplates.
	OpenStackClusterTemplates(namespace string) OpenStackClusterTemplateNamespaceLister
	OpenStackClusterTemplateListerExpansion
}

// openStackClusterTemplateLister implements the OpenStackClusterTemplateLister interface.
type openStackClusterTemplateLister struct {
	listers.ResourceIndexer[*apiv1alpha7.OpenStackClusterTemplate]
}

// NewOpenStackClusterTemplateLister returns a new OpenStackClusterTemplateLister.
func NewOpenStackClusterTemplateLister(indexer cache.Indexer) OpenStackClusterTemplateLister {
	return &openStackClusterTemplateLister{listers.New[*apiv1alpha7.OpenStackClusterTemplate](indexer, apiv1alpha7.Resource("openstackclustertemplate"))}
}

// OpenStackClusterTemplates returns an object that can list and get OpenStackClusterTemplates.
func (s *openStackClusterTemplateLister) OpenStackClusterTemplates(namespace string) OpenStackClusterTemplateNamespaceLister {
	return openStackClusterTemplateNamespaceLister{listers.NewNamespaced[*apiv1alpha7.OpenStackClusterTemplate](s.ResourceIndexer, namespace)}
}

// OpenStackClusterTemplateNamespaceLister helps list and get OpenStackClusterTemplates.
// All objects returned here must be treated as read-only.
type OpenStackClusterTemplateNamespaceLister interface {
	// List lists all OpenStackClusterTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apiv1alpha7.OpenStackClusterTemplate, err error)
	// Get retrieves the OpenStackClusterTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apiv1alpha7.OpenStackClusterTemplate, error)
	OpenStackClusterTemplateNamespaceListerExpansion
}

// openStackClusterTemplateNamespaceLister implements the OpenStackClusterTemplateNamespaceLister
// interface.
type openStackClusterTemplateNamespaceLister struct {
	listers.ResourceIndexer[*apiv1alpha7.OpenStackClusterTemplate]
}
