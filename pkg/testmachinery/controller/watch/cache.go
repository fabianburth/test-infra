// Copyright 2020 Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package watch

import (
	"context"
	"fmt"

	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/client-go/rest"
	toolscache "k8s.io/client-go/tools/cache"
	ctrlcache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

// CachedInformerType specifies the cached informer type.
const CachedInformerType InformerType = "cached"

type cachedInformer struct {
	log      logr.Logger
	client   client.Client
	cache    ctrlcache.Cache
	eventbus EventBus
}

func newCachedInformer(log logr.Logger, config *rest.Config, options *Options) (Informer, error) {
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, err
	}
	mapper, err := apiutil.NewDynamicRESTMapper(config, httpClient)
	if err != nil {
		return nil, err
	}

	// Create the informer for the cached read client and registering informers
	cache, err := ctrlcache.New(config, ctrlcache.Options{Scheme: options.Scheme, Mapper: mapper, SyncPeriod: options.SyncPeriod, DefaultNamespaces: map[string]ctrlcache.Config{
		options.Namespace: {},
	}})
	if err != nil {
		return nil, err
	}

	clientOpts := client.Options{}
	clientOpts.Scheme = options.Scheme
	clientOpts.Mapper = mapper
	clientOpts.Cache = &client.CacheOptions{}
	clientOpts.Cache.Reader = cache

	cachedClient, err := client.New(config, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("unable to create cached client: %w", err)
	}
	return &cachedInformer{
		log:    log,
		client: cachedClient,
		cache:  cache,
	}, nil
}

func (c *cachedInformer) Start(ctx context.Context) error {
	i, err := c.cache.GetInformer(ctx, &tmv1beta1.Testrun{})
	if err != nil {
		if kindMatchErr, ok := err.(*meta.NoKindMatchError); ok {
			c.log.Error(err, "if kind is a CRD, it should be installed before calling Start",
				"kind", kindMatchErr.GroupKind)
		}
		return err
	}
	_, err = i.AddEventHandler(toolscache.ResourceEventHandlerFuncs{
		AddFunc:    c.addItemToQueue,
		UpdateFunc: func(old, new interface{}) { c.addItemToQueue(new) },
		DeleteFunc: c.addItemToQueue,
	})
	if err != nil {
		return err
	}

	return c.cache.Start(ctx)
}

func (c *cachedInformer) WaitForCacheSync(ctx context.Context) bool {
	return c.cache.WaitForCacheSync(ctx)
}

func (c *cachedInformer) InjectEventBus(eb EventBus) {
	c.eventbus = eb
}

func (c *cachedInformer) Client() client.Client {
	return c.client
}

func (c *cachedInformer) addItemToQueue(obj interface{}) {
	// try to cast to testrun ignore the event if this is not possible
	tr, ok := obj.(*tmv1beta1.Testrun)
	if !ok {
		return
	}
	c.eventbus.Publish(keyOfTestrun(tr), tr)
}
