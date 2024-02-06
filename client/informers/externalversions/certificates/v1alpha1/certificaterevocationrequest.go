/*


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

package v1alpha1

import (
	"context"
	time "time"

	certificatesv1alpha1 "github.com/openshift/hypershift/api/certificates/v1alpha1"
	clientset "github.com/openshift/hypershift/client/clientset/clientset"
	internalinterfaces "github.com/openshift/hypershift/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openshift/hypershift/client/listers/certificates/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CertificateRevocationRequestInformer provides access to a shared informer and lister for
// CertificateRevocationRequests.
type CertificateRevocationRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CertificateRevocationRequestLister
}

type certificateRevocationRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCertificateRevocationRequestInformer constructs a new informer for CertificateRevocationRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateRevocationRequestInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateRevocationRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateRevocationRequestInformer constructs a new informer for CertificateRevocationRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateRevocationRequestInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1alpha1().CertificateRevocationRequests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CertificatesV1alpha1().CertificateRevocationRequests(namespace).Watch(context.TODO(), options)
			},
		},
		&certificatesv1alpha1.CertificateRevocationRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateRevocationRequestInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateRevocationRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certificateRevocationRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&certificatesv1alpha1.CertificateRevocationRequest{}, f.defaultInformer)
}

func (f *certificateRevocationRequestInformer) Lister() v1alpha1.CertificateRevocationRequestLister {
	return v1alpha1.NewCertificateRevocationRequestLister(f.Informer().GetIndexer())
}
