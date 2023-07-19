// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetCoursesForResourceReader is a Reader for the GetCoursesForResource structure.
type GetCoursesForResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoursesForResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoursesForResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCoursesForResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /resources/{id}/courses] getCoursesForResource", response, response.Code())
	}
}

// NewGetCoursesForResourceOK creates a GetCoursesForResourceOK with default headers values
func NewGetCoursesForResourceOK() *GetCoursesForResourceOK {
	return &GetCoursesForResourceOK{}
}

/*
GetCoursesForResourceOK describes a response with status code 200, with default header values.

OK Response
*/
type GetCoursesForResourceOK struct {
	Payload *models.CoursesResponse
}

// IsSuccess returns true when this get courses for resource o k response has a 2xx status code
func (o *GetCoursesForResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get courses for resource o k response has a 3xx status code
func (o *GetCoursesForResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get courses for resource o k response has a 4xx status code
func (o *GetCoursesForResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get courses for resource o k response has a 5xx status code
func (o *GetCoursesForResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get courses for resource o k response a status code equal to that given
func (o *GetCoursesForResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get courses for resource o k response
func (o *GetCoursesForResourceOK) Code() int {
	return 200
}

func (o *GetCoursesForResourceOK) Error() string {
	return fmt.Sprintf("[GET /resources/{id}/courses][%d] getCoursesForResourceOK  %+v", 200, o.Payload)
}

func (o *GetCoursesForResourceOK) String() string {
	return fmt.Sprintf("[GET /resources/{id}/courses][%d] getCoursesForResourceOK  %+v", 200, o.Payload)
}

func (o *GetCoursesForResourceOK) GetPayload() *models.CoursesResponse {
	return o.Payload
}

func (o *GetCoursesForResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoursesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoursesForResourceNotFound creates a GetCoursesForResourceNotFound with default headers values
func NewGetCoursesForResourceNotFound() *GetCoursesForResourceNotFound {
	return &GetCoursesForResourceNotFound{}
}

/*
GetCoursesForResourceNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetCoursesForResourceNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get courses for resource not found response has a 2xx status code
func (o *GetCoursesForResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get courses for resource not found response has a 3xx status code
func (o *GetCoursesForResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get courses for resource not found response has a 4xx status code
func (o *GetCoursesForResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get courses for resource not found response has a 5xx status code
func (o *GetCoursesForResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get courses for resource not found response a status code equal to that given
func (o *GetCoursesForResourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get courses for resource not found response
func (o *GetCoursesForResourceNotFound) Code() int {
	return 404
}

func (o *GetCoursesForResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /resources/{id}/courses][%d] getCoursesForResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetCoursesForResourceNotFound) String() string {
	return fmt.Sprintf("[GET /resources/{id}/courses][%d] getCoursesForResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetCoursesForResourceNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetCoursesForResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}