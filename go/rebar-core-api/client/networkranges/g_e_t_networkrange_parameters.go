package networkranges

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

// NewGETNetworkrangeParams creates a new GETNetworkrangeParams object
// with the default values initialized.
func NewGETNetworkrangeParams() *GETNetworkrangeParams {
	var ()
	return &GETNetworkrangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETNetworkrangeParamsWithTimeout creates a new GETNetworkrangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETNetworkrangeParamsWithTimeout(timeout time.Duration) *GETNetworkrangeParams {
	var ()
	return &GETNetworkrangeParams{

		timeout: timeout,
	}
}

// NewGETNetworkrangeParamsWithContext creates a new GETNetworkrangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETNetworkrangeParamsWithContext(ctx context.Context) *GETNetworkrangeParams {
	var ()
	return &GETNetworkrangeParams{

		Context: ctx,
	}
}

/*GETNetworkrangeParams contains all the parameters to send to the API endpoint
for the g e t networkrange operation typically these are written to a http.Request
*/
type GETNetworkrangeParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the g e t networkrange params
func (o *GETNetworkrangeParams) WithTimeout(timeout time.Duration) *GETNetworkrangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t networkrange params
func (o *GETNetworkrangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t networkrange params
func (o *GETNetworkrangeParams) WithContext(ctx context.Context) *GETNetworkrangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t networkrange params
func (o *GETNetworkrangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the g e t networkrange params
func (o *GETNetworkrangeParams) WithID(id string) *GETNetworkrangeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t networkrange params
func (o *GETNetworkrangeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETNetworkrangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
