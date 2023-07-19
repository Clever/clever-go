// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TermsUpdated terms updated
//
// swagger:model terms.updated
type TermsUpdated struct {
	createdField strfmt.DateTime

	idField string

	// data
	Data *TermObject `json:"data,omitempty"`

	// previous attributes
	PreviousAttributes map[string]interface{} `json:"previous_attributes,omitempty"`
}

// Created gets the created of this subtype
func (m *TermsUpdated) Created() strfmt.DateTime {
	return m.createdField
}

// SetCreated sets the created of this subtype
func (m *TermsUpdated) SetCreated(val strfmt.DateTime) {
	m.createdField = val
}

// ID gets the id of this subtype
func (m *TermsUpdated) ID() string {
	return m.idField
}

// SetID sets the id of this subtype
func (m *TermsUpdated) SetID(val string) {
	m.idField = val
}

// Type gets the type of this subtype
func (m *TermsUpdated) Type() string {
	return "terms.updated"
}

// SetType sets the type of this subtype
func (m *TermsUpdated) SetType(val string) {
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *TermsUpdated) UnmarshalJSON(raw []byte) error {
	var data struct {

		// data
		Data *TermObject `json:"data,omitempty"`

		// previous attributes
		PreviousAttributes map[string]interface{} `json:"previous_attributes,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		Created strfmt.DateTime `json:"created,omitempty"`

		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result TermsUpdated

	result.createdField = base.Created

	result.idField = base.ID

	if base.Type != result.Type() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid type value: %q", base.Type)
	}

	result.Data = data.Data
	result.PreviousAttributes = data.PreviousAttributes

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m TermsUpdated) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// data
		Data *TermObject `json:"data,omitempty"`

		// previous attributes
		PreviousAttributes map[string]interface{} `json:"previous_attributes,omitempty"`
	}{

		Data: m.Data,

		PreviousAttributes: m.PreviousAttributes,
	})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Created strfmt.DateTime `json:"created,omitempty"`

		ID string `json:"id,omitempty"`

		Type string `json:"type"`
	}{

		Created: m.Created(),

		ID: m.ID(),

		Type: m.Type(),
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this terms updated
func (m *TermsUpdated) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermsUpdated) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created()) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "datetime", m.Created().String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TermsUpdated) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terms updated based on the context it is used
func (m *TermsUpdated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TermsUpdated) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TermsUpdated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TermsUpdated) UnmarshalBinary(b []byte) error {
	var res TermsUpdated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
