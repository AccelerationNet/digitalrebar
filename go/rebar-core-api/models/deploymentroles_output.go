package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// DeploymentrolesOutput deploymentroles output
// swagger:model deploymentroles-output
type DeploymentrolesOutput struct {

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	DeploymentrolesInput
}

// Validate validates this deploymentroles output
func (m *DeploymentrolesOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.DeploymentrolesInput.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
