// Code generated by go-swagger; DO NOT EDIT.

package schools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetTermsForSchoolReader is a Reader for the GetTermsForSchool structure.
type GetTermsForSchoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTermsForSchoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTermsForSchoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTermsForSchoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTermsForSchoolOK creates a GetTermsForSchoolOK with default headers values
func NewGetTermsForSchoolOK() *GetTermsForSchoolOK {
	return &GetTermsForSchoolOK{}
}

/* GetTermsForSchoolOK describes a response with status code 200, with default header values.

OK Response
*/
type GetTermsForSchoolOK struct {
	Payload *models.TermsResponse
}

func (o *GetTermsForSchoolOK) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/terms][%d] getTermsForSchoolOK  %+v", 200, o.Payload)
}
func (o *GetTermsForSchoolOK) GetPayload() *models.TermsResponse {
	return o.Payload
}

func (o *GetTermsForSchoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTermsForSchoolNotFound creates a GetTermsForSchoolNotFound with default headers values
func NewGetTermsForSchoolNotFound() *GetTermsForSchoolNotFound {
	return &GetTermsForSchoolNotFound{}
}

/* GetTermsForSchoolNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetTermsForSchoolNotFound struct {
	Payload *models.NotFound
}

func (o *GetTermsForSchoolNotFound) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/terms][%d] getTermsForSchoolNotFound  %+v", 404, o.Payload)
}
func (o *GetTermsForSchoolNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetTermsForSchoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
