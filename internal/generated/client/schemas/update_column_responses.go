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

// UpdateColumnReader is a Reader for the UpdateColumn structure.
type UpdateColumnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateColumnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateColumnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateColumnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateColumnUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateColumnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateColumnNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateColumnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateColumnOK creates a UpdateColumnOK with default headers values
func NewUpdateColumnOK() *UpdateColumnOK {
	return &UpdateColumnOK{}
}

/*UpdateColumnOK handles this case with default header values.

OK
*/
type UpdateColumnOK struct {
	Payload *models.SuccessResponse
}

func (o *UpdateColumnOK) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnOK  %+v", 200, o.Payload)
}

func (o *UpdateColumnOK) GetPayload() *models.SuccessResponse {
	return o.Payload
}

func (o *UpdateColumnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColumnBadRequest creates a UpdateColumnBadRequest with default headers values
func NewUpdateColumnBadRequest() *UpdateColumnBadRequest {
	return &UpdateColumnBadRequest{}
}

/*UpdateColumnBadRequest handles this case with default header values.

Bad request
*/
type UpdateColumnBadRequest struct {
	Payload *models.Error
}

func (o *UpdateColumnBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateColumnBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColumnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColumnUnauthorized creates a UpdateColumnUnauthorized with default headers values
func NewUpdateColumnUnauthorized() *UpdateColumnUnauthorized {
	return &UpdateColumnUnauthorized{}
}

/*UpdateColumnUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateColumnUnauthorized struct {
	Payload *models.Error
}

func (o *UpdateColumnUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateColumnUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColumnUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColumnForbidden creates a UpdateColumnForbidden with default headers values
func NewUpdateColumnForbidden() *UpdateColumnForbidden {
	return &UpdateColumnForbidden{}
}

/*UpdateColumnForbidden handles this case with default header values.

Forbidden
*/
type UpdateColumnForbidden struct {
	Payload *models.Error
}

func (o *UpdateColumnForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnForbidden  %+v", 403, o.Payload)
}

func (o *UpdateColumnForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColumnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColumnNotFound creates a UpdateColumnNotFound with default headers values
func NewUpdateColumnNotFound() *UpdateColumnNotFound {
	return &UpdateColumnNotFound{}
}

/*UpdateColumnNotFound handles this case with default header values.

Not Found
*/
type UpdateColumnNotFound struct {
	Payload *models.Error
}

func (o *UpdateColumnNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnNotFound  %+v", 404, o.Payload)
}

func (o *UpdateColumnNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColumnNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateColumnInternalServerError creates a UpdateColumnInternalServerError with default headers values
func NewUpdateColumnInternalServerError() *UpdateColumnInternalServerError {
	return &UpdateColumnInternalServerError{}
}

/*UpdateColumnInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateColumnInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateColumnInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns/{columnName}][%d] updateColumnInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateColumnInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateColumnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
