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

// GetResourceReader is a Reader for the GetResource structure.
type GetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceOK creates a GetResourceOK with default headers values
func NewGetResourceOK() *GetResourceOK {
	return &GetResourceOK{}
}

/* GetResourceOK describes a response with status code 200, with default header values.

OK Response
*/
type GetResourceOK struct {
	Payload *models.ResourceResponse
}

func (o *GetResourceOK) Error() string {
	return fmt.Sprintf("[GET /resources/{id}][%d] getResourceOK  %+v", 200, o.Payload)
}
func (o *GetResourceOK) GetPayload() *models.ResourceResponse {
	return o.Payload
}

func (o *GetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceNotFound creates a GetResourceNotFound with default headers values
func NewGetResourceNotFound() *GetResourceNotFound {
	return &GetResourceNotFound{}
}

/* GetResourceNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetResourceNotFound struct {
	Payload *models.NotFound
}

func (o *GetResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /resources/{id}][%d] getResourceNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
