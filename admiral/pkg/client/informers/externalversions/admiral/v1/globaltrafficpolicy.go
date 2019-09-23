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
	time "time"

	admiralv1 "github.com/admiral/admiral/pkg/apis/admiral/v1"
	versioned "github.com/admiral/admiral/pkg/client/clientset/versioned"
	internalinterfaces "github.com/admiral/admiral/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/admiral/admiral/pkg/client/listers/admiral/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GlobalTrafficPolicyInformer provides access to a shared informer and lister for
// GlobalTrafficPolicies.
type GlobalTrafficPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.GlobalTrafficPolicyLister
}

type globalTrafficPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewGlobalTrafficPolicyInformer constructs a new informer for GlobalTrafficPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGlobalTrafficPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGlobalTrafficPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredGlobalTrafficPolicyInformer constructs a new informer for GlobalTrafficPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGlobalTrafficPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().GlobalTrafficPolicies(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmiralV1().GlobalTrafficPolicies(namespace).Watch(options)
			},
		},
		&admiralv1.GlobalTrafficPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *globalTrafficPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGlobalTrafficPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *globalTrafficPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admiralv1.GlobalTrafficPolicy{}, f.defaultInformer)
}

func (f *globalTrafficPolicyInformer) Lister() v1.GlobalTrafficPolicyLister {
	return v1.NewGlobalTrafficPolicyLister(f.Informer().GetIndexer())
}