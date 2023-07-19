// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetSectionsForResourceParams creates a new GetSectionsForResourceParams object
//
// There are no default values defined in the spec.
func NewGetSectionsForResourceParams() GetSectionsForResourceParams {

	return GetSectionsForResourceParams{}
}

// GetSectionsForResourceParams contains all the bound params for the get sections for resource operation
// typically these are obtained from a http.Request
//
// swagger:parameters getSectionsForResource
type GetSectionsForResourceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	EndingBefore *string
	/*
	  Required: true
	  In: path
	*/
	ID string
	/*
	  In: query
	*/
	Limit *int64
	/*
	  In: query
	*/
	StartingAfter *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetSectionsForResourceParams() beforehand.
func (o *GetSectionsForResourceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qEndingBefore, qhkEndingBefore, _ := qs.GetOK("ending_before")
	if err := o.bindEndingBefore(qEndingBefore, qhkEndingBefore, route.Formats); err != nil {
		res = append(res, err)
	}

	rID, rhkID, _ := route.Params.GetOK("id")
	if err := o.bindID(rID, rhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qLimit, qhkLimit, _ := qs.GetOK("limit")
	if err := o.bindLimit(qLimit, qhkLimit, route.Formats); err != nil {
		res = append(res, err)
	}

	qStartingAfter, qhkStartingAfter, _ := qs.GetOK("starting_after")
	if err := o.bindStartingAfter(qStartingAfter, qhkStartingAfter, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindEndingBefore binds and validates parameter EndingBefore from query.
func (o *GetSectionsForResourceParams) bindEndingBefore(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.EndingBefore = &raw

	return nil
}

// bindID binds and validates parameter ID from path.
func (o *GetSectionsForResourceParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.ID = raw

	return nil
}

// bindLimit binds and validates parameter Limit from query.
func (o *GetSectionsForResourceParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("limit", "query", "int64", raw)
	}
	o.Limit = &value

	return nil
}

// bindStartingAfter binds and validates parameter StartingAfter from query.
func (o *GetSectionsForResourceParams) bindStartingAfter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.StartingAfter = &raw

	return nil
}
