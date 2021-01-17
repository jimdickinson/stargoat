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

// NewDeleteColumnParams creates a new DeleteColumnParams object
// with the default values initialized.
func NewDeleteColumnParams() *DeleteColumnParams {
	var ()
	return &DeleteColumnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteColumnParamsWithTimeout creates a new DeleteColumnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteColumnParamsWithTimeout(timeout time.Duration) *DeleteColumnParams {
	var ()
	return &DeleteColumnParams{

		timeout: timeout,
	}
}

// NewDeleteColumnParamsWithContext creates a new DeleteColumnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteColumnParamsWithContext(ctx context.Context) *DeleteColumnParams {
	var ()
	return &DeleteColumnParams{

		Context: ctx,
	}
}

// NewDeleteColumnParamsWithHTTPClient creates a new DeleteColumnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteColumnParamsWithHTTPClient(client *http.Client) *DeleteColumnParams {
	var ()
	return &DeleteColumnParams{
		HTTPClient: client,
	}
}

/*DeleteColumnParams contains all the parameters to send to the API endpoint
for the delete column operation typically these are written to a http.Request
*/
type DeleteColumnParams struct {

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

// WithTimeout adds the timeout to the delete column params
func (o *DeleteColumnParams) WithTimeout(timeout time.Duration) *DeleteColumnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete column params
func (o *DeleteColumnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete column params
func (o *DeleteColumnParams) WithContext(ctx context.Context) *DeleteColumnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete column params
func (o *DeleteColumnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete column params
func (o *DeleteColumnParams) WithHTTPClient(client *http.Client) *DeleteColumnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete column params
func (o *DeleteColumnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the delete column params
func (o *DeleteColumnParams) WithXCassandraToken(xCassandraToken string) *DeleteColumnParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the delete column params
func (o *DeleteColumnParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithColumnName adds the columnName to the delete column params
func (o *DeleteColumnParams) WithColumnName(columnName string) *DeleteColumnParams {
	o.SetColumnName(columnName)
	return o
}

// SetColumnName adds the columnName to the delete column params
func (o *DeleteColumnParams) SetColumnName(columnName string) {
	o.ColumnName = columnName
}

// WithKeyspaceName adds the keyspaceName to the delete column params
func (o *DeleteColumnParams) WithKeyspaceName(keyspaceName string) *DeleteColumnParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the delete column params
func (o *DeleteColumnParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithTableName adds the tableName to the delete column params
func (o *DeleteColumnParams) WithTableName(tableName string) *DeleteColumnParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the delete column params
func (o *DeleteColumnParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteColumnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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