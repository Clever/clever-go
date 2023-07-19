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

// GetResourcesForSectionReader is a Reader for the GetResourcesForSection structure.
type GetResourcesForSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesForSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesForSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetResourcesForSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /sections/{id}/resources] getResourcesForSection", response, response.Code())
	}
}

// NewGetResourcesForSectionOK creates a GetResourcesForSectionOK with default headers values
func NewGetResourcesForSectionOK() *GetResourcesForSectionOK {
	return &GetResourcesForSectionOK{}
}

/*
GetResourcesForSectionOK describes a response with status code 200, with default header values.

OK Response
*/
type GetResourcesForSectionOK struct {
	Payload *models.ResourcesResponse
}

// IsSuccess returns true when this get resources for section o k response has a 2xx status code
func (o *GetResourcesForSectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources for section o k response has a 3xx status code
func (o *GetResourcesForSectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources for section o k response has a 4xx status code
func (o *GetResourcesForSectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources for section o k response has a 5xx status code
func (o *GetResourcesForSectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources for section o k response a status code equal to that given
func (o *GetResourcesForSectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get resources for section o k response
func (o *GetResourcesForSectionOK) Code() int {
	return 200
}

func (o *GetResourcesForSectionOK) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/resources][%d] getResourcesForSectionOK  %+v", 200, o.Payload)
}

func (o *GetResourcesForSectionOK) String() string {
	return fmt.Sprintf("[GET /sections/{id}/resources][%d] getResourcesForSectionOK  %+v", 200, o.Payload)
}

func (o *GetResourcesForSectionOK) GetPayload() *models.ResourcesResponse {
	return o.Payload
}

func (o *GetResourcesForSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesForSectionNotFound creates a GetResourcesForSectionNotFound with default headers values
func NewGetResourcesForSectionNotFound() *GetResourcesForSectionNotFound {
	return &GetResourcesForSectionNotFound{}
}

/*
GetResourcesForSectionNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetResourcesForSectionNotFound struct {
	Payload *models.NotFound
}

// IsSuccess returns true when this get resources for section not found response has a 2xx status code
func (o *GetResourcesForSectionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get resources for section not found response has a 3xx status code
func (o *GetResourcesForSectionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources for section not found response has a 4xx status code
func (o *GetResourcesForSectionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get resources for section not found response has a 5xx status code
func (o *GetResourcesForSectionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources for section not found response a status code equal to that given
func (o *GetResourcesForSectionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get resources for section not found response
func (o *GetResourcesForSectionNotFound) Code() int {
	return 404
}

func (o *GetResourcesForSectionNotFound) Error() string {
	return fmt.Sprintf("[GET /sections/{id}/resources][%d] getResourcesForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetResourcesForSectionNotFound) String() string {
	return fmt.Sprintf("[GET /sections/{id}/resources][%d] getResourcesForSectionNotFound  %+v", 404, o.Payload)
}

func (o *GetResourcesForSectionNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetResourcesForSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}