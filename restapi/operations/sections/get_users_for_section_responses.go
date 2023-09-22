// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetUsersForSectionOKCode is the HTTP code returned for type GetUsersForSectionOK
const GetUsersForSectionOKCode int = 200

/*
GetUsersForSectionOK OK Response

swagger:response getUsersForSectionOK
*/
type GetUsersForSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.UsersResponse `json:"body,omitempty"`
}

// NewGetUsersForSectionOK creates GetUsersForSectionOK with default headers values
func NewGetUsersForSectionOK() *GetUsersForSectionOK {

	return &GetUsersForSectionOK{}
}

// WithPayload adds the payload to the get users for section o k response
func (o *GetUsersForSectionOK) WithPayload(payload *models.UsersResponse) *GetUsersForSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users for section o k response
func (o *GetUsersForSectionOK) SetPayload(payload *models.UsersResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersForSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersForSectionNotFoundCode is the HTTP code returned for type GetUsersForSectionNotFound
const GetUsersForSectionNotFoundCode int = 404

/*
GetUsersForSectionNotFound Entity Not Found

swagger:response getUsersForSectionNotFound
*/
type GetUsersForSectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetUsersForSectionNotFound creates GetUsersForSectionNotFound with default headers values
func NewGetUsersForSectionNotFound() *GetUsersForSectionNotFound {

	return &GetUsersForSectionNotFound{}
}

// WithPayload adds the payload to the get users for section not found response
func (o *GetUsersForSectionNotFound) WithPayload(payload *models.NotFound) *GetUsersForSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users for section not found response
func (o *GetUsersForSectionNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersForSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
