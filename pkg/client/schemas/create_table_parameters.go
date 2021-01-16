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

	"github.com/jimdickinson/stargoat/pkg/models"
)

// NewCreateTableParams creates a new CreateTableParams object
// with the default values initialized.
func NewCreateTableParams() *CreateTableParams {
	var ()
	return &CreateTableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTableParamsWithTimeout creates a new CreateTableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTableParamsWithTimeout(timeout time.Duration) *CreateTableParams {
	var ()
	return &CreateTableParams{

		timeout: timeout,
	}
}

// NewCreateTableParamsWithContext creates a new CreateTableParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTableParamsWithContext(ctx context.Context) *CreateTableParams {
	var ()
	return &CreateTableParams{

		Context: ctx,
	}
}

// NewCreateTableParamsWithHTTPClient creates a new CreateTableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTableParamsWithHTTPClient(client *http.Client) *CreateTableParams {
	var ()
	return &CreateTableParams{
		HTTPClient: client,
	}
}

/*CreateTableParams contains all the parameters to send to the API endpoint
for the create table operation typically these are written to a http.Request
*/
type CreateTableParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*Body*/
	Body *models.TableAdd
	/*KeyspaceName
	  Name of the keyspace to use for the request.

	*/
	KeyspaceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create table params
func (o *CreateTableParams) WithTimeout(timeout time.Duration) *CreateTableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create table params
func (o *CreateTableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create table params
func (o *CreateTableParams) WithContext(ctx context.Context) *CreateTableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create table params
func (o *CreateTableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create table params
func (o *CreateTableParams) WithHTTPClient(client *http.Client) *CreateTableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create table params
func (o *CreateTableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the create table params
func (o *CreateTableParams) WithXCassandraToken(xCassandraToken string) *CreateTableParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the create table params
func (o *CreateTableParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithBody adds the body to the create table params
func (o *CreateTableParams) WithBody(body *models.TableAdd) *CreateTableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create table params
func (o *CreateTableParams) SetBody(body *models.TableAdd) {
	o.Body = body
}

// WithKeyspaceName adds the keyspaceName to the create table params
func (o *CreateTableParams) WithKeyspaceName(keyspaceName string) *CreateTableParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the create table params
func (o *CreateTableParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param keyspaceName
	if err := r.SetPathParam("keyspaceName", o.KeyspaceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
