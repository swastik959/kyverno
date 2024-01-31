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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	kyvernov1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v1"
	kyvernov1alpha2 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v1alpha2"
	kyvernov1beta1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v1beta1"
	kyvernov2 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v2"
	kyvernov2alpha1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v2alpha1"
	kyvernov2beta1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/kyverno/v2beta1"
	wgpolicyk8sv1alpha2 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/policyreport/v1alpha2"
	reportsv1 "github.com/kyverno/kyverno/pkg/client/clientset/versioned/typed/reports/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	KyvernoV1() kyvernov1.KyvernoV1Interface
	KyvernoV1alpha2() kyvernov1alpha2.KyvernoV1alpha2Interface
	KyvernoV1beta1() kyvernov1beta1.KyvernoV1beta1Interface
	KyvernoV2() kyvernov2.KyvernoV2Interface
	KyvernoV2beta1() kyvernov2beta1.KyvernoV2beta1Interface
	KyvernoV2alpha1() kyvernov2alpha1.KyvernoV2alpha1Interface
	Wgpolicyk8sV1alpha2() wgpolicyk8sv1alpha2.Wgpolicyk8sV1alpha2Interface
	ReportsV1() reportsv1.ReportsV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	kyvernoV1           *kyvernov1.KyvernoV1Client
	kyvernoV1alpha2     *kyvernov1alpha2.KyvernoV1alpha2Client
	kyvernoV1beta1      *kyvernov1beta1.KyvernoV1beta1Client
	kyvernoV2           *kyvernov2.KyvernoV2Client
	kyvernoV2beta1      *kyvernov2beta1.KyvernoV2beta1Client
	kyvernoV2alpha1     *kyvernov2alpha1.KyvernoV2alpha1Client
	wgpolicyk8sV1alpha2 *wgpolicyk8sv1alpha2.Wgpolicyk8sV1alpha2Client
	reportsV1           *reportsv1.ReportsV1Client
}

// KyvernoV1 retrieves the KyvernoV1Client
func (c *Clientset) KyvernoV1() kyvernov1.KyvernoV1Interface {
	return c.kyvernoV1
}

// KyvernoV1alpha2 retrieves the KyvernoV1alpha2Client
func (c *Clientset) KyvernoV1alpha2() kyvernov1alpha2.KyvernoV1alpha2Interface {
	return c.kyvernoV1alpha2
}

// KyvernoV1beta1 retrieves the KyvernoV1beta1Client
func (c *Clientset) KyvernoV1beta1() kyvernov1beta1.KyvernoV1beta1Interface {
	return c.kyvernoV1beta1
}

// KyvernoV2 retrieves the KyvernoV2Client
func (c *Clientset) KyvernoV2() kyvernov2.KyvernoV2Interface {
	return c.kyvernoV2
}

// KyvernoV2beta1 retrieves the KyvernoV2beta1Client
func (c *Clientset) KyvernoV2beta1() kyvernov2beta1.KyvernoV2beta1Interface {
	return c.kyvernoV2beta1
}

// KyvernoV2alpha1 retrieves the KyvernoV2alpha1Client
func (c *Clientset) KyvernoV2alpha1() kyvernov2alpha1.KyvernoV2alpha1Interface {
	return c.kyvernoV2alpha1
}

// Wgpolicyk8sV1alpha2 retrieves the Wgpolicyk8sV1alpha2Client
func (c *Clientset) Wgpolicyk8sV1alpha2() wgpolicyk8sv1alpha2.Wgpolicyk8sV1alpha2Interface {
	return c.wgpolicyk8sV1alpha2
}

// ReportsV1 retrieves the ReportsV1Client
func (c *Clientset) ReportsV1() reportsv1.ReportsV1Interface {
	return c.reportsV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.kyvernoV1, err = kyvernov1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kyvernoV1alpha2, err = kyvernov1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kyvernoV1beta1, err = kyvernov1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kyvernoV2, err = kyvernov2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kyvernoV2beta1, err = kyvernov2beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.kyvernoV2alpha1, err = kyvernov2alpha1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.wgpolicyk8sV1alpha2, err = wgpolicyk8sv1alpha2.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.reportsV1, err = reportsv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.kyvernoV1 = kyvernov1.New(c)
	cs.kyvernoV1alpha2 = kyvernov1alpha2.New(c)
	cs.kyvernoV1beta1 = kyvernov1beta1.New(c)
	cs.kyvernoV2 = kyvernov2.New(c)
	cs.kyvernoV2beta1 = kyvernov2beta1.New(c)
	cs.kyvernoV2alpha1 = kyvernov2alpha1.New(c)
	cs.wgpolicyk8sV1alpha2 = wgpolicyk8sv1alpha2.New(c)
	cs.reportsV1 = reportsv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
