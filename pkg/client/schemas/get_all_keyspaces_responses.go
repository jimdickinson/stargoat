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

// GetAllKeyspacesReader is a Reader for the GetAllKeyspaces structure.
type GetAllKeyspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllKeyspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllKeyspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllKeyspacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllKeyspacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllKeyspacesOK creates a GetAllKeyspacesOK with default headers values
func NewGetAllKeyspacesOK() *GetAllKeyspacesOK {
	return &GetAllKeyspacesOK{}
}

/*GetAllKeyspacesOK handles this case with default header values.

OK
*/
type GetAllKeyspacesOK struct {
	Payload *models.ResponseWrapper
}

func (o *GetAllKeyspacesOK) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces][%d] getAllKeyspacesOK  %+v", 200, o.Payload)
}

func (o *GetAllKeyspacesOK) GetPayload() *models.ResponseWrapper {
	return o.Payload
}

func (o *GetAllKeyspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllKeyspacesUnauthorized creates a GetAllKeyspacesUnauthorized with default headers values
func NewGetAllKeyspacesUnauthorized() *GetAllKeyspacesUnauthorized {
	return &GetAllKeyspacesUnauthorized{}
}

/*GetAllKeyspacesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllKeyspacesUnauthorized struct {
	Payload *models.Error
}

func (o *GetAllKeyspacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces][%d] getAllKeyspacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAllKeyspacesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllKeyspacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllKeyspacesInternalServerError creates a GetAllKeyspacesInternalServerError with default headers values
func NewGetAllKeyspacesInternalServerError() *GetAllKeyspacesInternalServerError {
	return &GetAllKeyspacesInternalServerError{}
}

/*GetAllKeyspacesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAllKeyspacesInternalServerError struct {
	Payload *models.Error
}

func (o *GetAllKeyspacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces][%d] getAllKeyspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllKeyspacesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllKeyspacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
