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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/api/v1alpha7"
	apiv1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/api/v1alpha7"
	typedapiv1alpha7 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/clientset/clientset/typed/api/v1alpha7"
)

// fakeOpenStackClusters implements OpenStackClusterInterface
type fakeOpenStackClusters struct {
	*gentype.FakeClientWithListAndApply[*v1alpha7.OpenStackCluster, *v1alpha7.OpenStackClusterList, *apiv1alpha7.OpenStackClusterApplyConfiguration]
	Fake *FakeInfrastructureV1alpha7
}

func newFakeOpenStackClusters(fake *FakeInfrastructureV1alpha7, namespace string) typedapiv1alpha7.OpenStackClusterInterface {
	return &fakeOpenStackClusters{
		gentype.NewFakeClientWithListAndApply[*v1alpha7.OpenStackCluster, *v1alpha7.OpenStackClusterList, *apiv1alpha7.OpenStackClusterApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha7.SchemeGroupVersion.WithResource("openstackclusters"),
			v1alpha7.SchemeGroupVersion.WithKind("OpenStackCluster"),
			func() *v1alpha7.OpenStackCluster { return &v1alpha7.OpenStackCluster{} },
			func() *v1alpha7.OpenStackClusterList { return &v1alpha7.OpenStackClusterList{} },
			func(dst, src *v1alpha7.OpenStackClusterList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha7.OpenStackClusterList) []*v1alpha7.OpenStackCluster {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha7.OpenStackClusterList, items []*v1alpha7.OpenStackCluster) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
