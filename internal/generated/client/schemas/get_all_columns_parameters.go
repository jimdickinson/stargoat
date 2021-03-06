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
	"github.com/go-openapi/swag"
)

// NewGetAllColumnsParams creates a new GetAllColumnsParams object
// with the default values initialized.
func NewGetAllColumnsParams() *GetAllColumnsParams {
	var (
		rawDefault = bool(false)
	)
	return &GetAllColumnsParams{
		Raw: &rawDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllColumnsParamsWithTimeout creates a new GetAllColumnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllColumnsParamsWithTimeout(timeout time.Duration) *GetAllColumnsParams {
	var (
		rawDefault = bool(false)
	)
	return &GetAllColumnsParams{
		Raw: &rawDefault,

		timeout: timeout,
	}
}

// NewGetAllColumnsParamsWithContext creates a new GetAllColumnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllColumnsParamsWithContext(ctx context.Context) *GetAllColumnsParams {
	var (
		rawDefault = bool(false)
	)
	return &GetAllColumnsParams{
		Raw: &rawDefault,

		Context: ctx,
	}
}

// NewGetAllColumnsParamsWithHTTPClient creates a new GetAllColumnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllColumnsParamsWithHTTPClient(client *http.Client) *GetAllColumnsParams {
	var (
		rawDefault = bool(false)
	)
	return &GetAllColumnsParams{
		Raw:        &rawDefault,
		HTTPClient: client,
	}
}

/*GetAllColumnsParams contains all the parameters to send to the API endpoint
for the get all columns operation typically these are written to a http.Request
*/
type GetAllColumnsParams struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*KeyspaceName
	  Name of the keyspace to use for the request.

	*/
	KeyspaceName string
	/*Raw
	  Unwrap results

	*/
	Raw *bool
	/*TableName
	  Name of the table to use for the request.

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all columns params
func (o *GetAllColumnsParams) WithTimeout(timeout time.Duration) *GetAllColumnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all columns params
func (o *GetAllColumnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all columns params
func (o *GetAllColumnsParams) WithContext(ctx context.Context) *GetAllColumnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all columns params
func (o *GetAllColumnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all columns params
func (o *GetAllColumnsParams) WithHTTPClient(client *http.Client) *GetAllColumnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all columns params
func (o *GetAllColumnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the get all columns params
func (o *GetAllColumnsParams) WithXCassandraToken(xCassandraToken string) *GetAllColumnsParams {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the get all columns params
func (o *GetAllColumnsParams) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithKeyspaceName adds the keyspaceName to the get all columns params
func (o *GetAllColumnsParams) WithKeyspaceName(keyspaceName string) *GetAllColumnsParams {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the get all columns params
func (o *GetAllColumnsParams) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithRaw adds the raw to the get all columns params
func (o *GetAllColumnsParams) WithRaw(raw *bool) *GetAllColumnsParams {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the get all columns params
func (o *GetAllColumnsParams) SetRaw(raw *bool) {
	o.Raw = raw
}

// WithTableName adds the tableName to the get all columns params
func (o *GetAllColumnsParams) WithTableName(tableName string) *GetAllColumnsParams {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get all columns params
func (o *GetAllColumnsParams) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllColumnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tableName
	if err := r.SetPathParam("tableName", o.TableName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
