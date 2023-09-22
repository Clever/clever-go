// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetEventOKCode is the HTTP code returned for type GetEventOK
const GetEventOKCode int = 200

/*
GetEventOK OK Response

swagger:response getEventOK
*/
type GetEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.EventResponse `json:"body,omitempty"`
}

// NewGetEventOK creates GetEventOK with default headers values
func NewGetEventOK() *GetEventOK {

	return &GetEventOK{}
}

// WithPayload adds the payload to the get event o k response
func (o *GetEventOK) WithPayload(payload *models.EventResponse) *GetEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event o k response
func (o *GetEventOK) SetPayload(payload *models.EventResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventNotFoundCode is the HTTP code returned for type GetEventNotFound
const GetEventNotFoundCode int = 404

/*
GetEventNotFound Entity Not Found

swagger:response getEventNotFound
*/
type GetEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetEventNotFound creates GetEventNotFound with default headers values
func NewGetEventNotFound() *GetEventNotFound {

	return &GetEventNotFound{}
}

// WithPayload adds the payload to the get event not found response
func (o *GetEventNotFound) WithPayload(payload *models.NotFound) *GetEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event not found response
func (o *GetEventNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
