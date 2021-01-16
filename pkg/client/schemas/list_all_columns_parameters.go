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

// NewListAllColumnsParams creates a new ListAllColumnsParams object
// with the default values initialized.
func NewListAllColumnsParams() *ListAllColumnsParams {
	var ()
	return &ListAllColumnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllColumnsParamsWithTimeout creates a new ListAllColumnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllColumnsParamsWithTimeout(timeout time.Duration) *ListAllColumnsParams {
	var ()
	return &ListAllColumnsParams{

		timeout: timeout,
	}
}

// NewListAllColumnsParamsWithContext creates a new ListAllColumnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllColumnsParamsWithContext(ctx context.Context) *ListAllColumnsParams {
	var ()
	return &ListAllColumnsParams{

		Context: ctx,
	}
}

// NewListAllColumnsParamsWithHTTPClient creates a new ListAllColumnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllColumnsParamsWithHTTPClient(client *http.Client) *ListAllColumnsParams {
	var ()
	return &ListAllColumnsParams{
		HTTPClient: client,
	}
}

/*ListAllColumnsParams contains all the parameters to send to the API endpoint
for the list all columns operation typically these are written to a http.Request
*/
type ListAllColumnsParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*KeyspaceName
	  Name of the keyspace to use for the request.

	*/
	KeyspaceName string
	/*TableName
	  Name of the table to use for the request.

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all columns params
func (o *ListAllColumnsParams) WithTimeout(timeout time.Duration) *ListAllColumnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all columns params
func (o *ListAllColumnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all columns params
func (o *ListAllColumnsParams) WithContext(ctx context.Context) *ListAllColumnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all columns params
func (o *ListAllColumnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all columns params
func (o *ListAllColumnsParams) WithHTTPClient(client *http.Client) *ListAllColumnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all columns params
func (o *ListAllColumnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the list all columns params
func (o *ListAllColumnsParams) WithXCassandraToken(xCassandraToken string) *ListAllColumnsParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the list all columns params
func (o *ListAllColumnsParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithKeyspaceName adds the keyspaceName to the list all columns params
func (o *ListAllColumnsParams) WithKeyspaceName(keyspaceName string) *ListAllColumnsParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the list all columns params
func (o *ListAllColumnsParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithTableName adds the tableName to the list all columns params
func (o *ListAllColumnsParams) WithTableName(tableName string) *ListAllColumnsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the list all columns params
func (o *ListAllColumnsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllColumnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	// path param keyspaceName
	if err := r.SetPathParam("keyspaceName", o.KeyspaceName); err != nil {
		return err
	}

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}