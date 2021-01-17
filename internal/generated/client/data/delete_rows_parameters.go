// Code generated by go-swagger; DO NOT EDIT.

package data

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

// NewDeleteRowsParams creates a new DeleteRowsParams object
// with the default values initialized.
func NewDeleteRowsParams() *DeleteRowsParams {
	var ()
	return &DeleteRowsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRowsParamsWithTimeout creates a new DeleteRowsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRowsParamsWithTimeout(timeout time.Duration) *DeleteRowsParams {
	var ()
	return &DeleteRowsParams{

		timeout: timeout,
	}
}

// NewDeleteRowsParamsWithContext creates a new DeleteRowsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRowsParamsWithContext(ctx context.Context) *DeleteRowsParams {
	var ()
	return &DeleteRowsParams{

		Context: ctx,
	}
}

// NewDeleteRowsParamsWithHTTPClient creates a new DeleteRowsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRowsParamsWithHTTPClient(client *http.Client) *DeleteRowsParams {
	var ()
	return &DeleteRowsParams{
		HTTPClient: client,
	}
}

/*DeleteRowsParams contains all the parameters to send to the API endpoint
for the delete rows operation typically these are written to a http.Request
*/
type DeleteRowsParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*KeyspaceName
	  Name of the keyspace to use for the request.

	*/
	KeyspaceName string
	/*PrimaryKey
	  Value from the primary key column for the table. Define composite keys by separating values with slashes (`val1/val2...`) in the order they were defined. </br> For example, if the composite key was defined as `PRIMARY KEY(race_year, race_name)` then the primary key in the path would be `race_year/race_name`

	*/
	PrimaryKey []string
	/*TableName
	  Name of the table to use for the request.

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete rows params
func (o *DeleteRowsParams) WithTimeout(timeout time.Duration) *DeleteRowsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete rows params
func (o *DeleteRowsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete rows params
func (o *DeleteRowsParams) WithContext(ctx context.Context) *DeleteRowsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete rows params
func (o *DeleteRowsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete rows params
func (o *DeleteRowsParams) WithHTTPClient(client *http.Client) *DeleteRowsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete rows params
func (o *DeleteRowsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the delete rows params
func (o *DeleteRowsParams) WithXCassandraToken(xCassandraToken string) *DeleteRowsParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the delete rows params
func (o *DeleteRowsParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithKeyspaceName adds the keyspaceName to the delete rows params
func (o *DeleteRowsParams) WithKeyspaceName(keyspaceName string) *DeleteRowsParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the delete rows params
func (o *DeleteRowsParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithPrimaryKey adds the primaryKey to the delete rows params
func (o *DeleteRowsParams) WithPrimaryKey(primaryKey []string) *DeleteRowsParams {
	o.SetPrimaryKey(primaryKey)
	return o
}

// SetPrimaryKey adds the primaryKey to the delete rows params
func (o *DeleteRowsParams) SetPrimaryKey(primaryKey []string) {
	o.PrimaryKey = primaryKey
}

// WithTableName adds the tableName to the delete rows params
func (o *DeleteRowsParams) WithTableName(tableName string) *DeleteRowsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the delete rows params
func (o *DeleteRowsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesPrimaryKey := o.PrimaryKey

	joinedPrimaryKey := swag.JoinByFormat(valuesPrimaryKey, "csv")
	// path array param primaryKey
	// SetPathParam does not support variadric arguments, since we used JoinByFormat
	// we can send the first item in the array as it's all the items of the previous
	// array joined together
	if len(joinedPrimaryKey) > 0 {
		if err := r.SetPathParam("primaryKey", joinedPrimaryKey[0]); err != nil {
			return err
		}
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