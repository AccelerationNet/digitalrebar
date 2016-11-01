package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewPATCHNodeParams creates a new PATCHNodeParams object
// with the default values initialized.
func NewPATCHNodeParams() PATCHNodeParams {
	var ()
	return PATCHNodeParams{}
}

// PATCHNodeParams contains all the bound params for the p a t c h node operation
// typically these are obtained from a http.Request
//
// swagger:parameters PATCH-node
type PATCHNodeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PATCHNodeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
