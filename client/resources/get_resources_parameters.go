// Code generated by go-swagger; DO NOT EDIT.

package resources

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

// NewGetResourcesParams creates a new GetResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResourcesParams() *GetResourcesParams {
	return &GetResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourcesParamsWithTimeout creates a new GetResourcesParams object
// with the ability to set a timeout on a request.
func NewGetResourcesParamsWithTimeout(timeout time.Duration) *GetResourcesParams {
	return &GetResourcesParams{
		timeout: timeout,
	}
}

// NewGetResourcesParamsWithContext creates a new GetResourcesParams object
// with the ability to set a context for a request.
func NewGetResourcesParamsWithContext(ctx context.Context) *GetResourcesParams {
	return &GetResourcesParams{
		Context: ctx,
	}
}

// NewGetResourcesParamsWithHTTPClient creates a new GetResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResourcesParamsWithHTTPClient(client *http.Client) *GetResourcesParams {
	return &GetResourcesParams{
		HTTPClient: client,
	}
}

/* GetResourcesParams contains all the parameters to send to the API endpoint
   for the get resources operation.

   Typically these are written to a http.Request.
*/
type GetResourcesParams struct {

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

// WithDefaults hydrates default values in the get resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourcesParams) WithDefaults() *GetResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResourcesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get resources params
func (o *GetResourcesParams) WithTimeout(timeout time.Duration) *GetResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resources params
func (o *GetResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resources params
func (o *GetResourcesParams) WithContext(ctx context.Context) *GetResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resources params
func (o *GetResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resources params
func (o *GetResourcesParams) WithHTTPClient(client *http.Client) *GetResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resources params
func (o *GetResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get resources params
func (o *GetResourcesParams) WithEndingBefore(endingBefore *string) *GetResourcesParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get resources params
func (o *GetResourcesParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithLimit adds the limit to the get resources params
func (o *GetResourcesParams) WithLimit(limit *int64) *GetResourcesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get resources params
func (o *GetResourcesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get resources params
func (o *GetResourcesParams) WithStartingAfter(startingAfter *string) *GetResourcesParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get resources params
func (o *GetResourcesParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
