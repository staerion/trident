// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ScopeReference Scope of the entity. Set to "cluster" for cluster owned objects and to "svm" for SVM owned objects.
//
// swagger:model scope_reference
type ScopeReference string

const (

	// ScopeReferenceCluster captures enum value "cluster"
	ScopeReferenceCluster ScopeReference = "cluster"

	// ScopeReferenceSvm captures enum value "svm"
	ScopeReferenceSvm ScopeReference = "svm"
)

// for schema
var scopeReferenceEnum []interface{}

func init() {
	var res []ScopeReference
	if err := json.Unmarshal([]byte(`["cluster","svm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scopeReferenceEnum = append(scopeReferenceEnum, v)
	}
}

func (m ScopeReference) validateScopeReferenceEnum(path, location string, value ScopeReference) error {
	if err := validate.EnumCase(path, location, value, scopeReferenceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this scope reference
func (m ScopeReference) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateScopeReferenceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this scope reference based on the context it is used
func (m ScopeReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", ScopeReference(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}