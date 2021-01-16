// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUpgradeCollectionParams creates a new UpgradeCollectionParams object
// with the default values initialized.
func NewUpgradeCollectionParams() *UpgradeCollectionParams {
	var (
		rawDefault = bool(false)
	)
	return &UpgradeCollectionParams{
		Raw: &rawDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCollectionParamsWithTimeout creates a new UpgradeCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeCollectionParamsWithTimeout(timeout time.Duration) *UpgradeCollectionParams {
	var (
		rawDefault = bool(false)
	)
	return &UpgradeCollectionParams{
		Raw: &rawDefault,

		timeout: timeout,
	}
}

// NewUpgradeCollectionParamsWithContext creates a new UpgradeCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeCollectionParamsWithContext(ctx context.Context) *UpgradeCollectionParams {
	var (
		rawDefault = bool(false)
	)
	return &UpgradeCollectionParams{
		Raw: &rawDefault,

		Context: ctx,
	}
}

// NewUpgradeCollectionParamsWithHTTPClient creates a new UpgradeCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeCollectionParamsWithHTTPClient(client *http.Client) *UpgradeCollectionParams {
	var (
		rawDefault = bool(false)
	)
	return &UpgradeCollectionParams{
		Raw:        &rawDefault,
		HTTPClient: client,
	}
}

/*UpgradeCollectionParams contains all the parameters to send to the API endpoint
for the upgrade collection operation typically these are written to a http.Request
*/
type UpgradeCollectionParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*Body
	  JSON with the upgrade type

	*/
	Body string
	/*CollectionID
	  the collection to upgrade

	*/
	CollectionID string
	/*NamespaceID
	  the namespace containing the collection to upgrade

	*/
	NamespaceID string
	/*Raw
	  Unwrap results

	*/
	Raw *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade collection params
func (o *UpgradeCollectionParams) WithTimeout(timeout time.Duration) *UpgradeCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade collection params
func (o *UpgradeCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade collection params
func (o *UpgradeCollectionParams) WithContext(ctx context.Context) *UpgradeCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade collection params
func (o *UpgradeCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade collection params
func (o *UpgradeCollectionParams) WithHTTPClient(client *http.Client) *UpgradeCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade collection params
func (o *UpgradeCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the upgrade collection params
func (o *UpgradeCollectionParams) WithXCassandraToken(xCassandraToken string) *UpgradeCollectionParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the upgrade collection params
func (o *UpgradeCollectionParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithBody adds the body to the upgrade collection params
func (o *UpgradeCollectionParams) WithBody(body string) *UpgradeCollectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade collection params
func (o *UpgradeCollectionParams) SetBody(body string) {
	o.Body = body
}

// WithCollectionID adds the collectionID to the upgrade collection params
func (o *UpgradeCollectionParams) WithCollectionID(collectionID string) *UpgradeCollectionParams {
	o.SetCollectionID(collectionID)
	return o
}

// SetCollectionID adds the collectionId to the upgrade collection params
func (o *UpgradeCollectionParams) SetCollectionID(collectionID string) {
	o.CollectionID = collectionID
}

// WithNamespaceID adds the namespaceID to the upgrade collection params
func (o *UpgradeCollectionParams) WithNamespaceID(namespaceID string) *UpgradeCollectionParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the upgrade collection params
func (o *UpgradeCollectionParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRaw adds the raw to the upgrade collection params
func (o *UpgradeCollectionParams) WithRaw(raw *bool) *UpgradeCollectionParams {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the upgrade collection params
func (o *UpgradeCollectionParams) SetRaw(raw *bool) {
	o.Raw = raw
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param collection-id
	if err := r.SetPathParam("collection-id", o.CollectionID); err != nil {
		return err
	}

	// path param namespace-id
	if err := r.SetPathParam("namespace-id", o.NamespaceID); err != nil {
		return err
	}

	if o.Raw != nil {

		// query param raw
		var qrRaw bool
		if o.Raw != nil {
			qrRaw = *o.Raw
		}
		qRaw := swag.FormatBool(qrRaw)
		if qRaw != "" {
			if err := r.SetQueryParam("raw", qRaw); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
