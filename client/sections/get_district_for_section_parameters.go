// Code generated by go-swagger; DO NOT EDIT.

package sections

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

// NewGetDistrictForSectionParams creates a new GetDistrictForSectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictForSectionParams() *GetDistrictForSectionParams {
	return &GetDistrictForSectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictForSectionParamsWithTimeout creates a new GetDistrictForSectionParams object
// with the ability to set a timeout on a request.
func NewGetDistrictForSectionParamsWithTimeout(timeout time.Duration) *GetDistrictForSectionParams {
	return &GetDistrictForSectionParams{
		timeout: timeout,
	}
}

// NewGetDistrictForSectionParamsWithContext creates a new GetDistrictForSectionParams object
// with the ability to set a context for a request.
func NewGetDistrictForSectionParamsWithContext(ctx context.Context) *GetDistrictForSectionParams {
	return &GetDistrictForSectionParams{
		Context: ctx,
	}
}

// NewGetDistrictForSectionParamsWithHTTPClient creates a new GetDistrictForSectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictForSectionParamsWithHTTPClient(client *http.Client) *GetDistrictForSectionParams {
	return &GetDistrictForSectionParams{
		HTTPClient: client,
	}
}

/*
GetDistrictForSectionParams contains all the parameters to send to the API endpoint

	for the get district for section operation.

	Typically these are written to a http.Request.
*/
type GetDistrictForSectionParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get district for section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForSectionParams) WithDefaults() *GetDistrictForSectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get district for section params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForSectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get district for section params
func (o *GetDistrictForSectionParams) WithTimeout(timeout time.Duration) *GetDistrictForSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get district for section params
func (o *GetDistrictForSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get district for section params
func (o *GetDistrictForSectionParams) WithContext(ctx context.Context) *GetDistrictForSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get district for section params
func (o *GetDistrictForSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get district for section params
func (o *GetDistrictForSectionParams) WithHTTPClient(client *http.Client) *GetDistrictForSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get district for section params
func (o *GetDistrictForSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get district for section params
func (o *GetDistrictForSectionParams) WithID(id string) *GetDistrictForSectionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get district for section params
func (o *GetDistrictForSectionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictForSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
