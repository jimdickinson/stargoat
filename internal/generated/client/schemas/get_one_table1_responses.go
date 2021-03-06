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

// GetOneTable1Reader is a Reader for the GetOneTable1 structure.
type GetOneTable1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOneTable1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOneTable1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOneTable1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOneTable1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOneTable1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOneTable1OK creates a GetOneTable1OK with default headers values
func NewGetOneTable1OK() *GetOneTable1OK {
	return &GetOneTable1OK{}
}

/*GetOneTable1OK handles this case with default header values.

OK
*/
type GetOneTable1OK struct {
	Payload *models.Table
}

func (o *GetOneTable1OK) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}][%d] getOneTable1OK  %+v", 200, o.Payload)
}

func (o *GetOneTable1OK) GetPayload() *models.Table {
	return o.Payload
}

func (o *GetOneTable1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Table)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneTable1Unauthorized creates a GetOneTable1Unauthorized with default headers values
func NewGetOneTable1Unauthorized() *GetOneTable1Unauthorized {
	return &GetOneTable1Unauthorized{}
}

/*GetOneTable1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetOneTable1Unauthorized struct {
	Payload *models.Error
}

func (o *GetOneTable1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}][%d] getOneTable1Unauthorized  %+v", 401, o.Payload)
}

func (o *GetOneTable1Unauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneTable1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneTable1NotFound creates a GetOneTable1NotFound with default headers values
func NewGetOneTable1NotFound() *GetOneTable1NotFound {
	return &GetOneTable1NotFound{}
}

/*GetOneTable1NotFound handles this case with default header values.

Not Found
*/
type GetOneTable1NotFound struct {
	Payload *models.Error
}

func (o *GetOneTable1NotFound) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}][%d] getOneTable1NotFound  %+v", 404, o.Payload)
}

func (o *GetOneTable1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneTable1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOneTable1InternalServerError creates a GetOneTable1InternalServerError with default headers values
func NewGetOneTable1InternalServerError() *GetOneTable1InternalServerError {
	return &GetOneTable1InternalServerError{}
}

/*GetOneTable1InternalServerError handles this case with default header values.

Internal server error
*/
type GetOneTable1InternalServerError struct {
	Payload *models.Error
}

func (o *GetOneTable1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}][%d] getOneTable1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetOneTable1InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOneTable1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
