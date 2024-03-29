// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Staff staff
//
// swagger:model Staff
type Staff struct {

	// credentials
	Credentials *Credentials `json:"credentials,omitempty"`

	// department
	Department *string `json:"department,omitempty"`

	// ext
	Ext interface{} `json:"ext,omitempty"`

	// legacy id
	LegacyID string `json:"legacy_id,omitempty"`

	// roles
	Roles []string `json:"roles"`

	// schools
	Schools []string `json:"schools"`

	// staff id
	StaffID string `json:"staff_id,omitempty"`

	// title
	Title *string `json:"title,omitempty"`
}

// Validate validates this staff
func (m *Staff) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Staff) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this staff based on the context it is used
func (m *Staff) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Staff) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Staff) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Staff) UnmarshalBinary(b []byte) error {
	var res Staff
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
