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

// GetResourcesReader is a Reader for the GetResources structure.
type GetResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /resources] getResources", response, response.Code())
	}
}

// NewGetResourcesOK creates a GetResourcesOK with default headers values
func NewGetResourcesOK() *GetResourcesOK {
	return &GetResourcesOK{}
}

/*
GetResourcesOK describes a response with status code 200, with default header values.

OK Response
*/
type GetResourcesOK struct {
	Payload *models.ResourcesResponse
}

// IsSuccess returns true when this get resources o k response has a 2xx status code
func (o *GetResourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resources o k response has a 3xx status code
func (o *GetResourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resources o k response has a 4xx status code
func (o *GetResourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resources o k response has a 5xx status code
func (o *GetResourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resources o k response a status code equal to that given
func (o *GetResourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get resources o k response
func (o *GetResourcesOK) Code() int {
	return 200
}

func (o *GetResourcesOK) Error() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) String() string {
	return fmt.Sprintf("[GET /resources][%d] getResourcesOK  %+v", 200, o.Payload)
}

func (o *GetResourcesOK) GetPayload() *models.ResourcesResponse {
	return o.Payload
}

func (o *GetResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
