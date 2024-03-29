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
	"github.com/go-openapi/swag"
)

// NewGetSchoolsForCourseParams creates a new GetSchoolsForCourseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchoolsForCourseParams() *GetSchoolsForCourseParams {
	return &GetSchoolsForCourseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchoolsForCourseParamsWithTimeout creates a new GetSchoolsForCourseParams object
// with the ability to set a timeout on a request.
func NewGetSchoolsForCourseParamsWithTimeout(timeout time.Duration) *GetSchoolsForCourseParams {
	return &GetSchoolsForCourseParams{
		timeout: timeout,
	}
}

// NewGetSchoolsForCourseParamsWithContext creates a new GetSchoolsForCourseParams object
// with the ability to set a context for a request.
func NewGetSchoolsForCourseParamsWithContext(ctx context.Context) *GetSchoolsForCourseParams {
	return &GetSchoolsForCourseParams{
		Context: ctx,
	}
}

// NewGetSchoolsForCourseParamsWithHTTPClient creates a new GetSchoolsForCourseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchoolsForCourseParamsWithHTTPClient(client *http.Client) *GetSchoolsForCourseParams {
	return &GetSchoolsForCourseParams{
		HTTPClient: client,
	}
}

/* GetSchoolsForCourseParams contains all the parameters to send to the API endpoint
   for the get schools for course operation.

   Typically these are written to a http.Request.
*/
type GetSchoolsForCourseParams struct {

	// EndingBefore.
	EndingBefore *string

	// ID.
	ID string

	// Limit.
	Limit *int64

	// StartingAfter.
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get schools for course params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchoolsForCourseParams) WithDefaults() *GetSchoolsForCourseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schools for course params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchoolsForCourseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schools for course params
func (o *GetSchoolsForCourseParams) WithTimeout(timeout time.Duration) *GetSchoolsForCourseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schools for course params
func (o *GetSchoolsForCourseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schools for course params
func (o *GetSchoolsForCourseParams) WithContext(ctx context.Context) *GetSchoolsForCourseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schools for course params
func (o *GetSchoolsForCourseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schools for course params
func (o *GetSchoolsForCourseParams) WithHTTPClient(client *http.Client) *GetSchoolsForCourseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schools for course params
func (o *GetSchoolsForCourseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get schools for course params
func (o *GetSchoolsForCourseParams) WithEndingBefore(endingBefore *string) *GetSchoolsForCourseParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get schools for course params
func (o *GetSchoolsForCourseParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get schools for course params
func (o *GetSchoolsForCourseParams) WithID(id string) *GetSchoolsForCourseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get schools for course params
func (o *GetSchoolsForCourseParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get schools for course params
func (o *GetSchoolsForCourseParams) WithLimit(limit *int64) *GetSchoolsForCourseParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get schools for course params
func (o *GetSchoolsForCourseParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get schools for course params
func (o *GetSchoolsForCourseParams) WithStartingAfter(startingAfter *string) *GetSchoolsForCourseParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get schools for course params
func (o *GetSchoolsForCourseParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchoolsForCourseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingBefore != nil {

		// query param ending_before
		var qrEndingBefore string

		if o.EndingBefore != nil {
			qrEndingBefore = *o.EndingBefore
		}
		qEndingBefore := qrEndingBefore
		if qEndingBefore != "" {

			if err := r.SetQueryParam("ending_before", qEndingBefore); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.StartingAfter != nil {

		// query param starting_after
		var qrStartingAfter string

		if o.StartingAfter != nil {
			qrStartingAfter = *o.StartingAfter
		}
		qStartingAfter := qrStartingAfter
		if qStartingAfter != "" {

			if err := r.SetQueryParam("starting_after", qStartingAfter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
