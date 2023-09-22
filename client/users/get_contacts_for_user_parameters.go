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
	"github.com/go-openapi/swag"
)

// NewGetContactsForUserParams creates a new GetContactsForUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactsForUserParams() *GetContactsForUserParams {
	return &GetContactsForUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactsForUserParamsWithTimeout creates a new GetContactsForUserParams object
// with the ability to set a timeout on a request.
func NewGetContactsForUserParamsWithTimeout(timeout time.Duration) *GetContactsForUserParams {
	return &GetContactsForUserParams{
		timeout: timeout,
	}
}

// NewGetContactsForUserParamsWithContext creates a new GetContactsForUserParams object
// with the ability to set a context for a request.
func NewGetContactsForUserParamsWithContext(ctx context.Context) *GetContactsForUserParams {
	return &GetContactsForUserParams{
		Context: ctx,
	}
}

// NewGetContactsForUserParamsWithHTTPClient creates a new GetContactsForUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactsForUserParamsWithHTTPClient(client *http.Client) *GetContactsForUserParams {
	return &GetContactsForUserParams{
		HTTPClient: client,
	}
}

/*
GetContactsForUserParams contains all the parameters to send to the API endpoint

	for the get contacts for user operation.

	Typically these are written to a http.Request.
*/
type GetContactsForUserParams struct {

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

// WithDefaults hydrates default values in the get contacts for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsForUserParams) WithDefaults() *GetContactsForUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contacts for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactsForUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contacts for user params
func (o *GetContactsForUserParams) WithTimeout(timeout time.Duration) *GetContactsForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contacts for user params
func (o *GetContactsForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contacts for user params
func (o *GetContactsForUserParams) WithContext(ctx context.Context) *GetContactsForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contacts for user params
func (o *GetContactsForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contacts for user params
func (o *GetContactsForUserParams) WithHTTPClient(client *http.Client) *GetContactsForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contacts for user params
func (o *GetContactsForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get contacts for user params
func (o *GetContactsForUserParams) WithEndingBefore(endingBefore *string) *GetContactsForUserParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get contacts for user params
func (o *GetContactsForUserParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get contacts for user params
func (o *GetContactsForUserParams) WithID(id string) *GetContactsForUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get contacts for user params
func (o *GetContactsForUserParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get contacts for user params
func (o *GetContactsForUserParams) WithLimit(limit *int64) *GetContactsForUserParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get contacts for user params
func (o *GetContactsForUserParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get contacts for user params
func (o *GetContactsForUserParams) WithStartingAfter(startingAfter *string) *GetContactsForUserParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get contacts for user params
func (o *GetContactsForUserParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactsForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
