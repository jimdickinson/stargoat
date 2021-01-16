// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jimdickinson/stargoat/pkg/models"
)

// DeleteRowReader is a Reader for the DeleteRow structure.
type DeleteRowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRowNoContent creates a DeleteRowNoContent with default headers values
func NewDeleteRowNoContent() *DeleteRowNoContent {
	return &DeleteRowNoContent{}
}

/*DeleteRowNoContent handles this case with default header values.

No Content
*/
type DeleteRowNoContent struct {
}

func (o *DeleteRowNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}][%d] deleteRowNoContent ", 204)
}

func (o *DeleteRowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRowBadRequest creates a DeleteRowBadRequest with default headers values
func NewDeleteRowBadRequest() *DeleteRowBadRequest {
	return &DeleteRowBadRequest{}
}

/*DeleteRowBadRequest handles this case with default header values.

Bad request
*/
type DeleteRowBadRequest struct {
	Payload *models.Error
}

func (o *DeleteRowBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}][%d] deleteRowBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRowBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRowUnauthorized creates a DeleteRowUnauthorized with default headers values
func NewDeleteRowUnauthorized() *DeleteRowUnauthorized {
	return &DeleteRowUnauthorized{}
}

/*DeleteRowUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteRowUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteRowUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}][%d] deleteRowUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRowUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRowForbidden creates a DeleteRowForbidden with default headers values
func NewDeleteRowForbidden() *DeleteRowForbidden {
	return &DeleteRowForbidden{}
}

/*DeleteRowForbidden handles this case with default header values.

Forbidden
*/
type DeleteRowForbidden struct {
	Payload *models.Error
}

func (o *DeleteRowForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}][%d] deleteRowForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRowForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRowInternalServerError creates a DeleteRowInternalServerError with default headers values
func NewDeleteRowInternalServerError() *DeleteRowInternalServerError {
	return &DeleteRowInternalServerError{}
}

/*DeleteRowInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteRowInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteRowInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/keyspaces/{keyspaceName}/tables/{tableName}/rows/{primaryKey}][%d] deleteRowInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRowInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteRowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
