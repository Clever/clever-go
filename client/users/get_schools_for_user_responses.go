// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetSchoolsForUserReader is a Reader for the GetSchoolsForUser structure.
type GetSchoolsForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchoolsForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchoolsForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchoolsForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{id}/schools] getSchoolsForUser", response, response.Code())
	}
}

// NewGetSchoolsForUserOK creates a GetSchoolsForUserOK with default headers values
func NewGetSchoolsForUserOK() *GetSchoolsForUserOK {
	return &GetSchoolsForUserOK{}
}

/*
GetSchoolsForUserOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSchoolsForUserOK struct {
	Payload *models.SchoolsResponse
}

// IsSuccess returns true when this get schools for user o k response has a 2xx status code
func (o *GetSchoolsForUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schools for user o k response has a 3xx status code
func (o *GetSchoolsForUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schools for user o k response has a 4xx status code
func (o *GetSchoolsForUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schools for user o k response has a 5xx status code
func (o *GetSchoolsForUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schools for user o k response a status code equal to that given
func (o *GetSchoolsForUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schools for user o k response
func (o *GetSchoolsForUserOK) Code() int {
	return 200
}

func (o *GetSchoolsForUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/schools][%d] getSchoolsForUserOK  %+v", 200, o.Payload)
}

func (o *GetSchoolsForUserOK) String() string {
	return fmt.Sprintf("[GET /users/{id}/schools][%d] getSchoolsForUserOK  %+v", 200, o.Payload)
}

func (o *GetSchoolsForUserOK) GetPayload() *models.SchoolsResponse {
	return o.Payload
}

func (o *GetSchoolsForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchoolsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchoolsForUserNotFound creates a GetSchoolsForUserNotFound with default headers values
func NewGetSchoolsForUserNotFound() *GetSchoolsForUserNotFound {
	return &GetSchoolsForUserNotFound{}
}

/*
GetSchoolsForUserNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetSchoolsForUserNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get schools for user not found response has a 2xx status code
func (o *GetSchoolsForUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get schools for user not found response has a 3xx status code
func (o *GetSchoolsForUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schools for user not found response has a 4xx status code
func (o *GetSchoolsForUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get schools for user not found response has a 5xx status code
func (o *GetSchoolsForUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get schools for user not found response a status code equal to that given
func (o *GetSchoolsForUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get schools for user not found response
func (o *GetSchoolsForUserNotFound) Code() int {
	return 404
}

func (o *GetSchoolsForUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}/schools][%d] getSchoolsForUserNotFound  %+v", 404, o.Payload)
}

func (o *GetSchoolsForUserNotFound) String() string {
	return fmt.Sprintf("[GET /users/{id}/schools][%d] getSchoolsForUserNotFound  %+v", 404, o.Payload)
}

func (o *GetSchoolsForUserNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSchoolsForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}