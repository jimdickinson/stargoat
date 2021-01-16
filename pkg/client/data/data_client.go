// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddRow(params *AddRowParams) (*AddRowCreated, error)

	CreateRow(params *CreateRowParams) (*CreateRowCreated, error)

	DeleteRow(params *DeleteRowParams) (*DeleteRowNoContent, error)

	DeleteRows(params *DeleteRowsParams) (*DeleteRowsNoContent, error)

	GetAllRows(params *GetAllRowsParams) (*GetAllRowsOK, error)

	GetRowWithWhere(params *GetRowWithWhereParams) (*GetRowWithWhereOK, error)

	GetRows(params *GetRowsParams) (*GetRowsOK, error)

	GetRows1(params *GetRows1Params) (*GetRows1OK, error)

	PatchRows(params *PatchRowsParams) (*PatchRowsOK, error)

	QueryRows(params *QueryRowsParams) (*QueryRowsOK, error)

	UpdateRow(params *UpdateRowParams) (*UpdateRowOK, error)

	UpdateRows(params *UpdateRowsParams) (*UpdateRowsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddRow adds row

  Add a row to a table in your database. If the new row has the same primary key as that of an existing row, the database processes it as an update to the existing row.
*/
func (a *Client) AddRow(params *AddRowParams) (*AddRowCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRow",
		Method:             "POST",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddRowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRowCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addRow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateRow adds row

  Add a row to a table in your database. If the new row has the same primary key as that of an existing row, the database processes it as an update to the existing row.
*/
func (a *Client) CreateRow(params *CreateRowParams) (*CreateRowCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRow",
		Method:             "POST",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRowCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRow deletes rows

  Delete individual rows from a table.
*/
func (a *Client) DeleteRow(params *DeleteRowParams) (*DeleteRowNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRow",
		Method:             "DELETE",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRowNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteRows deletes row s

  Delete one or more rows in a table
*/
func (a *Client) DeleteRows(params *DeleteRowsParams) (*DeleteRowsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRows",
		Method:             "DELETE",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRowsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllRows retrieves all rows

  Get all rows from a table.
*/
func (a *Client) GetAllRows(params *GetAllRowsParams) (*GetAllRowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllRows",
		Method:             "GET",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllRowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRowWithWhere searches a table

  Search a table using a json query as defined in the `where` query parameter
*/
func (a *Client) GetRowWithWhere(params *GetRowWithWhereParams) (*GetRowWithWhereOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRowWithWhereParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRowWithWhere",
		Method:             "GET",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRowWithWhereReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRowWithWhereOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRowWithWhere: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRows retrieves rows

  Get rows from a table based on the primary key.
*/
func (a *Client) GetRows(params *GetRowsParams) (*GetRowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRows",
		Method:             "GET",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRows1 gets row s

  Get rows from a table based on the primary key.
*/
func (a *Client) GetRows1(params *GetRows1Params) (*GetRows1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRows1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRows_1",
		Method:             "GET",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRows1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRows1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getRows_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchRows updates part of a row s

  Perform a partial update of one or more rows in a table
*/
func (a *Client) PatchRows(params *PatchRowsParams) (*PatchRowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchRows",
		Method:             "PATCH",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchRowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  QueryRows submits queries

  Submit queries to retrieve data from a table.
*/
func (a *Client) QueryRows(params *QueryRowsParams) (*QueryRowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryRows",
		Method:             "POST",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &QueryRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryRowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRow updates rows

  Update existing rows in a table.
*/
func (a *Client) UpdateRow(params *UpdateRowParams) (*UpdateRowOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRowParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRow",
		Method:             "PUT",
		PathPattern:        "/v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRowReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRowOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRow: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateRows replaces row s

  Update existing rows in a table.
*/
func (a *Client) UpdateRows(params *UpdateRowsParams) (*UpdateRowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRows",
		Method:             "PUT",
		PathPattern:        "/v2/keyspaces/{keyspaceName}/{tableName}/{primaryKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateRowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateRows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}