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

// NewPostDocParams creates a new PostDocParams object
// with the default values initialized.
func NewPostDocParams() *PostDocParams {
	var ()
	return &PostDocParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostDocParamsWithTimeout creates a new PostDocParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostDocParamsWithTimeout(timeout time.Duration) *PostDocParams {
	var ()
	return &PostDocParams{

		timeout: timeout,
	}
}

// NewPostDocParamsWithContext creates a new PostDocParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostDocParamsWithContext(ctx context.Context) *PostDocParams {
	var ()
	return &PostDocParams{

		Context: ctx,
	}
}

// NewPostDocParamsWithHTTPClient creates a new PostDocParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostDocParamsWithHTTPClient(client *http.Client) *PostDocParams {
	var ()
	return &PostDocParams{
		HTTPClient: client,
	}
}

/*PostDocParams contains all the parameters to send to the API endpoint
for the post doc operation typically these are written to a http.Request
*/
type PostDocParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*Body
	  The JSON document

	*/
	Body string
	/*CollectionID
	  the name of the collection

	*/
	CollectionID string
	/*NamespaceID
	  the namespace that the collection is in

	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post doc params
func (o *PostDocParams) WithTimeout(timeout time.Duration) *PostDocParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post doc params
func (o *PostDocParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post doc params
func (o *PostDocParams) WithContext(ctx context.Context) *PostDocParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post doc params
func (o *PostDocParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post doc params
func (o *PostDocParams) WithHTTPClient(client *http.Client) *PostDocParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post doc params
func (o *PostDocParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the post doc params
func (o *PostDocParams) WithXCassandraToken(xCassandraToken string) *PostDocParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the post doc params
func (o *PostDocParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithBody adds the body to the post doc params
func (o *PostDocParams) WithBody(body string) *PostDocParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post doc params
func (o *PostDocParams) SetBody(body string) {
	o.Body = body
}

// WithCollectionID adds the collectionID to the post doc params
func (o *PostDocParams) WithCollectionID(collectionID string) *PostDocParams {
	o.SetCollectionID(collectionID)
	return o
}

// SetCollectionID adds the collectionId to the post doc params
func (o *PostDocParams) SetCollectionID(collectionID string) {
	o.CollectionID = collectionID
}

// WithNamespaceID adds the namespaceID to the post doc params
func (o *PostDocParams) WithNamespaceID(namespaceID string) *PostDocParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the post doc params
func (o *PostDocParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PostDocParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
