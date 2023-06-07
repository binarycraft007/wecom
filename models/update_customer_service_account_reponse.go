// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateCustomerServiceAccountReponse update customer service account reponse
//
// swagger:model UpdateCustomerServiceAccountReponse
type UpdateCustomerServiceAccountReponse struct {

	// errcode
	// Required: true
	Errcode *int32 `json:"errcode"`

	// errmsg
	// Required: true
	Errmsg *string `json:"errmsg"`
}

// Validate validates this update customer service account reponse
func (m *UpdateCustomerServiceAccountReponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrcode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrmsg(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCustomerServiceAccountReponse) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

func (m *UpdateCustomerServiceAccountReponse) validateErrmsg(formats strfmt.Registry) error {

	if err := validate.Required("errmsg", "body", m.Errmsg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update customer service account reponse based on context it is used
func (m *UpdateCustomerServiceAccountReponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCustomerServiceAccountReponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCustomerServiceAccountReponse) UnmarshalBinary(b []byte) error {
	var res UpdateCustomerServiceAccountReponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}