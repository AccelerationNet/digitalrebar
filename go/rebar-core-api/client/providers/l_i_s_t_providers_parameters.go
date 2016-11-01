package providers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLISTProvidersParams creates a new LISTProvidersParams object
// with the default values initialized.
func NewLISTProvidersParams() *LISTProvidersParams {

	return &LISTProvidersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLISTProvidersParamsWithTimeout creates a new LISTProvidersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLISTProvidersParamsWithTimeout(timeout time.Duration) *LISTProvidersParams {

	return &LISTProvidersParams{

		timeout: timeout,
	}
}

// NewLISTProvidersParamsWithContext creates a new LISTProvidersParams object
// with the default values initialized, and the ability to set a context for a request
func NewLISTProvidersParamsWithContext(ctx context.Context) *LISTProvidersParams {

	return &LISTProvidersParams{

		Context: ctx,
	}
}

/*LISTProvidersParams contains all the parameters to send to the API endpoint
for the l i s t providers operation typically these are written to a http.Request
*/
type LISTProvidersParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the l i s t providers params
func (o *LISTProvidersParams) WithTimeout(timeout time.Duration) *LISTProvidersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the l i s t providers params
func (o *LISTProvidersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the l i s t providers params
func (o *LISTProvidersParams) WithContext(ctx context.Context) *LISTProvidersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the l i s t providers params
func (o *LISTProvidersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *LISTProvidersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
