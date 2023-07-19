// Code generated by go-swagger; DO NOT EDIT.

package schools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetDistrictForSchoolOKCode is the HTTP code returned for type GetDistrictForSchoolOK
const GetDistrictForSchoolOKCode int = 200

/*
GetDistrictForSchoolOK OK Response

swagger:response getDistrictForSchoolOK
*/
type GetDistrictForSchoolOK struct {

	/*
	  In: Body
	*/
	Payload *models.DistrictResponse `json:"body,omitempty"`
}

// NewGetDistrictForSchoolOK creates GetDistrictForSchoolOK with default headers values
func NewGetDistrictForSchoolOK() *GetDistrictForSchoolOK {

	return &GetDistrictForSchoolOK{}
}

// WithPayload adds the payload to the get district for school o k response
func (o *GetDistrictForSchoolOK) WithPayload(payload *models.DistrictResponse) *GetDistrictForSchoolOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for school o k response
func (o *GetDistrictForSchoolOK) SetPayload(payload *models.DistrictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForSchoolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDistrictForSchoolNotFoundCode is the HTTP code returned for type GetDistrictForSchoolNotFound
const GetDistrictForSchoolNotFoundCode int = 404

/*
GetDistrictForSchoolNotFound Entity Not Found

swagger:response getDistrictForSchoolNotFound
*/
type GetDistrictForSchoolNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetDistrictForSchoolNotFound creates GetDistrictForSchoolNotFound with default headers values
func NewGetDistrictForSchoolNotFound() *GetDistrictForSchoolNotFound {

	return &GetDistrictForSchoolNotFound{}
}

// WithPayload adds the payload to the get district for school not found response
func (o *GetDistrictForSchoolNotFound) WithPayload(payload *models.NotFound) *GetDistrictForSchoolNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for school not found response
func (o *GetDistrictForSchoolNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForSchoolNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}