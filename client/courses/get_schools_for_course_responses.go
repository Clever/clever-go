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

// GetSchoolsForCourseReader is a Reader for the GetSchoolsForCourse structure.
type GetSchoolsForCourseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchoolsForCourseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchoolsForCourseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSchoolsForCourseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchoolsForCourseOK creates a GetSchoolsForCourseOK with default headers values
func NewGetSchoolsForCourseOK() *GetSchoolsForCourseOK {
	return &GetSchoolsForCourseOK{}
}

/* GetSchoolsForCourseOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSchoolsForCourseOK struct {
	Payload *models.SchoolsResponse
}

func (o *GetSchoolsForCourseOK) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/schools][%d] getSchoolsForCourseOK  %+v", 200, o.Payload)
}
func (o *GetSchoolsForCourseOK) GetPayload() *models.SchoolsResponse {
	return o.Payload
}

func (o *GetSchoolsForCourseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchoolsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchoolsForCourseNotFound creates a GetSchoolsForCourseNotFound with default headers values
func NewGetSchoolsForCourseNotFound() *GetSchoolsForCourseNotFound {
	return &GetSchoolsForCourseNotFound{}
}

/* GetSchoolsForCourseNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetSchoolsForCourseNotFound struct {
	Payload *models.NotFound
}

func (o *GetSchoolsForCourseNotFound) Error() string {
	return fmt.Sprintf("[GET /courses/{id}/schools][%d] getSchoolsForCourseNotFound  %+v", 404, o.Payload)
}
func (o *GetSchoolsForCourseNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetSchoolsForCourseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}