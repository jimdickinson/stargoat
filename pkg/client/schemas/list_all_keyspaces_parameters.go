// Code generated by go-swagger; DO NOT EDIT.

package schemas

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

// NewListAllKeyspacesParams creates a new ListAllKeyspacesParams object
// with the default values initialized.
func NewListAllKeyspacesParams() *ListAllKeyspacesParams {
	var ()
	return &ListAllKeyspacesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllKeyspacesParamsWithTimeout creates a new ListAllKeyspacesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllKeyspacesParamsWithTimeout(timeout time.Duration) *ListAllKeyspacesParams {
	var ()
	return &ListAllKeyspacesParams{

		timeout: timeout,
	}
}

// NewListAllKeyspacesParamsWithContext creates a new ListAllKeyspacesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllKeyspacesParamsWithContext(ctx context.Context) *ListAllKeyspacesParams {
	var ()
	return &ListAllKeyspacesParams{

		Context: ctx,
	}
}

// NewListAllKeyspacesParamsWithHTTPClient creates a new ListAllKeyspacesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllKeyspacesParamsWithHTTPClient(client *http.Client) *ListAllKeyspacesParams {
	var ()
	return &ListAllKeyspacesParams{
		HTTPClient: client,
	}
}

/*ListAllKeyspacesParams contains all the parameters to send to the API endpoint
for the list all keyspaces operation typically these are written to a http.Request
*/
type ListAllKeyspacesParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all keyspaces params
func (o *ListAllKeyspacesParams) WithTimeout(timeout time.Duration) *ListAllKeyspacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all keyspaces params
func (o *ListAllKeyspacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all keyspaces params
func (o *ListAllKeyspacesParams) WithContext(ctx context.Context) *ListAllKeyspacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all keyspaces params
func (o *ListAllKeyspacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all keyspaces params
func (o *ListAllKeyspacesParams) WithHTTPClient(client *http.Client) *ListAllKeyspacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all keyspaces params
func (o *ListAllKeyspacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the list all keyspaces params
func (o *ListAllKeyspacesParams) WithXCassandraToken(xCassandraToken string) *ListAllKeyspacesParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the list all keyspaces params
func (o *ListAllKeyspacesParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllKeyspacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
