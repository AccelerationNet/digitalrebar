package attribs

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

// NewLISTAttribsParams creates a new LISTAttribsParams object
// with the default values initialized.
func NewLISTAttribsParams() *LISTAttribsParams {

	return &LISTAttribsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLISTAttribsParamsWithTimeout creates a new LISTAttribsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLISTAttribsParamsWithTimeout(timeout time.Duration) *LISTAttribsParams {

	return &LISTAttribsParams{

		timeout: timeout,
	}
}

// NewLISTAttribsParamsWithContext creates a new LISTAttribsParams object
// with the default values initialized, and the ability to set a context for a request
func NewLISTAttribsParamsWithContext(ctx context.Context) *LISTAttribsParams {

	return &LISTAttribsParams{

		Context: ctx,
	}
}

/*LISTAttribsParams contains all the parameters to send to the API endpoint
for the l i s t attribs operation typically these are written to a http.Request
*/
type LISTAttribsParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the l i s t attribs params
func (o *LISTAttribsParams) WithTimeout(timeout time.Duration) *LISTAttribsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the l i s t attribs params
func (o *LISTAttribsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the l i s t attribs params
func (o *LISTAttribsParams) WithContext(ctx context.Context) *LISTAttribsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the l i s t attribs params
func (o *LISTAttribsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *LISTAttribsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
