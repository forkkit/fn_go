// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new apps API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for apps API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateApp posts new app

Insert a new app
*/
func (a *Client) CreateApp(params *CreateAppParams) (*CreateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateApp",
		Method:             "POST",
		PathPattern:        "/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAppOK), nil

}

/*
DeleteApp deletes an app

Delete an app.
*/
func (a *Client) DeleteApp(params *DeleteAppParams) (*DeleteAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteApp",
		Method:             "DELETE",
		PathPattern:        "/apps/{appID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppOK), nil

}

/*
GetApp gets information for a app

This gives more details about a app, such as statistics.
*/
func (a *Client) GetApp(params *GetAppParams) (*GetAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetApp",
		Method:             "GET",
		PathPattern:        "/apps/{appID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppOK), nil

}

/*
ListApps gets applications

Get a filtered applications returned in alphabetical order.
*/
func (a *Client) ListApps(params *ListAppsParams) (*ListAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAppsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListApps",
		Method:             "GET",
		PathPattern:        "/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAppsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAppsOK), nil

}

/*
UpdateApp updates an app

Updates and application.
*/
func (a *Client) UpdateApp(params *UpdateAppParams) (*UpdateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateApp",
		Method:             "PUT",
		PathPattern:        "/apps/{appID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateAppOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
