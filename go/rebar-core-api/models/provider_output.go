package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// ProviderOutput provider output
// swagger:model provider-output
type ProviderOutput struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	ProviderInput
}

// Validate validates this provider output
func (m *ProviderOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.ProviderInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
