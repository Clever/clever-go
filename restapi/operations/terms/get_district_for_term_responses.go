// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetDistrictForTermOKCode is the HTTP code returned for type GetDistrictForTermOK
const GetDistrictForTermOKCode int = 200

/*
GetDistrictForTermOK OK Response

swagger:response getDistrictForTermOK
*/
type GetDistrictForTermOK struct {

	/*
	  In: Body
	*/
	Payload *models.DistrictResponse `json:"body,omitempty"`
}

// NewGetDistrictForTermOK creates GetDistrictForTermOK with default headers values
func NewGetDistrictForTermOK() *GetDistrictForTermOK {

	return &GetDistrictForTermOK{}
}

// WithPayload adds the payload to the get district for term o k response
func (o *GetDistrictForTermOK) WithPayload(payload *models.DistrictResponse) *GetDistrictForTermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for term o k response
func (o *GetDistrictForTermOK) SetPayload(payload *models.DistrictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForTermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDistrictForTermNotFoundCode is the HTTP code returned for type GetDistrictForTermNotFound
const GetDistrictForTermNotFoundCode int = 404

/*
GetDistrictForTermNotFound Entity Not Found

swagger:response getDistrictForTermNotFound
*/
type GetDistrictForTermNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetDistrictForTermNotFound creates GetDistrictForTermNotFound with default headers values
func NewGetDistrictForTermNotFound() *GetDistrictForTermNotFound {

	return &GetDistrictForTermNotFound{}
}

// WithPayload adds the payload to the get district for term not found response
func (o *GetDistrictForTermNotFound) WithPayload(payload *models.NotFound) *GetDistrictForTermNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for term not found response
func (o *GetDistrictForTermNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForTermNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
