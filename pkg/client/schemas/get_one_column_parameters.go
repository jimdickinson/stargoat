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

// NewGetOneColumnParams creates a new GetOneColumnParams object
// with the default values initialized.
func NewGetOneColumnParams() *GetOneColumnParams {
	var ()
	return &GetOneColumnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOneColumnParamsWithTimeout creates a new GetOneColumnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOneColumnParamsWithTimeout(timeout time.Duration) *GetOneColumnParams {
	var ()
	return &GetOneColumnParams{

		timeout: timeout,
	}
}

// NewGetOneColumnParamsWithContext creates a new GetOneColumnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOneColumnParamsWithContext(ctx context.Context) *GetOneColumnParams {
	var ()
	return &GetOneColumnParams{

		Context: ctx,
	}
}

// NewGetOneColumnParamsWithHTTPClient creates a new GetOneColumnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOneColumnParamsWithHTTPClient(client *http.Client) *GetOneColumnParams {
	var ()
	return &GetOneColumnParams{
		HTTPClient: client,
	}
}

/*GetOneColumnParams contains all the parameters to send to the API endpoint
for the get one column operation typically these are written to a http.Request
*/
type GetOneColumnParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*ColumnName
	  Name of the column to use for the request.

	*/
	ColumnName string
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

// WithTimeout adds the timeout to the get one column params
func (o *GetOneColumnParams) WithTimeout(timeout time.Duration) *GetOneColumnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get one column params
func (o *GetOneColumnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get one column params
func (o *GetOneColumnParams) WithContext(ctx context.Context) *GetOneColumnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get one column params
func (o *GetOneColumnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get one column params
func (o *GetOneColumnParams) WithHTTPClient(client *http.Client) *GetOneColumnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get one column params
func (o *GetOneColumnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the get one column params
func (o *GetOneColumnParams) WithXCassandraToken(xCassandraToken string) *GetOneColumnParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the get one column params
func (o *GetOneColumnParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithColumnName adds the columnName to the get one column params
func (o *GetOneColumnParams) WithColumnName(columnName string) *GetOneColumnParams {
	o.SetColumnName(columnName)
	return o
}

// SetColumnName adds the columnName to the get one column params
func (o *GetOneColumnParams) SetColumnName(columnName string) {
	o.ColumnName = columnName
}

// WithKeyspaceName adds the keyspaceName to the get one column params
func (o *GetOneColumnParams) WithKeyspaceName(keyspaceName string) *GetOneColumnParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the get one column params
func (o *GetOneColumnParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithTableName adds the tableName to the get one column params
func (o *GetOneColumnParams) WithTableName(tableName string) *GetOneColumnParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get one column params
func (o *GetOneColumnParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *GetOneColumnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	// path param columnName
	if err := r.SetPathParam("columnName", o.ColumnName); err != nil {
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
