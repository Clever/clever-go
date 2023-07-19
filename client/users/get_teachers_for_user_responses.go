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

// GetTeachersForUserReader is a Reader for the GetTeachersForUser structure.
type GetTeachersForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeachersForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeachersForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTeachersForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTeachersForUserOK creates a GetTeachersForUserOK with default headers values
func NewGetTeachersForUserOK() *GetTeachersForUserOK {
	return &GetTeachersForUserOK{}
}

/* GetTeachersForUserOK describes a response with status code 200, with default header values.

OK Response
*/
type GetTeachersForUserOK struct {
	Payload *models.UsersResponse
}

func (o *GetTeachersForUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/myteachers][%d] getTeachersForUserOK  %+v", 200, o.Payload)
}
func (o *GetTeachersForUserOK) GetPayload() *models.UsersResponse {
	return o.Payload
}

func (o *GetTeachersForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeachersForUserNotFound creates a GetTeachersForUserNotFound with default headers values
func NewGetTeachersForUserNotFound() *GetTeachersForUserNotFound {
	return &GetTeachersForUserNotFound{}
}

/* GetTeachersForUserNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetTeachersForUserNotFound struct {
	Payload *models.NotFound
}

func (o *GetTeachersForUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}/myteachers][%d] getTeachersForUserNotFound  %+v", 404, o.Payload)
}
func (o *GetTeachersForUserNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetTeachersForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
