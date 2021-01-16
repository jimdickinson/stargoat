// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Datacenter datacenter
//
// swagger:model Datacenter
type Datacenter struct {

	// The name of the datacenter.
	// Required: true
	Name *string `json:"name"`

	// The number of replicas in the datacenter. In other words, the number of copies of each row in the datacenter.
	// Required: true
	Replicas *int32 `json:"replicas"`
}

// Validate validates this datacenter
func (m *Datacenter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Datacenter) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Datacenter) validateReplicas(formats strfmt.Registry) error {

	if err := validate.Required("replicas", "body", m.Replicas); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Datacenter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Datacenter) UnmarshalBinary(b []byte) error {
	var res Datacenter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
