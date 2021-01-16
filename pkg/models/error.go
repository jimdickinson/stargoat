// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Error A description of an error state
//
// swagger:model Error
type Error struct {

	// The internal number referencing the error state
	Code int32 `json:"code,omitempty"`

	// A human readable description of the error state
	Description string `json:"description,omitempty"`

	// internal tx Id
	InternalTxID string `json:"internalTxId,omitempty"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}