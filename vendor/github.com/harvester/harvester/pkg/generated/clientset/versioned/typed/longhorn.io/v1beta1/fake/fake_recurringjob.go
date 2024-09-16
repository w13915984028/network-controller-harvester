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

package fake

import (
	"context"

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRecurringJobs implements RecurringJobInterface
type FakeRecurringJobs struct {
	Fake *FakeLonghornV1beta1
	ns   string
}

var recurringjobsResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "recurringjobs"}

var recurringjobsKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "RecurringJob"}

// Get takes name of the recurringJob, and returns the corresponding recurringJob object, and an error if there is any.
func (c *FakeRecurringJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.RecurringJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recurringjobsResource, c.ns, name), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

// List takes label and field selectors, and returns the list of RecurringJobs that match those selectors.
func (c *FakeRecurringJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.RecurringJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(recurringjobsResource, recurringjobsKind, c.ns, opts), &v1beta1.RecurringJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.RecurringJobList{ListMeta: obj.(*v1beta1.RecurringJobList).ListMeta}
	for _, item := range obj.(*v1beta1.RecurringJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested recurringJobs.
func (c *FakeRecurringJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recurringjobsResource, c.ns, opts))

}

// Create takes the representation of a recurringJob and creates it.  Returns the server's representation of the recurringJob, and an error, if there is any.
func (c *FakeRecurringJobs) Create(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.CreateOptions) (result *v1beta1.RecurringJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recurringjobsResource, c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

// Update takes the representation of a recurringJob and updates it. Returns the server's representation of the recurringJob, and an error, if there is any.
func (c *FakeRecurringJobs) Update(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.UpdateOptions) (result *v1beta1.RecurringJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recurringjobsResource, c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRecurringJobs) UpdateStatus(ctx context.Context, recurringJob *v1beta1.RecurringJob, opts v1.UpdateOptions) (*v1beta1.RecurringJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(recurringjobsResource, "status", c.ns, recurringJob), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}

// Delete takes name of the recurringJob and deletes it. Returns an error if one occurs.
func (c *FakeRecurringJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(recurringjobsResource, c.ns, name, opts), &v1beta1.RecurringJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecurringJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(recurringjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.RecurringJobList{})
	return err
}

// Patch applies the patch and returns the patched recurringJob.
func (c *FakeRecurringJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.RecurringJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recurringjobsResource, c.ns, name, pt, data, subresources...), &v1beta1.RecurringJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.RecurringJob), err
}
