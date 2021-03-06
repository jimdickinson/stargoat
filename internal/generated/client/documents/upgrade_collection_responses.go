// Code generated by go-swagger; DO NOT EDIT.

package documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/jimdickinson/stargoat/internal/generated/models"
)

// UpgradeCollectionReader is a Reader for the UpgradeCollection structure.
type UpgradeCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpgradeCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpgradeCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpgradeCollectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpgradeCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpgradeCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpgradeCollectionOK creates a UpgradeCollectionOK with default headers values
func NewUpgradeCollectionOK() *UpgradeCollectionOK {
	return &UpgradeCollectionOK{}
}

/*UpgradeCollectionOK handles this case with default header values.

OK
*/
type UpgradeCollectionOK struct {
	Payload *models.ResponseWrapper
}

func (o *UpgradeCollectionOK) Error() string {
	return fmt.Sprintf("[POST /v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade][%d] upgradeCollectionOK  %+v", 200, o.Payload)
}

func (o *UpgradeCollectionOK) GetPayload() *models.ResponseWrapper {
	return o.Payload
}

func (o *UpgradeCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpgradeCollectionBadRequest creates a UpgradeCollectionBadRequest with default headers values
func NewUpgradeCollectionBadRequest() *UpgradeCollectionBadRequest {
	return &UpgradeCollectionBadRequest{}
}

/*UpgradeCollectionBadRequest handles this case with default header values.

Bad Request
*/
type UpgradeCollectionBadRequest struct {
}

func (o *UpgradeCollectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade][%d] upgradeCollectionBadRequest ", 400)
}

func (o *UpgradeCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeCollectionUnauthorized creates a UpgradeCollectionUnauthorized with default headers values
func NewUpgradeCollectionUnauthorized() *UpgradeCollectionUnauthorized {
	return &UpgradeCollectionUnauthorized{}
}

/*UpgradeCollectionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpgradeCollectionUnauthorized struct {
}

func (o *UpgradeCollectionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade][%d] upgradeCollectionUnauthorized ", 401)
}

func (o *UpgradeCollectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeCollectionNotFound creates a UpgradeCollectionNotFound with default headers values
func NewUpgradeCollectionNotFound() *UpgradeCollectionNotFound {
	return &UpgradeCollectionNotFound{}
}

/*UpgradeCollectionNotFound handles this case with default header values.

Collection not found
*/
type UpgradeCollectionNotFound struct {
}

func (o *UpgradeCollectionNotFound) Error() string {
	return fmt.Sprintf("[POST /v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade][%d] upgradeCollectionNotFound ", 404)
}

func (o *UpgradeCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpgradeCollectionInternalServerError creates a UpgradeCollectionInternalServerError with default headers values
func NewUpgradeCollectionInternalServerError() *UpgradeCollectionInternalServerError {
	return &UpgradeCollectionInternalServerError{}
}

/*UpgradeCollectionInternalServerError handles this case with default header values.

Internal server error
*/
type UpgradeCollectionInternalServerError struct {
	Payload *models.Error
}

func (o *UpgradeCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/namespaces/{namespace-id}/collections/{collection-id}/upgrade][%d] upgradeCollectionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpgradeCollectionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpgradeCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
