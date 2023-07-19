// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetDistrictForCourseReader is a Reader for the GetDistrictForCourse structure.
type GetDistrictForCourseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistrictForCourseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistrictForCourseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDistrictForCourseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /courses/{id}/district] getDistrictForCourse", response, response.Code())
	}
}

// NewGetDistrictForCourseOK creates a GetDistrictForCourseOK with default headers values
func NewGetDistrictForCourseOK() *GetDistrictForCourseOK {
	return &GetDistrictForCourseOK{}
}

/*
GetDistrictForCourseOK describes a response with status code 200, with default header values.

OK Response
*/
type GetDistrictForCourseOK struct {
	Payload *models.DistrictResponse
}

// IsSuccess returns true when this get district for course o k response has a 2xx status code
func (o *GetDistrictForCourseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get district for course o k response has a 3xx status code
func (o *GetDistrictForCourseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for course o k response has a 4xx status code
func (o *GetDistrictForCourseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get district for course o k response has a 5xx status code
func (o *GetDistrictForCourseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for course o k response a status code equal to that given
func (o *GetDistrictForCourseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get district for course o k response
func (o *GetDistrictForCourseOK) Code() int {
	return 200
}

func (o *GetDistrictForCourseOK) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/district][%d] getDistrictForCourseOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForCourseOK) String() string {
	return fmt.Sprintf("[GET /courses/{id}/district][%d] getDistrictForCourseOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForCourseOK) GetPayload() *models.DistrictResponse {
	return o.Payload
}

func (o *GetDistrictForCourseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrictResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistrictForCourseNotFound creates a GetDistrictForCourseNotFound with default headers values
func NewGetDistrictForCourseNotFound() *GetDistrictForCourseNotFound {
	return &GetDistrictForCourseNotFound{}
}

/*
GetDistrictForCourseNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetDistrictForCourseNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get district for course not found response has a 2xx status code
func (o *GetDistrictForCourseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get district for course not found response has a 3xx status code
func (o *GetDistrictForCourseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for course not found response has a 4xx status code
func (o *GetDistrictForCourseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get district for course not found response has a 5xx status code
func (o *GetDistrictForCourseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for course not found response a status code equal to that given
func (o *GetDistrictForCourseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get district for course not found response
func (o *GetDistrictForCourseNotFound) Code() int {
	return 404
}

func (o *GetDistrictForCourseNotFound) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/district][%d] getDistrictForCourseNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForCourseNotFound) String() string {
	return fmt.Sprintf("[GET /courses/{id}/district][%d] getDistrictForCourseNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForCourseNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetDistrictForCourseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
