// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"log"
	"sort"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Compile-time assertion
	_ resources.Resource = new(KubeService)
)

func NewKubeServiceHashableResource() resources.HashableResource {
	return new(KubeService)
}

func NewKubeService(namespace, name string) *KubeService {
	kubeservice := &KubeService{}
	kubeservice.SetMetadata(&core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return kubeservice
}

func (r *KubeService) SetMetadata(meta *core.Metadata) {
	r.Metadata = meta
}

func (r *KubeService) MustHash() uint64 {
	hashVal, err := r.Hash(nil)
	if err != nil {
		log.Panicf("error while hashing: (%s) this should never happen", err)
	}
	return hashVal
}

func (r *KubeService) GroupVersionKind() schema.GroupVersionKind {
	return KubeServiceGVK
}

type KubeServiceList []*KubeService

func (list KubeServiceList) Find(namespace, name string) (*KubeService, error) {
	for _, kubeService := range list {
		if kubeService.GetMetadata().Name == name && kubeService.GetMetadata().Namespace == namespace {
			return kubeService, nil
		}
	}
	return nil, errors.Errorf("list did not find kubeService %v.%v", namespace, name)
}

func (list KubeServiceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, kubeService := range list {
		ress = append(ress, kubeService)
	}
	return ress
}

func (list KubeServiceList) Names() []string {
	var names []string
	for _, kubeService := range list {
		names = append(names, kubeService.GetMetadata().Name)
	}
	return names
}

func (list KubeServiceList) NamespacesDotNames() []string {
	var names []string
	for _, kubeService := range list {
		names = append(names, kubeService.GetMetadata().Namespace+"."+kubeService.GetMetadata().Name)
	}
	return names
}

func (list KubeServiceList) Sort() KubeServiceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list KubeServiceList) Clone() KubeServiceList {
	var kubeServiceList KubeServiceList
	for _, kubeService := range list {
		kubeServiceList = append(kubeServiceList, resources.Clone(kubeService).(*KubeService))
	}
	return kubeServiceList
}

func (list KubeServiceList) Each(f func(element *KubeService)) {
	for _, kubeService := range list {
		f(kubeService)
	}
}

func (list KubeServiceList) EachResource(f func(element resources.Resource)) {
	for _, kubeService := range list {
		f(kubeService)
	}
}

func (list KubeServiceList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *KubeService) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

// Kubernetes Adapter for KubeService

func (o *KubeService) GetObjectKind() schema.ObjectKind {
	t := KubeServiceCrd.TypeMeta()
	return &t
}

func (o *KubeService) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*KubeService)
}

func (o *KubeService) DeepCopyInto(out *KubeService) {
	clone := resources.Clone(o).(*KubeService)
	*out = *clone
}

var (
	KubeServiceCrd = crd.NewCrd(
		"services",
		KubeServiceGVK.Group,
		KubeServiceGVK.Version,
		KubeServiceGVK.Kind,
		"sv",
		false,
		&KubeService{})
)

var (
	KubeServiceGVK = schema.GroupVersionKind{
		Version: "v1",
		Group:   "ingress.solo.io",
		Kind:    "KubeService",
	}
)
