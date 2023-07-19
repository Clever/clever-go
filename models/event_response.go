// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EventResponse event response
//
// swagger:model EventResponse
type EventResponse struct {
	dataField Event
}

// Data gets the data of this base type
func (m *EventResponse) Data() Event {
	return m.dataField
}

// SetData sets the data of this base type
func (m *EventResponse) SetData(val Event) {
	m.dataField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *EventResponse) UnmarshalJSON(raw []byte) error {
	var data struct {
		Data json.RawMessage `json:"data,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propData Event
	if string(data.Data) != "null" {
		data, err := UnmarshalEvent(bytes.NewBuffer(data.Data), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propData = data
	}

	var result EventResponse

	// data
	result.dataField = propData

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m EventResponse) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
	}{})
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Data Event `json:"data,omitempty"`
	}{

		Data: m.dataField,
	})
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this event response
func (m *EventResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventResponse) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data()) { // not required
		return nil
	}

	if err := m.Data().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data")
		}
		return err
	}

	return nil
}

// ContextValidate validate this event response based on the context it is used
func (m *EventResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventResponse) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Data()) { // not required
		return nil
	}

	if err := m.Data().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("data")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("data")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventResponse) UnmarshalBinary(b []byte) error {
	var res EventResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}