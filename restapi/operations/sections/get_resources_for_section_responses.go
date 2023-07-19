// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetResourcesForSectionOKCode is the HTTP code returned for type GetResourcesForSectionOK
const GetResourcesForSectionOKCode int = 200

/*GetResourcesForSectionOK OK Response

swagger:response getResourcesForSectionOK
*/
type GetResourcesForSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResourcesResponse `json:"body,omitempty"`
}

// NewGetResourcesForSectionOK creates GetResourcesForSectionOK with default headers values
func NewGetResourcesForSectionOK() *GetResourcesForSectionOK {

	return &GetResourcesForSectionOK{}
}

// WithPayload adds the payload to the get resources for section o k response
func (o *GetResourcesForSectionOK) WithPayload(payload *models.ResourcesResponse) *GetResourcesForSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources for section o k response
func (o *GetResourcesForSectionOK) SetPayload(payload *models.ResourcesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesForSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetResourcesForSectionNotFoundCode is the HTTP code returned for type GetResourcesForSectionNotFound
const GetResourcesForSectionNotFoundCode int = 404

/*GetResourcesForSectionNotFound Entity Not Found

swagger:response getResourcesForSectionNotFound
*/
type GetResourcesForSectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetResourcesForSectionNotFound creates GetResourcesForSectionNotFound with default headers values
func NewGetResourcesForSectionNotFound() *GetResourcesForSectionNotFound {

	return &GetResourcesForSectionNotFound{}
}

// WithPayload adds the payload to the get resources for section not found response
func (o *GetResourcesForSectionNotFound) WithPayload(payload *models.NotFound) *GetResourcesForSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get resources for section not found response
func (o *GetResourcesForSectionNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetResourcesForSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
