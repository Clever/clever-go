// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Teacher teacher
//
// swagger:model Teacher
type Teacher struct {

	// created
	// Format: datetime
	Created strfmt.DateTime `json:"created,omitempty"`

	// credentials
	Credentials *Credentials `json:"credentials,omitempty"`

	// district
	District string `json:"district,omitempty"`

	// ext
	Ext interface{} `json:"ext,omitempty"`

	// last modified
	// Format: datetime
	LastModified strfmt.DateTime `json:"last_modified,omitempty"`

	// legacy id
	LegacyID string `json:"legacy_id,omitempty"`

	// name
	Name *Name `json:"name,omitempty"`

	// school
	School string `json:"school,omitempty"`

	// schools
	Schools []string `json:"schools"`

	// sis id
	SisID string `json:"sis_id,omitempty"`

	// state id
	StateID *string `json:"state_id,omitempty"`

	// teacher number
	TeacherNumber *string `json:"teacher_number,omitempty"`

	// title
	Title *string `json:"title,omitempty"`
}

// Validate validates this teacher
func (m *Teacher) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Teacher) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "datetime", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Teacher) validateCredentials(formats strfmt.Registry) error {
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

func (m *Teacher) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("last_modified", "body", "datetime", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Teacher) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this teacher based on the context it is used
func (m *Teacher) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Teacher) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Teacher) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if m.Name != nil {
		if err := m.Name.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Teacher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Teacher) UnmarshalBinary(b []byte) error {
	var res Teacher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
