// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetCourseForSectionOKCode is the HTTP code returned for type GetCourseForSectionOK
const GetCourseForSectionOKCode int = 200

/*GetCourseForSectionOK OK Response

swagger:response getCourseForSectionOK
*/
type GetCourseForSectionOK struct {

	/*
	  In: Body
	*/
	Payload *models.CourseResponse `json:"body,omitempty"`
}

// NewGetCourseForSectionOK creates GetCourseForSectionOK with default headers values
func NewGetCourseForSectionOK() *GetCourseForSectionOK {

	return &GetCourseForSectionOK{}
}

// WithPayload adds the payload to the get course for section o k response
func (o *GetCourseForSectionOK) WithPayload(payload *models.CourseResponse) *GetCourseForSectionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get course for section o k response
func (o *GetCourseForSectionOK) SetPayload(payload *models.CourseResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCourseForSectionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCourseForSectionNotFoundCode is the HTTP code returned for type GetCourseForSectionNotFound
const GetCourseForSectionNotFoundCode int = 404

/*GetCourseForSectionNotFound Entity Not Found

swagger:response getCourseForSectionNotFound
*/
type GetCourseForSectionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetCourseForSectionNotFound creates GetCourseForSectionNotFound with default headers values
func NewGetCourseForSectionNotFound() *GetCourseForSectionNotFound {

	return &GetCourseForSectionNotFound{}
}

// WithPayload adds the payload to the get course for section not found response
func (o *GetCourseForSectionNotFound) WithPayload(payload *models.NotFound) *GetCourseForSectionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get course for section not found response
func (o *GetCourseForSectionNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCourseForSectionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
