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

// GetSchoolsReader is a Reader for the GetSchools structure.
type GetSchoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSchoolsOK creates a GetSchoolsOK with default headers values
func NewGetSchoolsOK() *GetSchoolsOK {
	return &GetSchoolsOK{}
}

/*
	GetSchoolsOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSchoolsOK struct {
	Payload *models.SchoolsResponse
}

func (o *GetSchoolsOK) Error() string {
	return fmt.Sprintf("[GET /schools][%d] getSchoolsOK  %+v", 200, o.Payload)
}
func (o *GetSchoolsOK) GetPayload() *models.SchoolsResponse {
	return o.Payload
}

func (o *GetSchoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchoolsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
