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

// GetCoursesReader is a Reader for the GetCourses structure.
type GetCoursesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoursesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoursesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoursesOK creates a GetCoursesOK with default headers values
func NewGetCoursesOK() *GetCoursesOK {
	return &GetCoursesOK{}
}

/*
	GetCoursesOK describes a response with status code 200, with default header values.

OK Response
*/
type GetCoursesOK struct {
	Payload *models.CoursesResponse
}

func (o *GetCoursesOK) Error() string {
	return fmt.Sprintf("[GET /courses][%d] getCoursesOK  %+v", 200, o.Payload)
}
func (o *GetCoursesOK) GetPayload() *models.CoursesResponse {
	return o.Payload
}

func (o *GetCoursesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoursesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
