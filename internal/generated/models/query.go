// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Query query
//
// swagger:model Query
type Query struct {

	// A list of column names to return in the result set. An empty array returns all columns.
	ColumnNames []string `json:"columnNames"`

	// An array of filters to return results for, separated by `AND`. For example, `a > 1 AND b != 1`.
	// Required: true
	Filters []*Filter `json:"filters"`

	// The clustering order for rows returned.
	OrderBy *ClusteringExpression `json:"orderBy,omitempty"`

	// The size of the page to return in the result set.
	PageSize int32 `json:"pageSize,omitempty"`

	// A string returned from previous query requests representing the paging state.
	PageState string `json:"pageState,omitempty"`
}

// Validate validates this query
func (m *Query) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Query) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Query) validateOrderBy(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Query) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Query) UnmarshalBinary(b []byte) error {
	var res Query
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
