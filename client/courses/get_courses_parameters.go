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

// NewGetCoursesParams creates a new GetCoursesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCoursesParams() *GetCoursesParams {
	return &GetCoursesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCoursesParamsWithTimeout creates a new GetCoursesParams object
// with the ability to set a timeout on a request.
func NewGetCoursesParamsWithTimeout(timeout time.Duration) *GetCoursesParams {
	return &GetCoursesParams{
		timeout: timeout,
	}
}

// NewGetCoursesParamsWithContext creates a new GetCoursesParams object
// with the ability to set a context for a request.
func NewGetCoursesParamsWithContext(ctx context.Context) *GetCoursesParams {
	return &GetCoursesParams{
		Context: ctx,
	}
}

// NewGetCoursesParamsWithHTTPClient creates a new GetCoursesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCoursesParamsWithHTTPClient(client *http.Client) *GetCoursesParams {
	return &GetCoursesParams{
		HTTPClient: client,
	}
}

/* GetCoursesParams contains all the parameters to send to the API endpoint
   for the get courses operation.

   Typically these are written to a http.Request.
*/
type GetCoursesParams struct {

	// Count.
	Count *string

	// EndingBefore.
	EndingBefore *string

	// Limit.
	Limit *int64

	// StartingAfter.
	StartingAfter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get courses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoursesParams) WithDefaults() *GetCoursesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get courses params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCoursesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get courses params
func (o *GetCoursesParams) WithTimeout(timeout time.Duration) *GetCoursesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get courses params
func (o *GetCoursesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get courses params
func (o *GetCoursesParams) WithContext(ctx context.Context) *GetCoursesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get courses params
func (o *GetCoursesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get courses params
func (o *GetCoursesParams) WithHTTPClient(client *http.Client) *GetCoursesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get courses params
func (o *GetCoursesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get courses params
func (o *GetCoursesParams) WithCount(count *string) *GetCoursesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get courses params
func (o *GetCoursesParams) SetCount(count *string) {
	o.Count = count
}

// WithEndingBefore adds the endingBefore to the get courses params
func (o *GetCoursesParams) WithEndingBefore(endingBefore *string) *GetCoursesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get courses params
func (o *GetCoursesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithLimit adds the limit to the get courses params
func (o *GetCoursesParams) WithLimit(limit *int64) *GetCoursesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get courses params
func (o *GetCoursesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get courses params
func (o *GetCoursesParams) WithStartingAfter(startingAfter *string) *GetCoursesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get courses params
func (o *GetCoursesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetCoursesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount string

		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := qrCount
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

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
