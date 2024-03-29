// Code generated by go-swagger; DO NOT EDIT.

package schools

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

// NewGetSchoolsParams creates a new GetSchoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSchoolsParams() *GetSchoolsParams {
	return &GetSchoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSchoolsParamsWithTimeout creates a new GetSchoolsParams object
// with the ability to set a timeout on a request.
func NewGetSchoolsParamsWithTimeout(timeout time.Duration) *GetSchoolsParams {
	return &GetSchoolsParams{
		timeout: timeout,
	}
}

// NewGetSchoolsParamsWithContext creates a new GetSchoolsParams object
// with the ability to set a context for a request.
func NewGetSchoolsParamsWithContext(ctx context.Context) *GetSchoolsParams {
	return &GetSchoolsParams{
		Context: ctx,
	}
}

// NewGetSchoolsParamsWithHTTPClient creates a new GetSchoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSchoolsParamsWithHTTPClient(client *http.Client) *GetSchoolsParams {
	return &GetSchoolsParams{
		HTTPClient: client,
	}
}

/* GetSchoolsParams contains all the parameters to send to the API endpoint
   for the get schools operation.

   Typically these are written to a http.Request.
*/
type GetSchoolsParams struct {

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

// WithDefaults hydrates default values in the get schools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchoolsParams) WithDefaults() *GetSchoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get schools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSchoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get schools params
func (o *GetSchoolsParams) WithTimeout(timeout time.Duration) *GetSchoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get schools params
func (o *GetSchoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get schools params
func (o *GetSchoolsParams) WithContext(ctx context.Context) *GetSchoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get schools params
func (o *GetSchoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get schools params
func (o *GetSchoolsParams) WithHTTPClient(client *http.Client) *GetSchoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get schools params
func (o *GetSchoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get schools params
func (o *GetSchoolsParams) WithCount(count *string) *GetSchoolsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get schools params
func (o *GetSchoolsParams) SetCount(count *string) {
	o.Count = count
}

// WithEndingBefore adds the endingBefore to the get schools params
func (o *GetSchoolsParams) WithEndingBefore(endingBefore *string) *GetSchoolsParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get schools params
func (o *GetSchoolsParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithLimit adds the limit to the get schools params
func (o *GetSchoolsParams) WithLimit(limit *int64) *GetSchoolsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get schools params
func (o *GetSchoolsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get schools params
func (o *GetSchoolsParams) WithStartingAfter(startingAfter *string) *GetSchoolsParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get schools params
func (o *GetSchoolsParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetSchoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
