/*
Copyright 2020 The vm import Authors.

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

package v1alpha1

import (
	"time"

	scheme "github.com/kubevirt/vm-import-operator/pkg/api-client/clientset/versioned/scheme"
	v1alpha1 "github.com/kubevirt/vm-import-operator/pkg/apis/v2v/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceMappingsGetter has a method to return a ResourceMappingInterface.
// A group's client should implement this interface.
type ResourceMappingsGetter interface {
	ResourceMappings(namespace string) ResourceMappingInterface
}

// ResourceMappingInterface has methods to work with ResourceMapping resources.
type ResourceMappingInterface interface {
	Create(*v1alpha1.ResourceMapping) (*v1alpha1.ResourceMapping, error)
	Update(*v1alpha1.ResourceMapping) (*v1alpha1.ResourceMapping, error)
	UpdateStatus(*v1alpha1.ResourceMapping) (*v1alpha1.ResourceMapping, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ResourceMapping, error)
	List(opts v1.ListOptions) (*v1alpha1.ResourceMappingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceMapping, err error)
	ResourceMappingExpansion
}

// resourceMappings implements ResourceMappingInterface
type resourceMappings struct {
	client rest.Interface
	ns     string
}

// newResourceMappings returns a ResourceMappings
func newResourceMappings(c *V2vV1alpha1Client, namespace string) *resourceMappings {
	return &resourceMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceMapping, and returns the corresponding resourceMapping object, and an error if there is any.
func (c *resourceMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceMapping, err error) {
	result = &v1alpha1.ResourceMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcemappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceMappings that match those selectors.
func (c *resourceMappings) List(opts v1.ListOptions) (result *v1alpha1.ResourceMappingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceMappings.
func (c *resourceMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a resourceMapping and creates it.  Returns the server's representation of the resourceMapping, and an error, if there is any.
func (c *resourceMappings) Create(resourceMapping *v1alpha1.ResourceMapping) (result *v1alpha1.ResourceMapping, err error) {
	result = &v1alpha1.ResourceMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcemappings").
		Body(resourceMapping).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resourceMapping and updates it. Returns the server's representation of the resourceMapping, and an error, if there is any.
func (c *resourceMappings) Update(resourceMapping *v1alpha1.ResourceMapping) (result *v1alpha1.ResourceMapping, err error) {
	result = &v1alpha1.ResourceMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcemappings").
		Name(resourceMapping.Name).
		Body(resourceMapping).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resourceMappings) UpdateStatus(resourceMapping *v1alpha1.ResourceMapping) (result *v1alpha1.ResourceMapping, err error) {
	result = &v1alpha1.ResourceMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcemappings").
		Name(resourceMapping.Name).
		SubResource("status").
		Body(resourceMapping).
		Do().
		Into(result)
	return
}

// Delete takes name of the resourceMapping and deletes it. Returns an error if one occurs.
func (c *resourceMappings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcemappings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcemappings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resourceMapping.
func (c *resourceMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceMapping, err error) {
	result = &v1alpha1.ResourceMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcemappings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
