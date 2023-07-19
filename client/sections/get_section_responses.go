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

// GetSectionReader is a Reader for the GetSection structure.
type GetSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sections/{id}] getSection", response, response.Code())
	}
}

// NewGetSectionOK creates a GetSectionOK with default headers values
func NewGetSectionOK() *GetSectionOK {
	return &GetSectionOK{}
}

/*
GetSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSectionOK struct {
	Payload *models.SectionResponse
}

// IsSuccess returns true when this get section o k response has a 2xx status code
func (o *GetSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get section o k response has a 3xx status code
func (o *GetSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get section o k response has a 4xx status code
func (o *GetSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get section o k response has a 5xx status code
func (o *GetSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get section o k response a status code equal to that given
func (o *GetSectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get section o k response
func (o *GetSectionOK) Code() int {
	return 200
}

func (o *GetSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}][%d] getSectionOK  %+v", 200, o.Payload)
}

func (o *GetSectionOK) String() string {
	return fmt.Sprintf("[GET /sections/{id}][%d] getSectionOK  %+v", 200, o.Payload)
}

func (o *GetSectionOK) GetPayload() *models.SectionResponse {
	return o.Payload
}

func (o *GetSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSectionNotFound creates a GetSectionNotFound with default headers values
func NewGetSectionNotFound() *GetSectionNotFound {
	return &GetSectionNotFound{}
}

/*
GetSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetSectionNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get section not found response has a 2xx status code
func (o *GetSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get section not found response has a 3xx status code
func (o *GetSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get section not found response has a 4xx status code
func (o *GetSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get section not found response has a 5xx status code
func (o *GetSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get section not found response a status code equal to that given
func (o *GetSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get section not found response
func (o *GetSectionNotFound) Code() int {
	return 404
}

func (o *GetSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}][%d] getSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetSectionNotFound) String() string {
	return fmt.Sprintf("[GET /sections/{id}][%d] getSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}