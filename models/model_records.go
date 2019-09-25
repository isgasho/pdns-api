// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ModelRecords model records
// swagger:model model.Records
type ModelRecords []*ModelRecordsItems0

// Validate validates this model records
func (m ModelRecords) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {
		if swag.IsZero(m[i]) { // not required
			continue
		}

		if m[i] != nil {
			if err := m[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(strconv.Itoa(i))
				}
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ModelRecordsItems0 model records items0
// swagger:model ModelRecordsItems0
type ModelRecordsItems0 struct {

	// auth
	Auth bool `json:"auth,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// db
	Db string `json:"db,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// domain id
	DomainID int64 `json:"domain_id,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// older name
	OlderName string `json:"older_name,omitempty"`

	// prio
	Prio int64 `json:"prio,omitempty"`

	// ttl
	TTL int64 `json:"ttl,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this model records items0
func (m *ModelRecordsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelRecordsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelRecordsItems0) UnmarshalBinary(b []byte) error {
	var res ModelRecordsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}