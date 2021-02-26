// Code generated by go-swagger; DO NOT EDIT.

package districts

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

// NewGetDistrictsParams creates a new GetDistrictsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictsParams() *GetDistrictsParams {
	return &GetDistrictsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictsParamsWithTimeout creates a new GetDistrictsParams object
// with the ability to set a timeout on a request.
func NewGetDistrictsParamsWithTimeout(timeout time.Duration) *GetDistrictsParams {
	return &GetDistrictsParams{
		timeout: timeout,
	}
}

// NewGetDistrictsParamsWithContext creates a new GetDistrictsParams object
// with the ability to set a context for a request.
func NewGetDistrictsParamsWithContext(ctx context.Context) *GetDistrictsParams {
	return &GetDistrictsParams{
		Context: ctx,
	}
}

// NewGetDistrictsParamsWithHTTPClient creates a new GetDistrictsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictsParamsWithHTTPClient(client *http.Client) *GetDistrictsParams {
	return &GetDistrictsParams{
		HTTPClient: client,
	}
}

/* GetDistrictsParams contains all the parameters to send to the API endpoint
   for the get districts operation.

   Typically these are written to a http.Request.
*/
type GetDistrictsParams struct {

	// Count.
	Count *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get districts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictsParams) WithDefaults() *GetDistrictsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get districts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get districts params
func (o *GetDistrictsParams) WithTimeout(timeout time.Duration) *GetDistrictsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get districts params
func (o *GetDistrictsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get districts params
func (o *GetDistrictsParams) WithContext(ctx context.Context) *GetDistrictsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get districts params
func (o *GetDistrictsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get districts params
func (o *GetDistrictsParams) WithHTTPClient(client *http.Client) *GetDistrictsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get districts params
func (o *GetDistrictsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the get districts params
func (o *GetDistrictsParams) WithCount(count *string) *GetDistrictsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get districts params
func (o *GetDistrictsParams) SetCount(count *string) {
	o.Count = count
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}