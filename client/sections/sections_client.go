// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sections API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sections API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetCourseForSection(params *GetCourseForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetCourseForSectionOK, error)

	GetDistrictForSection(params *GetDistrictForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetDistrictForSectionOK, error)

	GetResourcesForSection(params *GetResourcesForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetResourcesForSectionOK, error)

	GetSchoolForSection(params *GetSchoolForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSchoolForSectionOK, error)

	GetSection(params *GetSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectionOK, error)

	GetSections(params *GetSectionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectionsOK, error)

	GetTermForSection(params *GetTermForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetTermForSectionOK, error)

	GetUsersForSection(params *GetUsersForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersForSectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetCourseForSection Returns the course for a section
*/
func (a *Client) GetCourseForSection(params *GetCourseForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetCourseForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCourseForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCourseForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/course",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCourseForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCourseForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCourseForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDistrictForSection Returns the district for a section
*/
func (a *Client) GetDistrictForSection(params *GetDistrictForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetDistrictForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDistrictForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDistrictForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/district",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDistrictForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDistrictForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDistrictForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetResourcesForSection Returns the resources for a section
*/
func (a *Client) GetResourcesForSection(params *GetResourcesForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetResourcesForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourcesForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResourcesForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourcesForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourcesForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getResourcesForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSchoolForSection Returns the school for a section
*/
func (a *Client) GetSchoolForSection(params *GetSchoolForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSchoolForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSchoolForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSchoolForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/school",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSchoolForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSchoolForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSchoolForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSection Returns a specific section
*/
func (a *Client) GetSection(params *GetSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSections Returns a list of sections
*/
func (a *Client) GetSections(params *GetSectionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSections",
		Method:             "GET",
		PathPattern:        "/sections",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTermForSection Returns the term for a section
*/
func (a *Client) GetTermForSection(params *GetTermForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetTermForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTermForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTermForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/term",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTermForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTermForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTermForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersForSection Returns the student and/or teacher users for a section
*/
func (a *Client) GetUsersForSection(params *GetUsersForSectionParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersForSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersForSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersForSection",
		Method:             "GET",
		PathPattern:        "/sections/{id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersForSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersForSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsersForSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
