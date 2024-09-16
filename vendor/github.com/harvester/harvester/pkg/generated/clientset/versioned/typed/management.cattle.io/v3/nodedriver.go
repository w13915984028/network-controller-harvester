/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeDriversGetter has a method to return a NodeDriverInterface.
// A group's client should implement this interface.
type NodeDriversGetter interface {
	NodeDrivers() NodeDriverInterface
}

// NodeDriverInterface has methods to work with NodeDriver resources.
type NodeDriverInterface interface {
	Create(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.CreateOptions) (*v3.NodeDriver, error)
	Update(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (*v3.NodeDriver, error)
	UpdateStatus(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (*v3.NodeDriver, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.NodeDriver, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.NodeDriverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.NodeDriver, err error)
	NodeDriverExpansion
}

// nodeDrivers implements NodeDriverInterface
type nodeDrivers struct {
	client rest.Interface
}

// newNodeDrivers returns a NodeDrivers
func newNodeDrivers(c *ManagementV3Client) *nodeDrivers {
	return &nodeDrivers{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeDriver, and returns the corresponding nodeDriver object, and an error if there is any.
func (c *nodeDrivers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.NodeDriver, err error) {
	result = &v3.NodeDriver{}
	err = c.client.Get().
		Resource("nodedrivers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeDrivers that match those selectors.
func (c *nodeDrivers) List(ctx context.Context, opts v1.ListOptions) (result *v3.NodeDriverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.NodeDriverList{}
	err = c.client.Get().
		Resource("nodedrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeDrivers.
func (c *nodeDrivers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodedrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeDriver and creates it.  Returns the server's representation of the nodeDriver, and an error, if there is any.
func (c *nodeDrivers) Create(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.CreateOptions) (result *v3.NodeDriver, err error) {
	result = &v3.NodeDriver{}
	err = c.client.Post().
		Resource("nodedrivers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeDriver).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeDriver and updates it. Returns the server's representation of the nodeDriver, and an error, if there is any.
func (c *nodeDrivers) Update(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (result *v3.NodeDriver, err error) {
	result = &v3.NodeDriver{}
	err = c.client.Put().
		Resource("nodedrivers").
		Name(nodeDriver.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeDriver).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodeDrivers) UpdateStatus(ctx context.Context, nodeDriver *v3.NodeDriver, opts v1.UpdateOptions) (result *v3.NodeDriver, err error) {
	result = &v3.NodeDriver{}
	err = c.client.Put().
		Resource("nodedrivers").
		Name(nodeDriver.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeDriver).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeDriver and deletes it. Returns an error if one occurs.
func (c *nodeDrivers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodedrivers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeDrivers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodedrivers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeDriver.
func (c *nodeDrivers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.NodeDriver, err error) {
	result = &v3.NodeDriver{}
	err = c.client.Patch(pt).
		Resource("nodedrivers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
