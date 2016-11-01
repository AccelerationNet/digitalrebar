package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// AvailablehammersInput availablehammers input
// swagger:model availablehammers-input
type AvailablehammersInput struct {

	// klass
	Klass string `json:"klass,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// priority
	Priority int64 `json:"priority,omitempty"`
}

// Validate validates this availablehammers input
func (m *AvailablehammersInput) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
