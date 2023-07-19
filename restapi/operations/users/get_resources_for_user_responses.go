// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetResourcesForUserOKCode is the HTTP code returned for type GetResourcesForUserOK
const GetResourcesForUserOKCode int = 200

/*
GetResourcesForUserOK OK Response

swagger:response getResourcesForUserOK
*/
type GetResourcesForUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResourcesResponse `json:"body,omitempty"`
}

// NewGetResourcesForUserOK creates GetResourcesForUserOK with default headers values
func NewGetResourcesForUserOK() *GetResourcesForUserOK {

	return &GetResourcesForUserOK{}
}

// WithPayload adds the payload to the get resources for user o k response
func (o *GetResourcesForUserOK) WithPayload(payload *models.ResourcesResponse) *GetResourcesForUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources for user o k response
func (o *GetResourcesForUserOK) SetPayload(payload *models.ResourcesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesForUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetResourcesForUserNotFoundCode is the HTTP code returned for type GetResourcesForUserNotFound
const GetResourcesForUserNotFoundCode int = 404

/*
GetResourcesForUserNotFound Entity Not Found

swagger:response getResourcesForUserNotFound
*/
type GetResourcesForUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetResourcesForUserNotFound creates GetResourcesForUserNotFound with default headers values
func NewGetResourcesForUserNotFound() *GetResourcesForUserNotFound {

	return &GetResourcesForUserNotFound{}
}

// WithPayload adds the payload to the get resources for user not found response
func (o *GetResourcesForUserNotFound) WithPayload(payload *models.NotFound) *GetResourcesForUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources for user not found response
func (o *GetResourcesForUserNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesForUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
