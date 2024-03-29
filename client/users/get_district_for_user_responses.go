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

// GetDistrictForUserReader is a Reader for the GetDistrictForUser structure.
type GetDistrictForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistrictForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistrictForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDistrictForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistrictForUserOK creates a GetDistrictForUserOK with default headers values
func NewGetDistrictForUserOK() *GetDistrictForUserOK {
	return &GetDistrictForUserOK{}
}

/* GetDistrictForUserOK describes a response with status code 200, with default header values.

OK Response
*/
type GetDistrictForUserOK struct {
	Payload *models.DistrictResponse
}

func (o *GetDistrictForUserOK) Error() string {
	return fmt.Sprintf("[GET /users/{id}/district][%d] getDistrictForUserOK  %+v", 200, o.Payload)
}
func (o *GetDistrictForUserOK) GetPayload() *models.DistrictResponse {
	return o.Payload
}

func (o *GetDistrictForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrictResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistrictForUserNotFound creates a GetDistrictForUserNotFound with default headers values
func NewGetDistrictForUserNotFound() *GetDistrictForUserNotFound {
	return &GetDistrictForUserNotFound{}
}

/* GetDistrictForUserNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetDistrictForUserNotFound struct {
	Payload *models.NotFound
}

func (o *GetDistrictForUserNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{id}/district][%d] getDistrictForUserNotFound  %+v", 404, o.Payload)
}
func (o *GetDistrictForUserNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetDistrictForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
