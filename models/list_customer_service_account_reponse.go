// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListCustomerServiceAccountReponse list customer service account reponse
//
// swagger:model ListCustomerServiceAccountReponse
type ListCustomerServiceAccountReponse struct {

	// account list
	AccountList []*CustomerServiceAccount `json:"account_list"`

	// errcode
	// Required: true
	Errcode *int32 `json:"errcode"`

	// errmsg
	// Required: true
	Errmsg *string `json:"errmsg"`
}

// Validate validates this list customer service account reponse
func (m *ListCustomerServiceAccountReponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountList(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ListCustomerServiceAccountReponse) validateAccountList(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountList) { // not required
		return nil
	}

	for i := 0; i < len(m.AccountList); i++ {
		if swag.IsZero(m.AccountList[i]) { // not required
			continue
		}

		if m.AccountList[i] != nil {
			if err := m.AccountList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("account_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("account_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ListCustomerServiceAccountReponse) validateErrcode(formats strfmt.Registry) error {

	if err := validate.Required("errcode", "body", m.Errcode); err != nil {
		return err
	}

	return nil
}

func (m *ListCustomerServiceAccountReponse) validateErrmsg(formats strfmt.Registry) error {

	if err := validate.Required("errmsg", "body", m.Errmsg); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this list customer service account reponse based on the context it is used
func (m *ListCustomerServiceAccountReponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccountList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListCustomerServiceAccountReponse) contextValidateAccountList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccountList); i++ {

		if m.AccountList[i] != nil {
			if err := m.AccountList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("account_list" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("account_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListCustomerServiceAccountReponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListCustomerServiceAccountReponse) UnmarshalBinary(b []byte) error {
	var res ListCustomerServiceAccountReponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}