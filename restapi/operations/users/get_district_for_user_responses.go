// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetDistrictForUserOKCode is the HTTP code returned for type GetDistrictForUserOK
const GetDistrictForUserOKCode int = 200

/*GetDistrictForUserOK OK Response

swagger:response getDistrictForUserOK
*/
type GetDistrictForUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.DistrictResponse `json:"body,omitempty"`
}

// NewGetDistrictForUserOK creates GetDistrictForUserOK with default headers values
func NewGetDistrictForUserOK() *GetDistrictForUserOK {

	return &GetDistrictForUserOK{}
}

// WithPayload adds the payload to the get district for user o k response
func (o *GetDistrictForUserOK) WithPayload(payload *models.DistrictResponse) *GetDistrictForUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for user o k response
func (o *GetDistrictForUserOK) SetPayload(payload *models.DistrictResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDistrictForUserNotFoundCode is the HTTP code returned for type GetDistrictForUserNotFound
const GetDistrictForUserNotFoundCode int = 404

/*GetDistrictForUserNotFound Entity Not Found

swagger:response getDistrictForUserNotFound
*/
type GetDistrictForUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetDistrictForUserNotFound creates GetDistrictForUserNotFound with default headers values
func NewGetDistrictForUserNotFound() *GetDistrictForUserNotFound {

	return &GetDistrictForUserNotFound{}
}

// WithPayload adds the payload to the get district for user not found response
func (o *GetDistrictForUserNotFound) WithPayload(payload *models.NotFound) *GetDistrictForUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get district for user not found response
func (o *GetDistrictForUserNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDistrictForUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
