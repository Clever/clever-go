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

// GetCoursesForSchoolReader is a Reader for the GetCoursesForSchool structure.
type GetCoursesForSchoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoursesForSchoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoursesForSchoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCoursesForSchoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoursesForSchoolOK creates a GetCoursesForSchoolOK with default headers values
func NewGetCoursesForSchoolOK() *GetCoursesForSchoolOK {
	return &GetCoursesForSchoolOK{}
}

/* GetCoursesForSchoolOK describes a response with status code 200, with default header values.

OK Response
*/
type GetCoursesForSchoolOK struct {
	Payload *models.CoursesResponse
}

func (o *GetCoursesForSchoolOK) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/courses][%d] getCoursesForSchoolOK  %+v", 200, o.Payload)
}
func (o *GetCoursesForSchoolOK) GetPayload() *models.CoursesResponse {
	return o.Payload
}

func (o *GetCoursesForSchoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoursesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoursesForSchoolNotFound creates a GetCoursesForSchoolNotFound with default headers values
func NewGetCoursesForSchoolNotFound() *GetCoursesForSchoolNotFound {
	return &GetCoursesForSchoolNotFound{}
}

/* GetCoursesForSchoolNotFound describes a response with status code 404, with default header values.

Entity Not Found
*/
type GetCoursesForSchoolNotFound struct {
	Payload *models.NotFound
}

func (o *GetCoursesForSchoolNotFound) Error() string {
	return fmt.Sprintf("[GET /schools/{id}/courses][%d] getCoursesForSchoolNotFound  %+v", 404, o.Payload)
}
func (o *GetCoursesForSchoolNotFound) GetPayload() *models.NotFound {
	return o.Payload
}

func (o *GetCoursesForSchoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotFound)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
