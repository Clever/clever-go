// Code generated by go-swagger; DO NOT EDIT.

package schools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetSchoolOKCode is the HTTP code returned for type GetSchoolOK
const GetSchoolOKCode int = 200

/*GetSchoolOK OK Response

swagger:response getSchoolOK
*/
type GetSchoolOK struct {

	/*
	  In: Body
	*/
	Payload *models.SchoolResponse `json:"body,omitempty"`
}

// NewGetSchoolOK creates GetSchoolOK with default headers values
func NewGetSchoolOK() *GetSchoolOK {

	return &GetSchoolOK{}
}

// WithPayload adds the payload to the get school o k response
func (o *GetSchoolOK) WithPayload(payload *models.SchoolResponse) *GetSchoolOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get school o k response
func (o *GetSchoolOK) SetPayload(payload *models.SchoolResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSchoolOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSchoolNotFoundCode is the HTTP code returned for type GetSchoolNotFound
const GetSchoolNotFoundCode int = 404

/*GetSchoolNotFound Entity Not Found

swagger:response getSchoolNotFound
*/
type GetSchoolNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetSchoolNotFound creates GetSchoolNotFound with default headers values
func NewGetSchoolNotFound() *GetSchoolNotFound {

	return &GetSchoolNotFound{}
}

// WithPayload adds the payload to the get school not found response
func (o *GetSchoolNotFound) WithPayload(payload *models.NotFound) *GetSchoolNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get school not found response
func (o *GetSchoolNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSchoolNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
