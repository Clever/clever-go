// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetCoursesParams creates a new GetCoursesParams object
//
// There are no default values defined in the spec.
func NewGetCoursesParams() GetCoursesParams {

	return GetCoursesParams{}
}

// GetCoursesParams contains all the bound params for the get courses operation
// typically these are obtained from a http.Request
//
// swagger:parameters getCourses
type GetCoursesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Count *string
	/*
	  In: query
	*/
	EndingBefore *string
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
// To ensure default values, the struct must have been initialized with NewGetCoursesParams() beforehand.
func (o *GetCoursesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCount, qhkCount, _ := qs.GetOK("count")
	if err := o.bindCount(qCount, qhkCount, route.Formats); err != nil {
		res = append(res, err)
	}

	qEndingBefore, qhkEndingBefore, _ := qs.GetOK("ending_before")
	if err := o.bindEndingBefore(qEndingBefore, qhkEndingBefore, route.Formats); err != nil {
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

// bindCount binds and validates parameter Count from query.
func (o *GetCoursesParams) bindCount(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Count = &raw

	if err := o.validateCount(formats); err != nil {
		return err
	}

	return nil
}

// validateCount carries on validations for parameter Count
func (o *GetCoursesParams) validateCount(formats strfmt.Registry) error {

	if err := validate.EnumCase("count", "query", *o.Count, []interface{}{"", "true"}, true); err != nil {
		return err
	}

	return nil
}

// bindEndingBefore binds and validates parameter EndingBefore from query.
func (o *GetCoursesParams) bindEndingBefore(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindLimit binds and validates parameter Limit from query.
func (o *GetCoursesParams) bindLimit(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetCoursesParams) bindStartingAfter(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
