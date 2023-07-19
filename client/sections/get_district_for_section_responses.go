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

// GetDistrictForSectionReader is a Reader for the GetDistrictForSection structure.
type GetDistrictForSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistrictForSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistrictForSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDistrictForSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sections/{id}/district] getDistrictForSection", response, response.Code())
	}
}

// NewGetDistrictForSectionOK creates a GetDistrictForSectionOK with default headers values
func NewGetDistrictForSectionOK() *GetDistrictForSectionOK {
	return &GetDistrictForSectionOK{}
}

/*
GetDistrictForSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetDistrictForSectionOK struct {
	Payload *models.DistrictResponse
}

// IsSuccess returns true when this get district for section o k response has a 2xx status code
func (o *GetDistrictForSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get district for section o k response has a 3xx status code
func (o *GetDistrictForSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for section o k response has a 4xx status code
func (o *GetDistrictForSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get district for section o k response has a 5xx status code
func (o *GetDistrictForSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for section o k response a status code equal to that given
func (o *GetDistrictForSectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get district for section o k response
func (o *GetDistrictForSectionOK) Code() int {
	return 200
}

func (o *GetDistrictForSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForSectionOK) String() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForSectionOK) GetPayload() *models.DistrictResponse {
	return o.Payload
}

func (o *GetDistrictForSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrictResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistrictForSectionNotFound creates a GetDistrictForSectionNotFound with default headers values
func NewGetDistrictForSectionNotFound() *GetDistrictForSectionNotFound {
	return &GetDistrictForSectionNotFound{}
}

/*
GetDistrictForSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetDistrictForSectionNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get district for section not found response has a 2xx status code
func (o *GetDistrictForSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get district for section not found response has a 3xx status code
func (o *GetDistrictForSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for section not found response has a 4xx status code
func (o *GetDistrictForSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get district for section not found response has a 5xx status code
func (o *GetDistrictForSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for section not found response a status code equal to that given
func (o *GetDistrictForSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get district for section not found response
func (o *GetDistrictForSectionNotFound) Code() int {
	return 404
}

func (o *GetDistrictForSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForSectionNotFound) String() string {
	return fmt.Sprintf("[GET /sections/{id}/district][%d] getDistrictForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetDistrictForSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
