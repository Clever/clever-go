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

// GetUsersForSchoolReader is a Reader for the GetUsersForSchool structure.
type GetUsersForSchoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersForSchoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersForSchoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUsersForSchoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersForSchoolOK creates a GetUsersForSchoolOK with default headers values
func NewGetUsersForSchoolOK() *GetUsersForSchoolOK {
	return &GetUsersForSchoolOK{}
}

/* GetUsersForSchoolOK describes a response with status code 200, with default header values.

OK Response
*/
type GetUsersForSchoolOK struct {
	Payload *models.UsersResponse
}

func (o *GetUsersForSchoolOK) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/users][%d] getUsersForSchoolOK  %+v", 200, o.Payload)
}
func (o *GetUsersForSchoolOK) GetPayload() *models.UsersResponse {
	return o.Payload
}

func (o *GetUsersForSchoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForSchoolNotFound creates a GetUsersForSchoolNotFound with default headers values
func NewGetUsersForSchoolNotFound() *GetUsersForSchoolNotFound {
	return &GetUsersForSchoolNotFound{}
}

/* GetUsersForSchoolNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetUsersForSchoolNotFound struct {
	Payload *models.NotFound
}

func (o *GetUsersForSchoolNotFound) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/users][%d] getUsersForSchoolNotFound  %+v", 404, o.Payload)
}
func (o *GetUsersForSchoolNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetUsersForSchoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
