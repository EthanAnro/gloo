/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	gloosoloiov1 "github.com/solo-io/gloo/projects/controller/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	versioned "github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1/kube/client/clientset/versioned"
	internalinterfaces "github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1/kube/client/informers/externalversions/internalinterfaces"
	v1 "github.com/solo-io/gloo/projects/controllerrollerroller/pkg/api/v1/kube/client/listers/gloo.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProxyInformer provides access to a shared informer and lister for
// Proxies.
type ProxyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProxyLister
}

type proxyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProxyInformer constructs a new informer for Proxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProxyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProxyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProxyInformer constructs a new informer for Proxy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProxyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GlooV1().Proxies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GlooV1().Proxies(namespace).Watch(context.TODO(), options)
			},
		},
		&gloosoloiov1.Proxy{},
		resyncPeriod,
		indexers,
	)
}

func (f *proxyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProxyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *proxyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gloosoloiov1.Proxy{}, f.defaultInformer)
}

func (f *proxyInformer) Lister() v1.ProxyLister {
	return v1.NewProxyLister(f.Informer().GetIndexer())
}