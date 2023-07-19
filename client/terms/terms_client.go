// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new terms API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for terms API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDistrictForTerm(params *GetDistrictForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistrictForTermOK, error)

	GetSchoolsForTerm(params *GetSchoolsForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolsForTermOK, error)

	GetSectionsForTerm(params *GetSectionsForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSectionsForTermOK, error)

	GetTerm(params *GetTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermOK, error)

	GetTerms(params *GetTermsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDistrictForTerm Returns the district for a term
*/
func (a *Client) GetDistrictForTerm(params *GetDistrictForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDistrictForTermOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistrictForTermParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDistrictForTerm",
		Method:             "GET",
		PathPattern:        "/terms/{id}/district",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDistrictForTermReader{formats: a.formats},
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
	success, ok := result.(*GetDistrictForTermOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDistrictForTerm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSchoolsForTerm Returns the schools for a term
*/
func (a *Client) GetSchoolsForTerm(params *GetSchoolsForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSchoolsForTermOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchoolsForTermParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSchoolsForTerm",
		Method:             "GET",
		PathPattern:        "/terms/{id}/schools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchoolsForTermReader{formats: a.formats},
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
	success, ok := result.(*GetSchoolsForTermOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchoolsForTerm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSectionsForTerm Returns the sections for a term
*/
func (a *Client) GetSectionsForTerm(params *GetSectionsForTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSectionsForTermOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectionsForTermParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSectionsForTerm",
		Method:             "GET",
		PathPattern:        "/terms/{id}/sections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectionsForTermReader{formats: a.formats},
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
	success, ok := result.(*GetSectionsForTermOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSectionsForTerm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTerm Returns a specific term
*/
func (a *Client) GetTerm(params *GetTermParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTermParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTerm",
		Method:             "GET",
		PathPattern:        "/terms/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTermReader{formats: a.formats},
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
	success, ok := result.(*GetTermOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTerm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTerms Returns a list of terms
*/
func (a *Client) GetTerms(params *GetTermsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTermsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTermsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTerms",
		Method:             "GET",
		PathPattern:        "/terms",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTermsReader{formats: a.formats},
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
	success, ok := result.(*GetTermsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTerms: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
