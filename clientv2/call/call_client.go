// Code generated by go-swagger; DO NOT EDIT.

package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new call API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for call API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCalls gets app bound calls

Get app-bound calls can filter to route-bound calls, results returned in created_at, descending order (newest first).
*/
func (a *Client) GetCalls(params *GetCallsParams) (*GetCallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCalls",
		Method:             "GET",
		PathPattern:        "/calls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCallsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCallsOK), nil

}

/*
GetCallsCallID gets call information

Get call information
*/
func (a *Client) GetCallsCallID(params *GetCallsCallIDParams) (*GetCallsCallIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCallsCallIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCallsCallID",
		Method:             "GET",
		PathPattern:        "/calls/{callID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCallsCallIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCallsCallIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
