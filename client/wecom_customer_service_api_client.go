// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/binarycraft007/wecom/client/account"
	"github.com/binarycraft007/wecom/client/message"
	"github.com/binarycraft007/wecom/client/service_state"
	"github.com/binarycraft007/wecom/client/token"
)

// Default wecom customer service API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "qyapi.weixin.qq.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/cgi-bin"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new wecom customer service API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *WecomCustomerServiceAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new wecom customer service API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *WecomCustomerServiceAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new wecom customer service API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *WecomCustomerServiceAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(WecomCustomerServiceAPI)
	cli.Transport = transport
	cli.Account = account.New(transport, formats)
	cli.Message = message.New(transport, formats)
	cli.ServiceState = service_state.New(transport, formats)
	cli.Token = token.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// WecomCustomerServiceAPI is a client for wecom customer service API
type WecomCustomerServiceAPI struct {
	Account account.ClientService

	Message message.ClientService

	ServiceState service_state.ClientService

	Token token.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *WecomCustomerServiceAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Message.SetTransport(transport)
	c.ServiceState.SetTransport(transport)
	c.Token.SetTransport(transport)
}
