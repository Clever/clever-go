// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetEvent(params *GetEventParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventOK, error)

	GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetEvent Returns the specific event
*/
func (a *Client) GetEvent(params *GetEventParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEvent",
		Method:             "GET",
		PathPattern:        "/events/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEvent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEvents Returns a list of events
*/
func (a *Client) GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEvents",
		Method:             "GET",
		PathPattern:        "/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEventsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEvents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
