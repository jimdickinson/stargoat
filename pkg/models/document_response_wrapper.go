// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DocumentResponseWrapper document response wrapper
//
// swagger:model DocumentResponseWrapper
type DocumentResponseWrapper struct {

	// The data returned by the request
	Data interface{} `json:"data,omitempty"`

	// The id of the document
	DocumentID string `json:"documentId,omitempty"`

	// A string representing the paging state to be used on future paging requests.
	PageState string `json:"pageState,omitempty"`
}

// Validate validates this document response wrapper
func (m *DocumentResponseWrapper) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentResponseWrapper) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentResponseWrapper) UnmarshalBinary(b []byte) error {
	var res DocumentResponseWrapper
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}