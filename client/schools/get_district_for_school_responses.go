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

// GetDistrictForSchoolReader is a Reader for the GetDistrictForSchool structure.
type GetDistrictForSchoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistrictForSchoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistrictForSchoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDistrictForSchoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /schools/{id}/district] getDistrictForSchool", response, response.Code())
	}
}

// NewGetDistrictForSchoolOK creates a GetDistrictForSchoolOK with default headers values
func NewGetDistrictForSchoolOK() *GetDistrictForSchoolOK {
	return &GetDistrictForSchoolOK{}
}

/*
GetDistrictForSchoolOK describes a response with status code 200, with default header values.

OK Response
*/
type GetDistrictForSchoolOK struct {
	Payload *models.DistrictResponse
}

// IsSuccess returns true when this get district for school o k response has a 2xx status code
func (o *GetDistrictForSchoolOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get district for school o k response has a 3xx status code
func (o *GetDistrictForSchoolOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for school o k response has a 4xx status code
func (o *GetDistrictForSchoolOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get district for school o k response has a 5xx status code
func (o *GetDistrictForSchoolOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for school o k response a status code equal to that given
func (o *GetDistrictForSchoolOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get district for school o k response
func (o *GetDistrictForSchoolOK) Code() int {
	return 200
}

func (o *GetDistrictForSchoolOK) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/district][%d] getDistrictForSchoolOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForSchoolOK) String() string {
	return fmt.Sprintf("[GET /schools/{id}/district][%d] getDistrictForSchoolOK  %+v", 200, o.Payload)
}

func (o *GetDistrictForSchoolOK) GetPayload() *models.DistrictResponse {
	return o.Payload
}

func (o *GetDistrictForSchoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DistrictResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistrictForSchoolNotFound creates a GetDistrictForSchoolNotFound with default headers values
func NewGetDistrictForSchoolNotFound() *GetDistrictForSchoolNotFound {
	return &GetDistrictForSchoolNotFound{}
}

/*
GetDistrictForSchoolNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetDistrictForSchoolNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get district for school not found response has a 2xx status code
func (o *GetDistrictForSchoolNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get district for school not found response has a 3xx status code
func (o *GetDistrictForSchoolNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get district for school not found response has a 4xx status code
func (o *GetDistrictForSchoolNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get district for school not found response has a 5xx status code
func (o *GetDistrictForSchoolNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get district for school not found response a status code equal to that given
func (o *GetDistrictForSchoolNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get district for school not found response
func (o *GetDistrictForSchoolNotFound) Code() int {
	return 404
}

func (o *GetDistrictForSchoolNotFound) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/district][%d] getDistrictForSchoolNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForSchoolNotFound) String() string {
	return fmt.Sprintf("[GET /schools/{id}/district][%d] getDistrictForSchoolNotFound  %+v", 404, o.Payload)
}

func (o *GetDistrictForSchoolNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetDistrictForSchoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
