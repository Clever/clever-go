// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetTermsReader is a Reader for the GetTerms structure.
type GetTermsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTermsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTermsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /terms] getTerms", response, response.Code())
	}
}

// NewGetTermsOK creates a GetTermsOK with default headers values
func NewGetTermsOK() *GetTermsOK {
	return &GetTermsOK{}
}

/*
GetTermsOK describes a response with status code 200, with default header values.

OK Response
*/
type GetTermsOK struct {
	Payload *models.TermsResponse
}

// IsSuccess returns true when this get terms o k response has a 2xx status code
func (o *GetTermsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get terms o k response has a 3xx status code
func (o *GetTermsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get terms o k response has a 4xx status code
func (o *GetTermsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get terms o k response has a 5xx status code
func (o *GetTermsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get terms o k response a status code equal to that given
func (o *GetTermsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get terms o k response
func (o *GetTermsOK) Code() int {
	return 200
}

func (o *GetTermsOK) Error() string {
	return fmt.Sprintf("[GET /terms][%d] getTermsOK  %+v", 200, o.Payload)
}

func (o *GetTermsOK) String() string {
	return fmt.Sprintf("[GET /terms][%d] getTermsOK  %+v", 200, o.Payload)
}

func (o *GetTermsOK) GetPayload() *models.TermsResponse {
	return o.Payload
}

func (o *GetTermsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}