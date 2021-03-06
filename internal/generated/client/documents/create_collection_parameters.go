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
)

// NewCreateCollectionParams creates a new CreateCollectionParams object
// with the default values initialized.
func NewCreateCollectionParams() *CreateCollectionParams {
	var ()
	return &CreateCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCollectionParamsWithTimeout creates a new CreateCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCollectionParamsWithTimeout(timeout time.Duration) *CreateCollectionParams {
	var ()
	return &CreateCollectionParams{

		timeout: timeout,
	}
}

// NewCreateCollectionParamsWithContext creates a new CreateCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCollectionParamsWithContext(ctx context.Context) *CreateCollectionParams {
	var ()
	return &CreateCollectionParams{

		Context: ctx,
	}
}

// NewCreateCollectionParamsWithHTTPClient creates a new CreateCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCollectionParamsWithHTTPClient(client *http.Client) *CreateCollectionParams {
	var ()
	return &CreateCollectionParams{
		HTTPClient: client,
	}
}

/*CreateCollectionParams contains all the parameters to send to the API endpoint
for the create collection operation typically these are written to a http.Request
*/
type CreateCollectionParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*Body
	  JSON with the name of the collection

	*/
	Body string
	/*NamespaceID
	  the namespace to create the collection in

	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create collection params
func (o *CreateCollectionParams) WithTimeout(timeout time.Duration) *CreateCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create collection params
func (o *CreateCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create collection params
func (o *CreateCollectionParams) WithContext(ctx context.Context) *CreateCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create collection params
func (o *CreateCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create collection params
func (o *CreateCollectionParams) WithHTTPClient(client *http.Client) *CreateCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create collection params
func (o *CreateCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the create collection params
func (o *CreateCollectionParams) WithXCassandraToken(xCassandraToken string) *CreateCollectionParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the create collection params
func (o *CreateCollectionParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithBody adds the body to the create collection params
func (o *CreateCollectionParams) WithBody(body string) *CreateCollectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create collection params
func (o *CreateCollectionParams) SetBody(body string) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the create collection params
func (o *CreateCollectionParams) WithNamespaceID(namespaceID string) *CreateCollectionParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the create collection params
func (o *CreateCollectionParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace-id
	if err := r.SetPathParam("namespace-id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
