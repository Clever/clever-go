// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewGetDistrictForUserParams creates a new GetDistrictForUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDistrictForUserParams() *GetDistrictForUserParams {
	return &GetDistrictForUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDistrictForUserParamsWithTimeout creates a new GetDistrictForUserParams object
// with the ability to set a timeout on a request.
func NewGetDistrictForUserParamsWithTimeout(timeout time.Duration) *GetDistrictForUserParams {
	return &GetDistrictForUserParams{
		timeout: timeout,
	}
}

// NewGetDistrictForUserParamsWithContext creates a new GetDistrictForUserParams object
// with the ability to set a context for a request.
func NewGetDistrictForUserParamsWithContext(ctx context.Context) *GetDistrictForUserParams {
	return &GetDistrictForUserParams{
		Context: ctx,
	}
}

// NewGetDistrictForUserParamsWithHTTPClient creates a new GetDistrictForUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDistrictForUserParamsWithHTTPClient(client *http.Client) *GetDistrictForUserParams {
	return &GetDistrictForUserParams{
		HTTPClient: client,
	}
}

/* GetDistrictForUserParams contains all the parameters to send to the API endpoint
   for the get district for user operation.

   Typically these are written to a http.Request.
*/
type GetDistrictForUserParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get district for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForUserParams) WithDefaults() *GetDistrictForUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get district for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDistrictForUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get district for user params
func (o *GetDistrictForUserParams) WithTimeout(timeout time.Duration) *GetDistrictForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get district for user params
func (o *GetDistrictForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get district for user params
func (o *GetDistrictForUserParams) WithContext(ctx context.Context) *GetDistrictForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get district for user params
func (o *GetDistrictForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get district for user params
func (o *GetDistrictForUserParams) WithHTTPClient(client *http.Client) *GetDistrictForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get district for user params
func (o *GetDistrictForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get district for user params
func (o *GetDistrictForUserParams) WithID(id string) *GetDistrictForUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get district for user params
func (o *GetDistrictForUserParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDistrictForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
