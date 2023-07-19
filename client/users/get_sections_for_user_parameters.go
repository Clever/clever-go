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

// NewGetSectionsForUserParams creates a new GetSectionsForUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSectionsForUserParams() *GetSectionsForUserParams {
	return &GetSectionsForUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSectionsForUserParamsWithTimeout creates a new GetSectionsForUserParams object
// with the ability to set a timeout on a request.
func NewGetSectionsForUserParamsWithTimeout(timeout time.Duration) *GetSectionsForUserParams {
	return &GetSectionsForUserParams{
		timeout: timeout,
	}
}

// NewGetSectionsForUserParamsWithContext creates a new GetSectionsForUserParams object
// with the ability to set a context for a request.
func NewGetSectionsForUserParamsWithContext(ctx context.Context) *GetSectionsForUserParams {
	return &GetSectionsForUserParams{
		Context: ctx,
	}
}

// NewGetSectionsForUserParamsWithHTTPClient creates a new GetSectionsForUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSectionsForUserParamsWithHTTPClient(client *http.Client) *GetSectionsForUserParams {
	return &GetSectionsForUserParams{
		HTTPClient: client,
	}
}

/* GetSectionsForUserParams contains all the parameters to send to the API endpoint
   for the get sections for user operation.

   Typically these are written to a http.Request.
*/
type GetSectionsForUserParams struct {

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

// WithDefaults hydrates default values in the get sections for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSectionsForUserParams) WithDefaults() *GetSectionsForUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sections for user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSectionsForUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sections for user params
func (o *GetSectionsForUserParams) WithTimeout(timeout time.Duration) *GetSectionsForUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sections for user params
func (o *GetSectionsForUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sections for user params
func (o *GetSectionsForUserParams) WithContext(ctx context.Context) *GetSectionsForUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sections for user params
func (o *GetSectionsForUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sections for user params
func (o *GetSectionsForUserParams) WithHTTPClient(client *http.Client) *GetSectionsForUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sections for user params
func (o *GetSectionsForUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingBefore adds the endingBefore to the get sections for user params
func (o *GetSectionsForUserParams) WithEndingBefore(endingBefore *string) *GetSectionsForUserParams {
	o.SetEndingBefore(endingBefore)
	return o
}

// SetEndingBefore adds the endingBefore to the get sections for user params
func (o *GetSectionsForUserParams) SetEndingBefore(endingBefore *string) {
	o.EndingBefore = endingBefore
}

// WithID adds the id to the get sections for user params
func (o *GetSectionsForUserParams) WithID(id string) *GetSectionsForUserParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get sections for user params
func (o *GetSectionsForUserParams) SetID(id string) {
	o.ID = id
}

// WithLimit adds the limit to the get sections for user params
func (o *GetSectionsForUserParams) WithLimit(limit *int64) *GetSectionsForUserParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get sections for user params
func (o *GetSectionsForUserParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithStartingAfter adds the startingAfter to the get sections for user params
func (o *GetSectionsForUserParams) WithStartingAfter(startingAfter *string) *GetSectionsForUserParams {
	o.SetStartingAfter(startingAfter)
	return o
}

// SetStartingAfter adds the startingAfter to the get sections for user params
func (o *GetSectionsForUserParams) SetStartingAfter(startingAfter *string) {
	o.StartingAfter = startingAfter
}

// WriteToRequest writes these params to a swagger request
func (o *GetSectionsForUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
