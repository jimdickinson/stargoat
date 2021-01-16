// Code generated by go-swagger; DO NOT EDIT.

package schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jimdickinson/stargoat/pkg/models"
)

// ListAllTablesReader is a Reader for the ListAllTables structure.
type ListAllTablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllTablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllTablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllTablesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAllTablesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllTablesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAllTablesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllTablesOK creates a ListAllTablesOK with default headers values
func NewListAllTablesOK() *ListAllTablesOK {
	return &ListAllTablesOK{}
}

/*ListAllTablesOK handles this case with default header values.

OK
*/
type ListAllTablesOK struct {
	Payload []string
}

func (o *ListAllTablesOK) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables][%d] listAllTablesOK  %+v", 200, o.Payload)
}

func (o *ListAllTablesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListAllTablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllTablesBadRequest creates a ListAllTablesBadRequest with default headers values
func NewListAllTablesBadRequest() *ListAllTablesBadRequest {
	return &ListAllTablesBadRequest{}
}

/*ListAllTablesBadRequest handles this case with default header values.

Bad request
*/
type ListAllTablesBadRequest struct {
	Payload *models.Error
}

func (o *ListAllTablesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables][%d] listAllTablesBadRequest  %+v", 400, o.Payload)
}

func (o *ListAllTablesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllTablesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllTablesUnauthorized creates a ListAllTablesUnauthorized with default headers values
func NewListAllTablesUnauthorized() *ListAllTablesUnauthorized {
	return &ListAllTablesUnauthorized{}
}

/*ListAllTablesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllTablesUnauthorized struct {
	Payload *models.Error
}

func (o *ListAllTablesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables][%d] listAllTablesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllTablesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllTablesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllTablesForbidden creates a ListAllTablesForbidden with default headers values
func NewListAllTablesForbidden() *ListAllTablesForbidden {
	return &ListAllTablesForbidden{}
}

/*ListAllTablesForbidden handles this case with default header values.

Forbidden
*/
type ListAllTablesForbidden struct {
	Payload *models.Error
}

func (o *ListAllTablesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables][%d] listAllTablesForbidden  %+v", 403, o.Payload)
}

func (o *ListAllTablesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllTablesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllTablesInternalServerError creates a ListAllTablesInternalServerError with default headers values
func NewListAllTablesInternalServerError() *ListAllTablesInternalServerError {
	return &ListAllTablesInternalServerError{}
}

/*ListAllTablesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllTablesInternalServerError struct {
	Payload *models.Error
}

func (o *ListAllTablesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables][%d] listAllTablesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllTablesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllTablesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}