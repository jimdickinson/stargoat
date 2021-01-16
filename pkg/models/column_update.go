// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ColumnUpdate Changes the name of a primary key column and preserves the existing values.
//
// swagger:model ColumnUpdate
type ColumnUpdate struct {

	// The new name of the column.
	NewName string `json:"newName,omitempty"`
}

// Validate validates this column update
func (m *ColumnUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ColumnUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ColumnUpdate) UnmarshalBinary(b []byte) error {
	var res ColumnUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}