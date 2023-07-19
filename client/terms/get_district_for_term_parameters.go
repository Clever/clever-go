// Code generated by go-swagger; DO NOT EDIT.

package terms

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

// NewGetDistrictForTermParams creates a new GetDistrictForTermParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictForTermParams() *GetDistrictForTermParams {
	return &GetDistrictForTermParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictForTermParamsWithTimeout creates a new GetDistrictForTermParams object
// with the ability to set a timeout on a request.
func NewGetDistrictForTermParamsWithTimeout(timeout time.Duration) *GetDistrictForTermParams {
	return &GetDistrictForTermParams{
		timeout: timeout,
	}
}

// NewGetDistrictForTermParamsWithContext creates a new GetDistrictForTermParams object
// with the ability to set a context for a request.
func NewGetDistrictForTermParamsWithContext(ctx context.Context) *GetDistrictForTermParams {
	return &GetDistrictForTermParams{
		Context: ctx,
	}
}

// NewGetDistrictForTermParamsWithHTTPClient creates a new GetDistrictForTermParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictForTermParamsWithHTTPClient(client *http.Client) *GetDistrictForTermParams {
	return &GetDistrictForTermParams{
		HTTPClient: client,
	}
}

/* GetDistrictForTermParams contains all the parameters to send to the API endpoint
   for the get district for term operation.

   Typically these are written to a http.Request.
*/
type GetDistrictForTermParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get district for term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForTermParams) WithDefaults() *GetDistrictForTermParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get district for term params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForTermParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get district for term params
func (o *GetDistrictForTermParams) WithTimeout(timeout time.Duration) *GetDistrictForTermParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get district for term params
func (o *GetDistrictForTermParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get district for term params
func (o *GetDistrictForTermParams) WithContext(ctx context.Context) *GetDistrictForTermParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get district for term params
func (o *GetDistrictForTermParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get district for term params
func (o *GetDistrictForTermParams) WithHTTPClient(client *http.Client) *GetDistrictForTermParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get district for term params
func (o *GetDistrictForTermParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get district for term params
func (o *GetDistrictForTermParams) WithID(id string) *GetDistrictForTermParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get district for term params
func (o *GetDistrictForTermParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictForTermParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
