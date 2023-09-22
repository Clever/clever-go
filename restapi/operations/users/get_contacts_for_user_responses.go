// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetContactsForUserOKCode is the HTTP code returned for type GetContactsForUserOK
const GetContactsForUserOKCode int = 200

/*
GetContactsForUserOK OK Response

swagger:response getContactsForUserOK
*/
type GetContactsForUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.UsersResponse `json:"body,omitempty"`
}

// NewGetContactsForUserOK creates GetContactsForUserOK with default headers values
func NewGetContactsForUserOK() *GetContactsForUserOK {

	return &GetContactsForUserOK{}
}

// WithPayload adds the payload to the get contacts for user o k response
func (o *GetContactsForUserOK) WithPayload(payload *models.UsersResponse) *GetContactsForUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get contacts for user o k response
func (o *GetContactsForUserOK) SetPayload(payload *models.UsersResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContactsForUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetContactsForUserNotFoundCode is the HTTP code returned for type GetContactsForUserNotFound
const GetContactsForUserNotFoundCode int = 404

/*
GetContactsForUserNotFound Entity Not Found

swagger:response getContactsForUserNotFound
*/
type GetContactsForUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetContactsForUserNotFound creates GetContactsForUserNotFound with default headers values
func NewGetContactsForUserNotFound() *GetContactsForUserNotFound {

	return &GetContactsForUserNotFound{}
}

// WithPayload adds the payload to the get contacts for user not found response
func (o *GetContactsForUserNotFound) WithPayload(payload *models.NotFound) *GetContactsForUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get contacts for user not found response
func (o *GetContactsForUserNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetContactsForUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
