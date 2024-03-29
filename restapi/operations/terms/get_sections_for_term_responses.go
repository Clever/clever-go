// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Clever/clever-go/models"
)

// GetSectionsForTermOKCode is the HTTP code returned for type GetSectionsForTermOK
const GetSectionsForTermOKCode int = 200

/*GetSectionsForTermOK OK Response

swagger:response getSectionsForTermOK
*/
type GetSectionsForTermOK struct {

	/*
	  In: Body
	*/
	Payload *models.SectionsResponse `json:"body,omitempty"`
}

// NewGetSectionsForTermOK creates GetSectionsForTermOK with default headers values
func NewGetSectionsForTermOK() *GetSectionsForTermOK {

	return &GetSectionsForTermOK{}
}

// WithPayload adds the payload to the get sections for term o k response
func (o *GetSectionsForTermOK) WithPayload(payload *models.SectionsResponse) *GetSectionsForTermOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sections for term o k response
func (o *GetSectionsForTermOK) SetPayload(payload *models.SectionsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSectionsForTermOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetSectionsForTermNotFoundCode is the HTTP code returned for type GetSectionsForTermNotFound
const GetSectionsForTermNotFoundCode int = 404

/*GetSectionsForTermNotFound Entity Not Found

swagger:response getSectionsForTermNotFound
*/
type GetSectionsForTermNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.NotFound `json:"body,omitempty"`
}

// NewGetSectionsForTermNotFound creates GetSectionsForTermNotFound with default headers values
func NewGetSectionsForTermNotFound() *GetSectionsForTermNotFound {

	return &GetSectionsForTermNotFound{}
}

// WithPayload adds the payload to the get sections for term not found response
func (o *GetSectionsForTermNotFound) WithPayload(payload *models.NotFound) *GetSectionsForTermNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sections for term not found response
func (o *GetSectionsForTermNotFound) SetPayload(payload *models.NotFound) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSectionsForTermNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
