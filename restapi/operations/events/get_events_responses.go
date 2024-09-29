// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetEventsOKCode is the HTTP code returned for type GetEventsOK
const GetEventsOKCode int = 200

/*
GetEventsOK OK Response

swagger:response getEventsOK
*/
type GetEventsOK struct {

	/*
	  In: Body
	*/
	Payload *models.EventsResponse `json:"body,omitempty"`
}

// NewGetEventsOK creates GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {

	return &GetEventsOK{}
}

// WithPayload adds the payload to the get events o k response
func (o *GetEventsOK) WithPayload(payload *models.EventsResponse) *GetEventsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events o k response
func (o *GetEventsOK) SetPayload(payload *models.EventsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventsNotFoundCode is the HTTP code returned for type GetEventsNotFound
const GetEventsNotFoundCode int = 404

/*
GetEventsNotFound Entity Not Found

swagger:response getEventsNotFound
*/
type GetEventsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetEventsNotFound creates GetEventsNotFound with default headers values
func NewGetEventsNotFound() *GetEventsNotFound {

	return &GetEventsNotFound{}
}

// WithPayload adds the payload to the get events not found response
func (o *GetEventsNotFound) WithPayload(payload *models.NotFound) *GetEventsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get events not found response
func (o *GetEventsNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
