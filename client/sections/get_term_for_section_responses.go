// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetTermForSectionReader is a Reader for the GetTermForSection structure.
type GetTermForSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTermForSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTermForSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTermForSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sections/{id}/term] getTermForSection", response, response.Code())
	}
}

// NewGetTermForSectionOK creates a GetTermForSectionOK with default headers values
func NewGetTermForSectionOK() *GetTermForSectionOK {
	return &GetTermForSectionOK{}
}

/*
GetTermForSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetTermForSectionOK struct {
	Payload *models.TermResponse
}

// IsSuccess returns true when this get term for section o k response has a 2xx status code
func (o *GetTermForSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get term for section o k response has a 3xx status code
func (o *GetTermForSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get term for section o k response has a 4xx status code
func (o *GetTermForSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get term for section o k response has a 5xx status code
func (o *GetTermForSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get term for section o k response a status code equal to that given
func (o *GetTermForSectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get term for section o k response
func (o *GetTermForSectionOK) Code() int {
	return 200
}

func (o *GetTermForSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/term][%d] getTermForSectionOK  %+v", 200, o.Payload)
}

func (o *GetTermForSectionOK) String() string {
	return fmt.Sprintf("[GET /sections/{id}/term][%d] getTermForSectionOK  %+v", 200, o.Payload)
}

func (o *GetTermForSectionOK) GetPayload() *models.TermResponse {
	return o.Payload
}

func (o *GetTermForSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TermResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTermForSectionNotFound creates a GetTermForSectionNotFound with default headers values
func NewGetTermForSectionNotFound() *GetTermForSectionNotFound {
	return &GetTermForSectionNotFound{}
}

/*
GetTermForSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetTermForSectionNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get term for section not found response has a 2xx status code
func (o *GetTermForSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get term for section not found response has a 3xx status code
func (o *GetTermForSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get term for section not found response has a 4xx status code
func (o *GetTermForSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get term for section not found response has a 5xx status code
func (o *GetTermForSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get term for section not found response a status code equal to that given
func (o *GetTermForSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get term for section not found response
func (o *GetTermForSectionNotFound) Code() int {
	return 404
}

func (o *GetTermForSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/term][%d] getTermForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetTermForSectionNotFound) String() string {
	return fmt.Sprintf("[GET /sections/{id}/term][%d] getTermForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetTermForSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetTermForSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
