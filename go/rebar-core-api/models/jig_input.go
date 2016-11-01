package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// JigInput jig input
// swagger:model jig-input
type JigInput struct {

	// active
	Active bool `json:"active,omitempty"`

	// client name
	ClientName string `json:"client_name,omitempty"`

	// client role name
	ClientRoleName string `json:"client_role_name,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// server
	Server string `json:"server,omitempty"`
}

// Validate validates this jig input
func (m *JigInput) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
