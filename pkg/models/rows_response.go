// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RowsResponse rows response
//
// swagger:model RowsResponse
type RowsResponse struct {

	// Number of rows modified/added by the request.
	// Read Only: true
	RowsModified int32 `json:"rowsModified,omitempty"`

	// Whether the request was successful.
	// Read Only: true
	Success *bool `json:"success,omitempty"`
}

// Validate validates this rows response
func (m *RowsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RowsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RowsResponse) UnmarshalBinary(b []byte) error {
	var res RowsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}