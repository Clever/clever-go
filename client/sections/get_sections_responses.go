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

// GetSectionsReader is a Reader for the GetSections structure.
type GetSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /sections] getSections", response, response.Code())
	}
}

// NewGetSectionsOK creates a GetSectionsOK with default headers values
func NewGetSectionsOK() *GetSectionsOK {
	return &GetSectionsOK{}
}

/*
GetSectionsOK describes a response with status code 200, with default header values.

OK Response
*/
type GetSectionsOK struct {
	Payload *models.SectionsResponse
}

// IsSuccess returns true when this get sections o k response has a 2xx status code
func (o *GetSectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sections o k response has a 3xx status code
func (o *GetSectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sections o k response has a 4xx status code
func (o *GetSectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sections o k response has a 5xx status code
func (o *GetSectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sections o k response a status code equal to that given
func (o *GetSectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sections o k response
func (o *GetSectionsOK) Code() int {
	return 200
}

func (o *GetSectionsOK) Error() string {
	return fmt.Sprintf("[GET /sections][%d] getSectionsOK  %+v", 200, o.Payload)
}

func (o *GetSectionsOK) String() string {
	return fmt.Sprintf("[GET /sections][%d] getSectionsOK  %+v", 200, o.Payload)
}

func (o *GetSectionsOK) GetPayload() *models.SectionsResponse {
	return o.Payload
}

func (o *GetSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SectionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
