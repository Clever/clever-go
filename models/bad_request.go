// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BadRequest bad request
//
// swagger:model BadRequest
type BadRequest struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this bad request
func (m *BadRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this bad request based on context it is used
func (m *BadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BadRequest) UnmarshalBinary(b []byte) error {
	var res BadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
