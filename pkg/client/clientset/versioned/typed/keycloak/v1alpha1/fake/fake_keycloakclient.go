// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKeycloakClients implements KeycloakClientInterface
type FakeKeycloakClients struct {
	Fake *FakeKeycloakV1alpha1
	ns   string
}

var keycloakclientsResource = schema.GroupVersionResource{Group: "keycloak.org", Version: "v1alpha1", Resource: "keycloakclients"}

var keycloakclientsKind = schema.GroupVersionKind{Group: "keycloak.org", Version: "v1alpha1", Kind: "KeycloakClient"}

// Get takes name of the keycloakClient, and returns the corresponding keycloakClient object, and an error if there is any.
func (c *FakeKeycloakClients) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeycloakClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keycloakclientsResource, c.ns, name), &v1alpha1.KeycloakClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakClient), err
}

// List takes label and field selectors, and returns the list of KeycloakClients that match those selectors.
func (c *FakeKeycloakClients) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeycloakClientList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keycloakclientsResource, keycloakclientsKind, c.ns, opts), &v1alpha1.KeycloakClientList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeycloakClientList{ListMeta: obj.(*v1alpha1.KeycloakClientList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeycloakClientList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keycloakClients.
func (c *FakeKeycloakClients) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keycloakclientsResource, c.ns, opts))

}

// Create takes the representation of a keycloakClient and creates it.  Returns the server's representation of the keycloakClient, and an error, if there is any.
func (c *FakeKeycloakClients) Create(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.CreateOptions) (result *v1alpha1.KeycloakClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keycloakclientsResource, c.ns, keycloakClient), &v1alpha1.KeycloakClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakClient), err
}

// Update takes the representation of a keycloakClient and updates it. Returns the server's representation of the keycloakClient, and an error, if there is any.
func (c *FakeKeycloakClients) Update(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (result *v1alpha1.KeycloakClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keycloakclientsResource, c.ns, keycloakClient), &v1alpha1.KeycloakClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakClient), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeycloakClients) UpdateStatus(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (*v1alpha1.KeycloakClient, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keycloakclientsResource, "status", c.ns, keycloakClient), &v1alpha1.KeycloakClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakClient), err
}

// Delete takes name of the keycloakClient and deletes it. Returns an error if one occurs.
func (c *FakeKeycloakClients) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keycloakclientsResource, c.ns, name), &v1alpha1.KeycloakClient{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeycloakClients) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keycloakclientsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeycloakClientList{})
	return err
}

// Patch applies the patch and returns the patched keycloakClient.
func (c *FakeKeycloakClients) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeycloakClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keycloakclientsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeycloakClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakClient), err
}