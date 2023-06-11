// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetServiceStateReponse get service state reponse
//
// swagger:model GetServiceStateReponse
type GetServiceStateReponse struct {

	// errcode
	Errcode int32 `json:"errcode,omitempty"`

	// errmsg
	Errmsg string `json:"errmsg,omitempty"`

	// service state
	ServiceState int32 `json:"service_state,omitempty"`

	// servicer userid
	ServicerUserid string `json:"servicer_userid,omitempty"`
}

// Validate validates this get service state reponse
func (m *GetServiceStateReponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get service state reponse based on context it is used
func (m *GetServiceStateReponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetServiceStateReponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetServiceStateReponse) UnmarshalBinary(b []byte) error {
	var res GetServiceStateReponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}