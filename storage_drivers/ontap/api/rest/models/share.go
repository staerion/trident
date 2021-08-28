// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Share share
//
// swagger:model share
type Share struct {

	// Displays the file or directory effective permission for the mentioned user, only for files and directories contained where the
	// specified path is relative to the root of the specified share. If this parameter is not specified, the SVM root volume is
	// taken as the default. If this parameter is specified, the effective share permission of the user is also displayed.
	// Wildcard query characters are not supported.
	//
	Name string `json:"name,omitempty"`

	// Displays the CIFS share path.
	Path string `json:"path,omitempty"`
}

// Validate validates this share
func (m *Share) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this share based on context it is used
func (m *Share) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Share) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Share) UnmarshalBinary(b []byte) error {
	var res Share
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}