// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilds implements BuildInterface
type FakeBuilds struct {
	Fake *FakeConfigV1
}

var buildsResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "builds"}

var buildsKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "Build"}

// Get takes name of the build, and returns the corresponding build object, and an error if there is any.
func (c *FakeBuilds) Get(ctx context.Context, name string, options v1.GetOptions) (result *configv1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(buildsResource, name), &configv1.Build{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Build), err
}

// List takes label and field selectors, and returns the list of Builds that match those selectors.
func (c *FakeBuilds) List(ctx context.Context, opts v1.ListOptions) (result *configv1.BuildList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(buildsResource, buildsKind, opts), &configv1.BuildList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.BuildList{ListMeta: obj.(*configv1.BuildList).ListMeta}
	for _, item := range obj.(*configv1.BuildList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builds.
func (c *FakeBuilds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(buildsResource, opts))
}

// Create takes the representation of a build and creates it.  Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuilds) Create(ctx context.Context, build *configv1.Build, opts v1.CreateOptions) (result *configv1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(buildsResource, build), &configv1.Build{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Build), err
}

// Update takes the representation of a build and updates it. Returns the server's representation of the build, and an error, if there is any.
func (c *FakeBuilds) Update(ctx context.Context, build *configv1.Build, opts v1.UpdateOptions) (result *configv1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(buildsResource, build), &configv1.Build{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Build), err
}

// Delete takes name of the build and deletes it. Returns an error if one occurs.
func (c *FakeBuilds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(buildsResource, name), &configv1.Build{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuilds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(buildsResource, listOpts)

	_, err := c.Fake.Invokes(action, &configv1.BuildList{})
	return err
}

// Patch applies the patch and returns the patched build.
func (c *FakeBuilds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configv1.Build, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(buildsResource, name, pt, data, subresources...), &configv1.Build{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Build), err
}
