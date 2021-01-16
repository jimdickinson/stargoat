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

// CreateTableReader is a Reader for the CreateTable structure.
type CreateTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateTableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateTableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTableCreated creates a CreateTableCreated with default headers values
func NewCreateTableCreated() *CreateTableCreated {
	return &CreateTableCreated{}
}

/*CreateTableCreated handles this case with default header values.

Created
*/
type CreateTableCreated struct {
	Payload map[string]interface{}
}

func (o *CreateTableCreated) Error() string {
	return fmt.Sprintf("[POST /v2/schemas/keyspaces/{keyspaceName}/tables][%d] createTableCreated  %+v", 201, o.Payload)
}

func (o *CreateTableCreated) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *CreateTableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableBadRequest creates a CreateTableBadRequest with default headers values
func NewCreateTableBadRequest() *CreateTableBadRequest {
	return &CreateTableBadRequest{}
}

/*CreateTableBadRequest handles this case with default header values.

Bad Request
*/
type CreateTableBadRequest struct {
	Payload *models.Error
}

func (o *CreateTableBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/schemas/keyspaces/{keyspaceName}/tables][%d] createTableBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTableBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableUnauthorized creates a CreateTableUnauthorized with default headers values
func NewCreateTableUnauthorized() *CreateTableUnauthorized {
	return &CreateTableUnauthorized{}
}

/*CreateTableUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateTableUnauthorized struct {
	Payload *models.Error
}

func (o *CreateTableUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/schemas/keyspaces/{keyspaceName}/tables][%d] createTableUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTableUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableConflict creates a CreateTableConflict with default headers values
func NewCreateTableConflict() *CreateTableConflict {
	return &CreateTableConflict{}
}

/*CreateTableConflict handles this case with default header values.

Conflict
*/
type CreateTableConflict struct {
	Payload *models.Error
}

func (o *CreateTableConflict) Error() string {
	return fmt.Sprintf("[POST /v2/schemas/keyspaces/{keyspaceName}/tables][%d] createTableConflict  %+v", 409, o.Payload)
}

func (o *CreateTableConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTableInternalServerError creates a CreateTableInternalServerError with default headers values
func NewCreateTableInternalServerError() *CreateTableInternalServerError {
	return &CreateTableInternalServerError{}
}

/*CreateTableInternalServerError handles this case with default header values.

Internal server error
*/
type CreateTableInternalServerError struct {
	Payload *models.Error
}

func (o *CreateTableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/schemas/keyspaces/{keyspaceName}/tables][%d] createTableInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTableInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateTableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
