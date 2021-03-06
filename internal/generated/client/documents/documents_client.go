// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new documents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for documents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateCollection(params *CreateCollectionParams) (*CreateCollectionCreated, error)

	CreateNamespace(params *CreateNamespaceParams) (*CreateNamespaceCreated, error)

	DeleteCollection(params *DeleteCollectionParams) (*DeleteCollectionNoContent, error)

	DeleteDoc(params *DeleteDocParams) (*DeleteDocNoContent, error)

	DeleteDocPath(params *DeleteDocPathParams) (*DeleteDocPathNoContent, error)

	DeleteNamespace(params *DeleteNamespaceParams) (*DeleteNamespaceNoContent, error)

	GetAllNamespaces(params *GetAllNamespacesParams) (*GetAllNamespacesOK, error)

	GetCollections(params *GetCollectionsParams) (*GetCollectionsOK, error)

	GetDoc(params *GetDocParams) (*GetDocOK, *GetDocNoContent, error)

	GetDocPath(params *GetDocPathParams) (*GetDocPathOK, *GetDocPathNoContent, error)

	GetOneNamespace(params *GetOneNamespaceParams) (*GetOneNamespaceOK, error)

	PatchDoc(params *PatchDocParams) (*PatchDocOK, error)

	PatchDocPath(params *PatchDocPathParams) (*PatchDocPathOK, error)

	PostDoc(params *PostDocParams) (*PostDocCreated, error)

	PutDoc(params *PutDocParams) (*PutDocOK, error)

	PutDocPath(params *PutDocPathParams) (*PutDocPathOK, error)

	SearchDoc(params *SearchDocParams) (*SearchDocOK, *SearchDocNoContent, error)

	UpgradeCollection(params *UpgradeCollectionParams) (*UpgradeCollectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateCollection creates a new empty collection in a namespace
*/
func (a *Client) CreateCollection(params *CreateCollectionParams) (*CreateCollectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCollection",
		Method:             "POST",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCollectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateNamespace creates a namespace

  Create a new namespace.
*/
func (a *Client) CreateNamespace(params *CreateNamespaceParams) (*CreateNamespaceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNamespace",
		Method:             "POST",
		PathPattern:        "/v2/schemas/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNamespaceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createNamespace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteCollection deletes a collection in a namespace
*/
func (a *Client) DeleteCollection(params *DeleteCollectionParams) (*DeleteCollectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCollection",
		Method:             "DELETE",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCollectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDoc deletes a document

  Delete a document
*/
func (a *Client) DeleteDoc(params *DeleteDocParams) (*DeleteDocNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDoc",
		Method:             "DELETE",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDocNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteDocPath deletes a path in a document

  Delete a path in a document
*/
func (a *Client) DeleteDocPath(params *DeleteDocPathParams) (*DeleteDocPathNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDocPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDocPath",
		Method:             "DELETE",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}/{document-path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDocPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDocPathNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDocPath: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNamespace deletes a namespace

  Delete a single namespace.
*/
func (a *Client) DeleteNamespace(params *DeleteNamespaceParams) (*DeleteNamespaceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNamespace",
		Method:             "DELETE",
		PathPattern:        "/v2/schemas/namespaces/{namespace-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNamespaceNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNamespace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllNamespaces gets all namespaces

  Retrieve all available namespaces.
*/
func (a *Client) GetAllNamespaces(params *GetAllNamespacesParams) (*GetAllNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllNamespacesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllNamespaces",
		Method:             "GET",
		PathPattern:        "/v2/schemas/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllNamespacesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCollections lists collections in namespace
*/
func (a *Client) GetCollections(params *GetCollectionsParams) (*GetCollectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCollectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCollections",
		Method:             "GET",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCollectionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCollectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCollections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDoc gets a document

  Retrieve the JSON representation of the document
*/
func (a *Client) GetDoc(params *GetDocParams) (*GetDocOK, *GetDocNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDoc",
		Method:             "GET",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetDocOK:
		return value, nil, nil
	case *GetDocNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for documents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDocPath gets a path in a document

  Retrieve the JSON representation of the document at a provided path, with optional search parameters.
*/
func (a *Client) GetDocPath(params *GetDocPathParams) (*GetDocPathOK, *GetDocPathNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDocPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDocPath",
		Method:             "GET",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}/{document-path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDocPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GetDocPathOK:
		return value, nil, nil
	case *GetDocPathNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for documents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetOneNamespace gets a namespace

  Return a single namespace specification.
*/
func (a *Client) GetOneNamespace(params *GetOneNamespaceParams) (*GetOneNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOneNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOneNamespace",
		Method:             "GET",
		PathPattern:        "/v2/schemas/namespaces/{namespace-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetOneNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOneNamespaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOneNamespace: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchDoc updates data at the root of a document

  Merges data at the root with requested data.
*/
func (a *Client) PatchDoc(params *PatchDocParams) (*PatchDocOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDoc",
		Method:             "PATCH",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchDocOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchDocPath updates data at a path in a document

  Merges data at the path with requested data, assumes that the data at the path is already an object.
*/
func (a *Client) PatchDocPath(params *PatchDocPathParams) (*PatchDocPathOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDocPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDocPath",
		Method:             "PATCH",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}/{document-path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchDocPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchDocPathOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchDocPath: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostDoc creates a new document

  Auto-generates an ID for the newly created document
*/
func (a *Client) PostDoc(params *PostDocParams) (*PostDocCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDoc",
		Method:             "POST",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDocCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutDoc creates or update a document with the provided document id
*/
func (a *Client) PutDoc(params *PutDocParams) (*PutDocOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putDoc",
		Method:             "PUT",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutDocOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutDocPath replaces data at a path in a document

  Removes whatever was previously present at the path
*/
func (a *Client) PutDocPath(params *PutDocPathParams) (*PutDocPathOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDocPathParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putDocPath",
		Method:             "PUT",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/{document-id}/{document-path}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutDocPathReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutDocPathOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putDocPath: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SearchDoc searches documents in a collection

  Page over documents in a collection, with optional search parameters. Does not perform well for large documents.
*/
func (a *Client) SearchDoc(params *SearchDocParams) (*SearchDocOK, *SearchDocNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchDoc",
		Method:             "GET",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *SearchDocOK:
		return value, nil, nil
	case *SearchDocNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for documents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpgradeCollection upgrades a collection in a namespace

  WARNING: This endpoint is expected to cause some down-time for the collection you choose.
*/
func (a *Client) UpgradeCollection(params *UpgradeCollectionParams) (*UpgradeCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeCollection",
		Method:             "POST",
		PathPattern:        "/v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpgradeCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpgradeCollectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upgradeCollection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
