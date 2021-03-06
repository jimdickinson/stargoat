// Code generated by go-swagger; DO NOT EDIT.

package schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jimdickinson/stargoat/internal/generated/models"
)

// ListAllKeyspacesReader is a Reader for the ListAllKeyspaces structure.
type ListAllKeyspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllKeyspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllKeyspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllKeyspacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAllKeyspacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllKeyspacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAllKeyspacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllKeyspacesOK creates a ListAllKeyspacesOK with default headers values
func NewListAllKeyspacesOK() *ListAllKeyspacesOK {
	return &ListAllKeyspacesOK{}
}

/*ListAllKeyspacesOK handles this case with default header values.

OK
*/
type ListAllKeyspacesOK struct {
	Payload []string
}

func (o *ListAllKeyspacesOK) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces][%d] listAllKeyspacesOK  %+v", 200, o.Payload)
}

func (o *ListAllKeyspacesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListAllKeyspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllKeyspacesBadRequest creates a ListAllKeyspacesBadRequest with default headers values
func NewListAllKeyspacesBadRequest() *ListAllKeyspacesBadRequest {
	return &ListAllKeyspacesBadRequest{}
}

/*ListAllKeyspacesBadRequest handles this case with default header values.

Bad request
*/
type ListAllKeyspacesBadRequest struct {
	Payload *models.Error
}

func (o *ListAllKeyspacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces][%d] listAllKeyspacesBadRequest  %+v", 400, o.Payload)
}

func (o *ListAllKeyspacesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllKeyspacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllKeyspacesUnauthorized creates a ListAllKeyspacesUnauthorized with default headers values
func NewListAllKeyspacesUnauthorized() *ListAllKeyspacesUnauthorized {
	return &ListAllKeyspacesUnauthorized{}
}

/*ListAllKeyspacesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllKeyspacesUnauthorized struct {
	Payload *models.Error
}

func (o *ListAllKeyspacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces][%d] listAllKeyspacesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllKeyspacesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllKeyspacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllKeyspacesForbidden creates a ListAllKeyspacesForbidden with default headers values
func NewListAllKeyspacesForbidden() *ListAllKeyspacesForbidden {
	return &ListAllKeyspacesForbidden{}
}

/*ListAllKeyspacesForbidden handles this case with default header values.

Forbidden
*/
type ListAllKeyspacesForbidden struct {
	Payload *models.Error
}

func (o *ListAllKeyspacesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces][%d] listAllKeyspacesForbidden  %+v", 403, o.Payload)
}

func (o *ListAllKeyspacesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllKeyspacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllKeyspacesInternalServerError creates a ListAllKeyspacesInternalServerError with default headers values
func NewListAllKeyspacesInternalServerError() *ListAllKeyspacesInternalServerError {
	return &ListAllKeyspacesInternalServerError{}
}

/*ListAllKeyspacesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllKeyspacesInternalServerError struct {
	Payload *models.Error
}

func (o *ListAllKeyspacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces][%d] listAllKeyspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllKeyspacesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllKeyspacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
