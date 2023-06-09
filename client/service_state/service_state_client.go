// Code generated by go-swagger; DO NOT EDIT.

package service_state

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service state API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service state API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetState(params *GetStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStateOK, error)

	Transform(params *TransformParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TransformOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetState get service state
*/
func (a *Client) GetState(params *GetStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getState",
		Method:             "POST",
		PathPattern:        "/kf/service_state/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStateReader{formats: a.formats},
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
	success, ok := result.(*GetStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetStateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Transform transform service state
*/
func (a *Client) Transform(params *TransformParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TransformOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTransformParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "transform",
		Method:             "POST",
		PathPattern:        "/kf/service_state/trans",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TransformReader{formats: a.formats},
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
	success, ok := result.(*TransformOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TransformDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
