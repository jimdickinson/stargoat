// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Rows rows
//
// swagger:model Rows
type Rows struct {

	// Number of records being returned by the request.
	// Read Only: true
	Count int32 `json:"count,omitempty"`

	// A string representing the paging state to be used on future paging requests.
	// Read Only: true
	PageState string `json:"pageState,omitempty"`

	// The rows returned by the request.
	// Read Only: true
	Rows []map[string]interface{} `json:"rows"`
}

// Validate validates this rows
func (m *Rows) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Rows) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rows) UnmarshalBinary(b []byte) error {
	var res Rows
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
