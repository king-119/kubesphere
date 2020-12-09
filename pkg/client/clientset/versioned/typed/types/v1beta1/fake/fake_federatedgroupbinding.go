/*
Copyright 2020 The KubeSphere Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
)

// FakeFederatedGroupBindings implements FederatedGroupBindingInterface
type FakeFederatedGroupBindings struct {
	Fake *FakeTypesV1beta1
}

var federatedgroupbindingsResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedgroupbindings"}

var federatedgroupbindingsKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedGroupBinding"}

// Get takes name of the federatedGroupBinding, and returns the corresponding federatedGroupBinding object, and an error if there is any.
func (c *FakeFederatedGroupBindings) Get(name string, options v1.GetOptions) (result *v1beta1.FederatedGroupBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(federatedgroupbindingsResource, name), &v1beta1.FederatedGroupBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroupBinding), err
}

// List takes label and field selectors, and returns the list of FederatedGroupBindings that match those selectors.
func (c *FakeFederatedGroupBindings) List(opts v1.ListOptions) (result *v1beta1.FederatedGroupBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(federatedgroupbindingsResource, federatedgroupbindingsKind, opts), &v1beta1.FederatedGroupBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedGroupBindingList{ListMeta: obj.(*v1beta1.FederatedGroupBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedGroupBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedGroupBindings.
func (c *FakeFederatedGroupBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(federatedgroupbindingsResource, opts))
}

// Create takes the representation of a federatedGroupBinding and creates it.  Returns the server's representation of the federatedGroupBinding, and an error, if there is any.
func (c *FakeFederatedGroupBindings) Create(federatedGroupBinding *v1beta1.FederatedGroupBinding) (result *v1beta1.FederatedGroupBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(federatedgroupbindingsResource, federatedGroupBinding), &v1beta1.FederatedGroupBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroupBinding), err
}

// Update takes the representation of a federatedGroupBinding and updates it. Returns the server's representation of the federatedGroupBinding, and an error, if there is any.
func (c *FakeFederatedGroupBindings) Update(federatedGroupBinding *v1beta1.FederatedGroupBinding) (result *v1beta1.FederatedGroupBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(federatedgroupbindingsResource, federatedGroupBinding), &v1beta1.FederatedGroupBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroupBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedGroupBindings) UpdateStatus(federatedGroupBinding *v1beta1.FederatedGroupBinding) (*v1beta1.FederatedGroupBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(federatedgroupbindingsResource, "status", federatedGroupBinding), &v1beta1.FederatedGroupBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroupBinding), err
}

// Delete takes name of the federatedGroupBinding and deletes it. Returns an error if one occurs.
func (c *FakeFederatedGroupBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(federatedgroupbindingsResource, name), &v1beta1.FederatedGroupBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedGroupBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(federatedgroupbindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedGroupBindingList{})
	return err
}

// Patch applies the patch and returns the patched federatedGroupBinding.
func (c *FakeFederatedGroupBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FederatedGroupBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(federatedgroupbindingsResource, name, pt, data, subresources...), &v1beta1.FederatedGroupBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedGroupBinding), err
}
