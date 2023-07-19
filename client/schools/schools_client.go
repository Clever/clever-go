// Code generated by go-swagger; DO NOT EDIT.

package schools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schools API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schools API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetCoursesForSchool(params *GetCoursesForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCoursesForSchoolOK, error)

	GetDistrictForSchool(params *GetDistrictForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistrictForSchoolOK, error)

	GetSchool(params *GetSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolOK, error)

	GetSchools(params *GetSchoolsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolsOK, error)

	GetSectionsForSchool(params *GetSectionsForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSectionsForSchoolOK, error)

	GetTermsForSchool(params *GetTermsForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermsForSchoolOK, error)

	GetUsersForSchool(params *GetUsersForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersForSchoolOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCoursesForSchool Returns the courses for a school
*/
func (a *Client) GetCoursesForSchool(params *GetCoursesForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCoursesForSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCoursesForSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCoursesForSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}/courses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCoursesForSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCoursesForSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCoursesForSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDistrictForSchool Returns the district for a school
*/
func (a *Client) GetDistrictForSchool(params *GetDistrictForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistrictForSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistrictForSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDistrictForSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}/district",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDistrictForSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDistrictForSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDistrictForSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSchool Returns a specific school
*/
func (a *Client) GetSchool(params *GetSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSchools Returns a list of schools
*/
func (a *Client) GetSchools(params *GetSchoolsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchoolsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSchools",
		Method:             "GET",
		PathPattern:        "/schools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchoolsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchoolsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchools: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSectionsForSchool Returns the sections for a school
*/
func (a *Client) GetSectionsForSchool(params *GetSectionsForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSectionsForSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectionsForSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSectionsForSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}/sections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectionsForSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSectionsForSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSectionsForSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTermsForSchool Returns the terms for a school
*/
func (a *Client) GetTermsForSchool(params *GetTermsForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermsForSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTermsForSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTermsForSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}/terms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTermsForSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTermsForSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTermsForSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersForSchool Returns the staff, student, and/or teacher users for a school
*/
func (a *Client) GetUsersForSchool(params *GetUsersForSchoolParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersForSchoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersForSchoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsersForSchool",
		Method:             "GET",
		PathPattern:        "/schools/{id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersForSchoolReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersForSchoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsersForSchool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
