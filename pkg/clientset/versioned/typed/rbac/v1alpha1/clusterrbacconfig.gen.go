// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "istio.io/client-go/pkg/apis/rbac/v1alpha1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterRbacConfigsGetter has a method to return a ClusterRbacConfigInterface.
// A group's client should implement this interface.
type ClusterRbacConfigsGetter interface {
	ClusterRbacConfigs() ClusterRbacConfigInterface
}

// ClusterRbacConfigInterface has methods to work with ClusterRbacConfig resources.
type ClusterRbacConfigInterface interface {
	Create(ctx context.Context, clusterRbacConfig *v1alpha1.ClusterRbacConfig, opts v1.CreateOptions) (*v1alpha1.ClusterRbacConfig, error)
	Update(ctx context.Context, clusterRbacConfig *v1alpha1.ClusterRbacConfig, opts v1.UpdateOptions) (*v1alpha1.ClusterRbacConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterRbacConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterRbacConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRbacConfig, err error)
	ClusterRbacConfigExpansion
}

// clusterRbacConfigs implements ClusterRbacConfigInterface
type clusterRbacConfigs struct {
	client rest.Interface
}

// newClusterRbacConfigs returns a ClusterRbacConfigs
func newClusterRbacConfigs(c *RbacV1alpha1Client) *clusterRbacConfigs {
	return &clusterRbacConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterRbacConfig, and returns the corresponding clusterRbacConfig object, and an error if there is any.
func (c *clusterRbacConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterRbacConfig, err error) {
	result = &v1alpha1.ClusterRbacConfig{}
	err = c.client.Get().
		Resource("clusterrbacconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterRbacConfigs that match those selectors.
func (c *clusterRbacConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterRbacConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterRbacConfigList{}
	err = c.client.Get().
		Resource("clusterrbacconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterRbacConfigs.
func (c *clusterRbacConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clusterrbacconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterRbacConfig and creates it.  Returns the server's representation of the clusterRbacConfig, and an error, if there is any.
func (c *clusterRbacConfigs) Create(ctx context.Context, clusterRbacConfig *v1alpha1.ClusterRbacConfig, opts v1.CreateOptions) (result *v1alpha1.ClusterRbacConfig, err error) {
	result = &v1alpha1.ClusterRbacConfig{}
	err = c.client.Post().
		Resource("clusterrbacconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRbacConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterRbacConfig and updates it. Returns the server's representation of the clusterRbacConfig, and an error, if there is any.
func (c *clusterRbacConfigs) Update(ctx context.Context, clusterRbacConfig *v1alpha1.ClusterRbacConfig, opts v1.UpdateOptions) (result *v1alpha1.ClusterRbacConfig, err error) {
	result = &v1alpha1.ClusterRbacConfig{}
	err = c.client.Put().
		Resource("clusterrbacconfigs").
		Name(clusterRbacConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterRbacConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterRbacConfig and deletes it. Returns an error if one occurs.
func (c *clusterRbacConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clusterrbacconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterRbacConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clusterrbacconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterRbacConfig.
func (c *clusterRbacConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRbacConfig, err error) {
	result = &v1alpha1.ClusterRbacConfig{}
	err = c.client.Patch(pt).
		Resource("clusterrbacconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
