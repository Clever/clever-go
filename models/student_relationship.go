// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StudentRelationship student relationship
//
// swagger:model StudentRelationship
type StudentRelationship struct {

	// relationship
	// Enum: [Parent Grandparent Self Aunt/Uncle Sibling Other ]
	Relationship *string `json:"relationship,omitempty"`

	// student
	Student string `json:"student,omitempty"`

	// type
	// Enum: [Parent/Guardian Emergency Primary Secondary Family Other ]
	Type *string `json:"type,omitempty"`
}

// Validate validates this student relationship
func (m *StudentRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelationship(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var studentRelationshipTypeRelationshipPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Parent","Grandparent","Self","Aunt/Uncle","Sibling","Other",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentRelationshipTypeRelationshipPropEnum = append(studentRelationshipTypeRelationshipPropEnum, v)
	}
}

const (

	// StudentRelationshipRelationshipParent captures enum value "Parent"
	StudentRelationshipRelationshipParent string = "Parent"

	// StudentRelationshipRelationshipGrandparent captures enum value "Grandparent"
	StudentRelationshipRelationshipGrandparent string = "Grandparent"

	// StudentRelationshipRelationshipSelf captures enum value "Self"
	StudentRelationshipRelationshipSelf string = "Self"

	// StudentRelationshipRelationshipAuntUncle captures enum value "Aunt/Uncle"
	StudentRelationshipRelationshipAuntUncle string = "Aunt/Uncle"

	// StudentRelationshipRelationshipSibling captures enum value "Sibling"
	StudentRelationshipRelationshipSibling string = "Sibling"

	// StudentRelationshipRelationshipOther captures enum value "Other"
	StudentRelationshipRelationshipOther string = "Other"

	// StudentRelationshipRelationshipEmpty captures enum value ""
	StudentRelationshipRelationshipEmpty string = ""
)

// prop value enum
func (m *StudentRelationship) validateRelationshipEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentRelationshipTypeRelationshipPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StudentRelationship) validateRelationship(formats strfmt.Registry) error {
	if swag.IsZero(m.Relationship) { // not required
		return nil
	}

	// value enum
	if err := m.validateRelationshipEnum("relationship", "body", *m.Relationship); err != nil {
		return err
	}

	return nil
}

var studentRelationshipTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Parent/Guardian","Emergency","Primary","Secondary","Family","Other",""]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		studentRelationshipTypeTypePropEnum = append(studentRelationshipTypeTypePropEnum, v)
	}
}

const (

	// StudentRelationshipTypeParentGuardian captures enum value "Parent/Guardian"
	StudentRelationshipTypeParentGuardian string = "Parent/Guardian"

	// StudentRelationshipTypeEmergency captures enum value "Emergency"
	StudentRelationshipTypeEmergency string = "Emergency"

	// StudentRelationshipTypePrimary captures enum value "Primary"
	StudentRelationshipTypePrimary string = "Primary"

	// StudentRelationshipTypeSecondary captures enum value "Secondary"
	StudentRelationshipTypeSecondary string = "Secondary"

	// StudentRelationshipTypeFamily captures enum value "Family"
	StudentRelationshipTypeFamily string = "Family"

	// StudentRelationshipTypeOther captures enum value "Other"
	StudentRelationshipTypeOther string = "Other"

	// StudentRelationshipTypeEmpty captures enum value ""
	StudentRelationshipTypeEmpty string = ""
)

// prop value enum
func (m *StudentRelationship) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, studentRelationshipTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StudentRelationship) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this student relationship based on context it is used
func (m *StudentRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StudentRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StudentRelationship) UnmarshalBinary(b []byte) error {
	var res StudentRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}