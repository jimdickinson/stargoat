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

// ListAllColumnsReader is a Reader for the ListAllColumns structure.
type ListAllColumnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllColumnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAllColumnsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAllColumnsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAllColumnsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAllColumnsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAllColumnsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListAllColumnsOK creates a ListAllColumnsOK with default headers values
func NewListAllColumnsOK() *ListAllColumnsOK {
	return &ListAllColumnsOK{}
}

/*ListAllColumnsOK handles this case with default header values.

OK
*/
type ListAllColumnsOK struct {
	Payload []*models.ColumnDefinition
}

func (o *ListAllColumnsOK) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns][%d] listAllColumnsOK  %+v", 200, o.Payload)
}

func (o *ListAllColumnsOK) GetPayload() []*models.ColumnDefinition {
	return o.Payload
}

func (o *ListAllColumnsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllColumnsBadRequest creates a ListAllColumnsBadRequest with default headers values
func NewListAllColumnsBadRequest() *ListAllColumnsBadRequest {
	return &ListAllColumnsBadRequest{}
}

/*ListAllColumnsBadRequest handles this case with default header values.

Bad request
*/
type ListAllColumnsBadRequest struct {
	Payload *models.Error
}

func (o *ListAllColumnsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns][%d] listAllColumnsBadRequest  %+v", 400, o.Payload)
}

func (o *ListAllColumnsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllColumnsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllColumnsUnauthorized creates a ListAllColumnsUnauthorized with default headers values
func NewListAllColumnsUnauthorized() *ListAllColumnsUnauthorized {
	return &ListAllColumnsUnauthorized{}
}

/*ListAllColumnsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllColumnsUnauthorized struct {
	Payload *models.Error
}

func (o *ListAllColumnsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns][%d] listAllColumnsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllColumnsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllColumnsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllColumnsForbidden creates a ListAllColumnsForbidden with default headers values
func NewListAllColumnsForbidden() *ListAllColumnsForbidden {
	return &ListAllColumnsForbidden{}
}

/*ListAllColumnsForbidden handles this case with default header values.

Forbidden
*/
type ListAllColumnsForbidden struct {
	Payload *models.Error
}

func (o *ListAllColumnsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns][%d] listAllColumnsForbidden  %+v", 403, o.Payload)
}

func (o *ListAllColumnsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllColumnsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllColumnsInternalServerError creates a ListAllColumnsInternalServerError with default headers values
func NewListAllColumnsInternalServerError() *ListAllColumnsInternalServerError {
	return &ListAllColumnsInternalServerError{}
}

/*ListAllColumnsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllColumnsInternalServerError struct {
	Payload *models.Error
}

func (o *ListAllColumnsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/keyspaces/{keyspaceName}/tables/{tableName}/columns][%d] listAllColumnsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllColumnsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListAllColumnsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
