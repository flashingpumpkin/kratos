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

// Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity Identity identity
//
// swagger:model Identity
type Identity struct {

	// id
	// Required: true
	// Format: uuid4
	ID *UUID `json:"id"`

	// RecoveryAddresses contains all the addresses that can be used to recover an identity.
	RecoveryAddresses []*RecoveryAddress `json:"recovery_addresses,omitempty"`

	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits.
	// Required: true
	SchemaID *string `json:"schema_id"`

	// SchemaURL is the URL of the endpoint where the identity's traits schema can be fetched from.
	//
	// format: url
	// Required: true
	SchemaURL *string `json:"schema_url"`

	// traits
	// Required: true
	Traits Traits `json:"traits"`

	// VerifiableAddresses contains all the addresses that can be verified by the user.
	VerifiableAddresses []*VerifiableAddress `json:"verifiable_addresses,omitempty"`
}

// Validate validates this identity
func (m *Identity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemaID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemaURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVerifiableAddresses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identity) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if m.ID != nil {
		if err := m.ID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *Identity) validateRecoveryAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoveryAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.RecoveryAddresses); i++ {
		if swag.IsZero(m.RecoveryAddresses[i]) { // not required
			continue
		}

		if m.RecoveryAddresses[i] != nil {
			if err := m.RecoveryAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recovery_addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Identity) validateSchemaID(formats strfmt.Registry) error {

	if err := validate.Required("schema_id", "body", m.SchemaID); err != nil {
		return err
	}

	return nil
}

func (m *Identity) validateSchemaURL(formats strfmt.Registry) error {

	if err := validate.Required("schema_url", "body", m.SchemaURL); err != nil {
		return err
	}

	return nil
}

func (m *Identity) validateTraits(formats strfmt.Registry) error {

	if m.Traits == nil {
		return errors.Required("traits", "body", nil)
	}

	return nil
}

func (m *Identity) validateVerifiableAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.VerifiableAddresses) { // not required
		return nil
	}

	for i := 0; i < len(m.VerifiableAddresses); i++ {
		if swag.IsZero(m.VerifiableAddresses[i]) { // not required
			continue
		}

		if m.VerifiableAddresses[i] != nil {
			if err := m.VerifiableAddresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verifiable_addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this identity based on the context it is used
func (m *Identity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecoveryAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVerifiableAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Identity) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if m.ID != nil {
		if err := m.ID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id")
			}
			return err
		}
	}

	return nil
}

func (m *Identity) contextValidateRecoveryAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecoveryAddresses); i++ {

		if m.RecoveryAddresses[i] != nil {
			if err := m.RecoveryAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recovery_addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Identity) contextValidateVerifiableAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VerifiableAddresses); i++ {

		if m.VerifiableAddresses[i] != nil {
			if err := m.VerifiableAddresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("verifiable_addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Identity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Identity) UnmarshalBinary(b []byte) error {
	var res Identity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}