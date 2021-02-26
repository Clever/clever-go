// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetTermForSectionOKCode is the HTTP code returned for type GetTermForSectionOK
const GetTermForSectionOKCode int = 200

/*GetTermForSectionOK OK Response

swagger:response getTermForSectionOK
*/
type GetTermForSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.TermResponse `json:"body,omitempty"`
}

// NewGetTermForSectionOK creates GetTermForSectionOK with default headers values
func NewGetTermForSectionOK() *GetTermForSectionOK {

	return &GetTermForSectionOK{}
}

// WithPayload adds the payload to the get term for section o k response
func (o *GetTermForSectionOK) WithPayload(payload *models.TermResponse) *GetTermForSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get term for section o k response
func (o *GetTermForSectionOK) SetPayload(payload *models.TermResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTermForSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTermForSectionNotFoundCode is the HTTP code returned for type GetTermForSectionNotFound
const GetTermForSectionNotFoundCode int = 404

/*GetTermForSectionNotFound Entity Not Found

swagger:response getTermForSectionNotFound
*/
type GetTermForSectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetTermForSectionNotFound creates GetTermForSectionNotFound with default headers values
func NewGetTermForSectionNotFound() *GetTermForSectionNotFound {

	return &GetTermForSectionNotFound{}
}

// WithPayload adds the payload to the get term for section not found response
func (o *GetTermForSectionNotFound) WithPayload(payload *models.NotFound) *GetTermForSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get term for section not found response
func (o *GetTermForSectionNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTermForSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}