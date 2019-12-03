// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/gardener/cert-management/pkg/apis/cert/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertificates implements CertificateInterface
type FakeCertificates struct {
	Fake *FakeCertV1alpha1
	ns   string
}

var certificatesResource = schema.GroupVersionResource{Group: "cert", Version: "v1alpha1", Resource: "certificates"}

var certificatesKind = schema.GroupVersionKind{Group: "cert", Version: "v1alpha1", Kind: "Certificate"}

// Get takes name of the certificate, and returns the corresponding certificate object, and an error if there is any.
func (c *FakeCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certificatesResource, c.ns, name), &v1alpha1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Certificate), err
}

// List takes label and field selectors, and returns the list of Certificates that match those selectors.
func (c *FakeCertificates) List(opts v1.ListOptions) (result *v1alpha1.CertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certificatesResource, certificatesKind, c.ns, opts), &v1alpha1.CertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertificateList{ListMeta: obj.(*v1alpha1.CertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.CertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certificates.
func (c *FakeCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certificatesResource, c.ns, opts))

}

// Create takes the representation of a certificate and creates it.  Returns the server's representation of the certificate, and an error, if there is any.
func (c *FakeCertificates) Create(certificate *v1alpha1.Certificate) (result *v1alpha1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certificatesResource, c.ns, certificate), &v1alpha1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Certificate), err
}

// Update takes the representation of a certificate and updates it. Returns the server's representation of the certificate, and an error, if there is any.
func (c *FakeCertificates) Update(certificate *v1alpha1.Certificate) (result *v1alpha1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certificatesResource, c.ns, certificate), &v1alpha1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Certificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCertificates) UpdateStatus(certificate *v1alpha1.Certificate) (*v1alpha1.Certificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(certificatesResource, "status", c.ns, certificate), &v1alpha1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Certificate), err
}

// Delete takes name of the certificate and deletes it. Returns an error if one occurs.
func (c *FakeCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certificatesResource, c.ns, name), &v1alpha1.Certificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertificateList{})
	return err
}

// Patch applies the patch and returns the patched certificate.
func (c *FakeCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Certificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Certificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Certificate), err
}
