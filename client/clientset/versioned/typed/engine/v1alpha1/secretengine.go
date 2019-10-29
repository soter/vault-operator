/*
Copyright The KubeVault Authors.

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

	v1alpha1 "kubevault.dev/operator/apis/engine/v1alpha1"
	scheme "kubevault.dev/operator/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecretEnginesGetter has a method to return a SecretEngineInterface.
// A group's client should implement this interface.
type SecretEnginesGetter interface {
	SecretEngines(namespace string) SecretEngineInterface
}

// SecretEngineInterface has methods to work with SecretEngine resources.
type SecretEngineInterface interface {
	Create(*v1alpha1.SecretEngine) (*v1alpha1.SecretEngine, error)
	Update(*v1alpha1.SecretEngine) (*v1alpha1.SecretEngine, error)
	UpdateStatus(*v1alpha1.SecretEngine) (*v1alpha1.SecretEngine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecretEngine, error)
	List(opts v1.ListOptions) (*v1alpha1.SecretEngineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretEngine, err error)
	SecretEngineExpansion
}

// secretEngines implements SecretEngineInterface
type secretEngines struct {
	client rest.Interface
	ns     string
}

// newSecretEngines returns a SecretEngines
func newSecretEngines(c *EngineV1alpha1Client, namespace string) *secretEngines {
	return &secretEngines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretEngine, and returns the corresponding secretEngine object, and an error if there is any.
func (c *secretEngines) Get(name string, options v1.GetOptions) (result *v1alpha1.SecretEngine, err error) {
	result = &v1alpha1.SecretEngine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretengines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretEngines that match those selectors.
func (c *secretEngines) List(opts v1.ListOptions) (result *v1alpha1.SecretEngineList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecretEngineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretengines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretEngines.
func (c *secretEngines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretengines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a secretEngine and creates it.  Returns the server's representation of the secretEngine, and an error, if there is any.
func (c *secretEngines) Create(secretEngine *v1alpha1.SecretEngine) (result *v1alpha1.SecretEngine, err error) {
	result = &v1alpha1.SecretEngine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretengines").
		Body(secretEngine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secretEngine and updates it. Returns the server's representation of the secretEngine, and an error, if there is any.
func (c *secretEngines) Update(secretEngine *v1alpha1.SecretEngine) (result *v1alpha1.SecretEngine, err error) {
	result = &v1alpha1.SecretEngine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretengines").
		Name(secretEngine.Name).
		Body(secretEngine).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *secretEngines) UpdateStatus(secretEngine *v1alpha1.SecretEngine) (result *v1alpha1.SecretEngine, err error) {
	result = &v1alpha1.SecretEngine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretengines").
		Name(secretEngine.Name).
		SubResource("status").
		Body(secretEngine).
		Do().
		Into(result)
	return
}

// Delete takes name of the secretEngine and deletes it. Returns an error if one occurs.
func (c *secretEngines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretengines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretEngines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretengines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secretEngine.
func (c *secretEngines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretEngine, err error) {
	result = &v1alpha1.SecretEngine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretengines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
