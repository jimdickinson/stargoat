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

// DeleteColumn1Reader is a Reader for the DeleteColumn1 structure.
type DeleteColumn1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteColumn1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteColumn1NoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteColumn1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteColumn1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteColumn1NoContent creates a DeleteColumn1NoContent with default headers values
func NewDeleteColumn1NoContent() *DeleteColumn1NoContent {
	return &DeleteColumn1NoContent{}
}

/*DeleteColumn1NoContent handles this case with default header values.

No Content
*/
type DeleteColumn1NoContent struct {
}

func (o *DeleteColumn1NoContent) Error() string {
	return fmt.Sprintf("[DELETE /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] deleteColumn1NoContent ", 204)
}

func (o *DeleteColumn1NoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteColumn1Unauthorized creates a DeleteColumn1Unauthorized with default headers values
func NewDeleteColumn1Unauthorized() *DeleteColumn1Unauthorized {
	return &DeleteColumn1Unauthorized{}
}

/*DeleteColumn1Unauthorized handles this case with default header values.

Unauthorized
*/
type DeleteColumn1Unauthorized struct {
	Payload *models.Error
}

func (o *DeleteColumn1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] deleteColumn1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteColumn1Unauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColumn1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteColumn1InternalServerError creates a DeleteColumn1InternalServerError with default headers values
func NewDeleteColumn1InternalServerError() *DeleteColumn1InternalServerError {
	return &DeleteColumn1InternalServerError{}
}

/*DeleteColumn1InternalServerError handles this case with default header values.

Internal server error
*/
type DeleteColumn1InternalServerError struct {
	Payload *models.Error
}

func (o *DeleteColumn1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v2/schemas/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] deleteColumn1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteColumn1InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteColumn1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
