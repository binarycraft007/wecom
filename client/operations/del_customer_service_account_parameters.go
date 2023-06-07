// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/binarycraft007/wecom/models"
)

// NewDelCustomerServiceAccountParams creates a new DelCustomerServiceAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDelCustomerServiceAccountParams() *DelCustomerServiceAccountParams {
	return &DelCustomerServiceAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDelCustomerServiceAccountParamsWithTimeout creates a new DelCustomerServiceAccountParams object
// with the ability to set a timeout on a request.
func NewDelCustomerServiceAccountParamsWithTimeout(timeout time.Duration) *DelCustomerServiceAccountParams {
	return &DelCustomerServiceAccountParams{
		timeout: timeout,
	}
}

// NewDelCustomerServiceAccountParamsWithContext creates a new DelCustomerServiceAccountParams object
// with the ability to set a context for a request.
func NewDelCustomerServiceAccountParamsWithContext(ctx context.Context) *DelCustomerServiceAccountParams {
	return &DelCustomerServiceAccountParams{
		Context: ctx,
	}
}

// NewDelCustomerServiceAccountParamsWithHTTPClient creates a new DelCustomerServiceAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewDelCustomerServiceAccountParamsWithHTTPClient(client *http.Client) *DelCustomerServiceAccountParams {
	return &DelCustomerServiceAccountParams{
		HTTPClient: client,
	}
}

/*
DelCustomerServiceAccountParams contains all the parameters to send to the API endpoint

	for the del customer service account operation.

	Typically these are written to a http.Request.
*/
type DelCustomerServiceAccountParams struct {

	/* DelCustomerServiceAccountRequest.

	   delete customer service account request
	*/
	DelCustomerServiceAccountRequest *models.DelCustomerServiceAccountRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the del customer service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelCustomerServiceAccountParams) WithDefaults() *DelCustomerServiceAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the del customer service account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelCustomerServiceAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the del customer service account params
func (o *DelCustomerServiceAccountParams) WithTimeout(timeout time.Duration) *DelCustomerServiceAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the del customer service account params
func (o *DelCustomerServiceAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the del customer service account params
func (o *DelCustomerServiceAccountParams) WithContext(ctx context.Context) *DelCustomerServiceAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the del customer service account params
func (o *DelCustomerServiceAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the del customer service account params
func (o *DelCustomerServiceAccountParams) WithHTTPClient(client *http.Client) *DelCustomerServiceAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the del customer service account params
func (o *DelCustomerServiceAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelCustomerServiceAccountRequest adds the delCustomerServiceAccountRequest to the del customer service account params
func (o *DelCustomerServiceAccountParams) WithDelCustomerServiceAccountRequest(delCustomerServiceAccountRequest *models.DelCustomerServiceAccountRequest) *DelCustomerServiceAccountParams {
	o.SetDelCustomerServiceAccountRequest(delCustomerServiceAccountRequest)
	return o
}

// SetDelCustomerServiceAccountRequest adds the delCustomerServiceAccountRequest to the del customer service account params
func (o *DelCustomerServiceAccountParams) SetDelCustomerServiceAccountRequest(delCustomerServiceAccountRequest *models.DelCustomerServiceAccountRequest) {
	o.DelCustomerServiceAccountRequest = delCustomerServiceAccountRequest
}

// WriteToRequest writes these params to a swagger request
func (o *DelCustomerServiceAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.DelCustomerServiceAccountRequest != nil {
		if err := r.SetBodyParam(o.DelCustomerServiceAccountRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}