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

// GetUsersForSectionReader is a Reader for the GetUsersForSection structure.
type GetUsersForSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersForSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersForSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUsersForSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sections/{id}/users] getUsersForSection", response, response.Code())
	}
}

// NewGetUsersForSectionOK creates a GetUsersForSectionOK with default headers values
func NewGetUsersForSectionOK() *GetUsersForSectionOK {
	return &GetUsersForSectionOK{}
}

/*
GetUsersForSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetUsersForSectionOK struct {
	Payload *models.UsersResponse
}

// IsSuccess returns true when this get users for section o k response has a 2xx status code
func (o *GetUsersForSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users for section o k response has a 3xx status code
func (o *GetUsersForSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for section o k response has a 4xx status code
func (o *GetUsersForSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users for section o k response has a 5xx status code
func (o *GetUsersForSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for section o k response a status code equal to that given
func (o *GetUsersForSectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users for section o k response
func (o *GetUsersForSectionOK) Code() int {
	return 200
}

func (o *GetUsersForSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/users][%d] getUsersForSectionOK  %+v", 200, o.Payload)
}

func (o *GetUsersForSectionOK) String() string {
	return fmt.Sprintf("[GET /sections/{id}/users][%d] getUsersForSectionOK  %+v", 200, o.Payload)
}

func (o *GetUsersForSectionOK) GetPayload() *models.UsersResponse {
	return o.Payload
}

func (o *GetUsersForSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForSectionNotFound creates a GetUsersForSectionNotFound with default headers values
func NewGetUsersForSectionNotFound() *GetUsersForSectionNotFound {
	return &GetUsersForSectionNotFound{}
}

/*
GetUsersForSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetUsersForSectionNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get users for section not found response has a 2xx status code
func (o *GetUsersForSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users for section not found response has a 3xx status code
func (o *GetUsersForSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users for section not found response has a 4xx status code
func (o *GetUsersForSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users for section not found response has a 5xx status code
func (o *GetUsersForSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get users for section not found response a status code equal to that given
func (o *GetUsersForSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get users for section not found response
func (o *GetUsersForSectionNotFound) Code() int {
	return 404
}

func (o *GetUsersForSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/users][%d] getUsersForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersForSectionNotFound) String() string {
	return fmt.Sprintf("[GET /sections/{id}/users][%d] getUsersForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersForSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetUsersForSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
