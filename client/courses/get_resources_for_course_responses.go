// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Clever/clever-go/models"
)

// GetResourcesForCourseReader is a Reader for the GetResourcesForCourse structure.
type GetResourcesForCourseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcesForCourseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcesForCourseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetResourcesForCourseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourcesForCourseOK creates a GetResourcesForCourseOK with default headers values
func NewGetResourcesForCourseOK() *GetResourcesForCourseOK {
	return &GetResourcesForCourseOK{}
}

/* GetResourcesForCourseOK describes a response with status code 200, with default header values.

OK Response
*/
type GetResourcesForCourseOK struct {
	Payload *models.ResourcesResponse
}

func (o *GetResourcesForCourseOK) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/resources][%d] getResourcesForCourseOK  %+v", 200, o.Payload)
}
func (o *GetResourcesForCourseOK) GetPayload() *models.ResourcesResponse {
	return o.Payload
}

func (o *GetResourcesForCourseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcesForCourseNotFound creates a GetResourcesForCourseNotFound with default headers values
func NewGetResourcesForCourseNotFound() *GetResourcesForCourseNotFound {
	return &GetResourcesForCourseNotFound{}
}

/* GetResourcesForCourseNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetResourcesForCourseNotFound struct {
	Payload *models.NotFound
}

func (o *GetResourcesForCourseNotFound) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/resources][%d] getResourcesForCourseNotFound  %+v", 404, o.Payload)
}
func (o *GetResourcesForCourseNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetResourcesForCourseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
