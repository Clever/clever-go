// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetDistrictForCourseParams creates a new GetDistrictForCourseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictForCourseParams() *GetDistrictForCourseParams {
	return &GetDistrictForCourseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictForCourseParamsWithTimeout creates a new GetDistrictForCourseParams object
// with the ability to set a timeout on a request.
func NewGetDistrictForCourseParamsWithTimeout(timeout time.Duration) *GetDistrictForCourseParams {
	return &GetDistrictForCourseParams{
		timeout: timeout,
	}
}

// NewGetDistrictForCourseParamsWithContext creates a new GetDistrictForCourseParams object
// with the ability to set a context for a request.
func NewGetDistrictForCourseParamsWithContext(ctx context.Context) *GetDistrictForCourseParams {
	return &GetDistrictForCourseParams{
		Context: ctx,
	}
}

// NewGetDistrictForCourseParamsWithHTTPClient creates a new GetDistrictForCourseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictForCourseParamsWithHTTPClient(client *http.Client) *GetDistrictForCourseParams {
	return &GetDistrictForCourseParams{
		HTTPClient: client,
	}
}

/* GetDistrictForCourseParams contains all the parameters to send to the API endpoint
   for the get district for course operation.

   Typically these are written to a http.Request.
*/
type GetDistrictForCourseParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get district for course params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForCourseParams) WithDefaults() *GetDistrictForCourseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get district for course params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForCourseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get district for course params
func (o *GetDistrictForCourseParams) WithTimeout(timeout time.Duration) *GetDistrictForCourseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get district for course params
func (o *GetDistrictForCourseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get district for course params
func (o *GetDistrictForCourseParams) WithContext(ctx context.Context) *GetDistrictForCourseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get district for course params
func (o *GetDistrictForCourseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get district for course params
func (o *GetDistrictForCourseParams) WithHTTPClient(client *http.Client) *GetDistrictForCourseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get district for course params
func (o *GetDistrictForCourseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get district for course params
func (o *GetDistrictForCourseParams) WithID(id string) *GetDistrictForCourseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get district for course params
func (o *GetDistrictForCourseParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictForCourseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
