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

// NewGetRows1Params creates a new GetRows1Params object
// with the default values initialized.
func NewGetRows1Params() *GetRows1Params {
	var (
		rawDefault = bool(false)
	)
	return &GetRows1Params{
		Raw: &rawDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRows1ParamsWithTimeout creates a new GetRows1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRows1ParamsWithTimeout(timeout time.Duration) *GetRows1Params {
	var (
		rawDefault = bool(false)
	)
	return &GetRows1Params{
		Raw: &rawDefault,

		timeout: timeout,
	}
}

// NewGetRows1ParamsWithContext creates a new GetRows1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRows1ParamsWithContext(ctx context.Context) *GetRows1Params {
	var (
		rawDefault = bool(false)
	)
	return &GetRows1Params{
		Raw: &rawDefault,

		Context: ctx,
	}
}

// NewGetRows1ParamsWithHTTPClient creates a new GetRows1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRows1ParamsWithHTTPClient(client *http.Client) *GetRows1Params {
	var (
		rawDefault = bool(false)
	)
	return &GetRows1Params{
		Raw:        &rawDefault,
		HTTPClient: client,
	}
}

/*GetRows1Params contains all the parameters to send to the API endpoint
for the get rows 1 operation typically these are written to a http.Request
*/
type GetRows1Params struct {

	/*XCassandraToken
	  The token returned from the authorization endpoint. Use this token in each request.

	*/
	XCassandraToken string
	/*Fields
	  URL escaped, comma delimited list of keys to include

	*/
	Fields *string
	/*KeyspaceName
	  Name of the keyspace to use for the request.

	*/
	KeyspaceName string
	/*PageSize
	  Restrict the number of returned items

	*/
	PageSize *int32
	/*PageState
	  Move the cursor to a particular result

	*/
	PageState *string
	/*PrimaryKey
	  Value from the primary key column for the table. Define composite keys by separating values with slashes (`val1/val2...`) in the order they were defined. </br> For example, if the composite key was defined as `PRIMARY KEY(race_year, race_name)` then the primary key in the path would be `race_year/race_name`

	*/
	PrimaryKey []string
	/*Raw
	  Unwrap results

	*/
	Raw *bool
	/*Sort
	  Keys to sort by

	*/
	Sort *string
	/*TableName
	  Name of the table to use for the request.

	*/
	TableName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get rows 1 params
func (o *GetRows1Params) WithTimeout(timeout time.Duration) *GetRows1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rows 1 params
func (o *GetRows1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rows 1 params
func (o *GetRows1Params) WithContext(ctx context.Context) *GetRows1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rows 1 params
func (o *GetRows1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rows 1 params
func (o *GetRows1Params) WithHTTPClient(client *http.Client) *GetRows1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rows 1 params
func (o *GetRows1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXCassandraToken adds the xCassandraToken to the get rows 1 params
func (o *GetRows1Params) WithXCassandraToken(xCassandraToken string) *GetRows1Params {
	o.SetXCassandraToken(xCassandraToken)
	return o
}

// SetXCassandraToken adds the xCassandraToken to the get rows 1 params
func (o *GetRows1Params) SetXCassandraToken(xCassandraToken string) {
	o.XCassandraToken = xCassandraToken
}

// WithFields adds the fields to the get rows 1 params
func (o *GetRows1Params) WithFields(fields *string) *GetRows1Params {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get rows 1 params
func (o *GetRows1Params) SetFields(fields *string) {
	o.Fields = fields
}

// WithKeyspaceName adds the keyspaceName to the get rows 1 params
func (o *GetRows1Params) WithKeyspaceName(keyspaceName string) *GetRows1Params {
	o.SetKeyspaceName(keyspaceName)
	return o
}

// SetKeyspaceName adds the keyspaceName to the get rows 1 params
func (o *GetRows1Params) SetKeyspaceName(keyspaceName string) {
	o.KeyspaceName = keyspaceName
}

// WithPageSize adds the pageSize to the get rows 1 params
func (o *GetRows1Params) WithPageSize(pageSize *int32) *GetRows1Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get rows 1 params
func (o *GetRows1Params) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPageState adds the pageState to the get rows 1 params
func (o *GetRows1Params) WithPageState(pageState *string) *GetRows1Params {
	o.SetPageState(pageState)
	return o
}

// SetPageState adds the pageState to the get rows 1 params
func (o *GetRows1Params) SetPageState(pageState *string) {
	o.PageState = pageState
}

// WithPrimaryKey adds the primaryKey to the get rows 1 params
func (o *GetRows1Params) WithPrimaryKey(primaryKey []string) *GetRows1Params {
	o.SetPrimaryKey(primaryKey)
	return o
}

// SetPrimaryKey adds the primaryKey to the get rows 1 params
func (o *GetRows1Params) SetPrimaryKey(primaryKey []string) {
	o.PrimaryKey = primaryKey
}

// WithRaw adds the raw to the get rows 1 params
func (o *GetRows1Params) WithRaw(raw *bool) *GetRows1Params {
	o.SetRaw(raw)
	return o
}

// SetRaw adds the raw to the get rows 1 params
func (o *GetRows1Params) SetRaw(raw *bool) {
	o.Raw = raw
}

// WithSort adds the sort to the get rows 1 params
func (o *GetRows1Params) WithSort(sort *string) *GetRows1Params {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get rows 1 params
func (o *GetRows1Params) SetSort(sort *string) {
	o.Sort = sort
}

// WithTableName adds the tableName to the get rows 1 params
func (o *GetRows1Params) WithTableName(tableName string) *GetRows1Params {
	o.SetTableName(tableName)
	return o
}

// SetTableName adds the tableName to the get rows 1 params
func (o *GetRows1Params) SetTableName(tableName string) {
	o.TableName = tableName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRows1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Cassandra-Token
	if err := r.SetHeaderParam("X-Cassandra-Token", o.XCassandraToken); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param keyspaceName
	if err := r.SetPathParam("keyspaceName", o.KeyspaceName); err != nil {
		return err
	}

	if o.PageSize != nil {

		// query param page-size
		var qrPageSize int32
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page-size", qPageSize); err != nil {
				return err
			}
		}

	}

	if o.PageState != nil {

		// query param page-state
		var qrPageState string
		if o.PageState != nil {
			qrPageState = *o.PageState
		}
		qPageState := qrPageState
		if qPageState != "" {
			if err := r.SetQueryParam("page-state", qPageState); err != nil {
				return err
			}
		}

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

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
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
