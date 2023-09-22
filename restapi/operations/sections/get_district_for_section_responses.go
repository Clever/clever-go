// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetDistrictForSectionOKCode is the HTTP code returned for type GetDistrictForSectionOK
const GetDistrictForSectionOKCode int = 200

/*
GetDistrictForSectionOK OK Response

swagger:response getDistrictForSectionOK
*/
type GetDistrictForSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.DistrictResponse `json:"body,omitempty"`
}

// NewGetDistrictForSectionOK creates GetDistrictForSectionOK with default headers values
func NewGetDistrictForSectionOK() *GetDistrictForSectionOK {

	return &GetDistrictForSectionOK{}
}

// WithPayload adds the payload to the get district for section o k response
func (o *GetDistrictForSectionOK) WithPayload(payload *models.DistrictResponse) *GetDistrictForSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for section o k response
func (o *GetDistrictForSectionOK) SetPayload(payload *models.DistrictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDistrictForSectionNotFoundCode is the HTTP code returned for type GetDistrictForSectionNotFound
const GetDistrictForSectionNotFoundCode int = 404

/*
GetDistrictForSectionNotFound Entity Not Found

swagger:response getDistrictForSectionNotFound
*/
type GetDistrictForSectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetDistrictForSectionNotFound creates GetDistrictForSectionNotFound with default headers values
func NewGetDistrictForSectionNotFound() *GetDistrictForSectionNotFound {

	return &GetDistrictForSectionNotFound{}
}

// WithPayload adds the payload to the get district for section not found response
func (o *GetDistrictForSectionNotFound) WithPayload(payload *models.NotFound) *GetDistrictForSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for section not found response
func (o *GetDistrictForSectionNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
