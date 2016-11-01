package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// NewPUTNetworkParams creates a new PUTNetworkParams object
// with the default values initialized.
func NewPUTNetworkParams() *PUTNetworkParams {
	var ()
	return &PUTNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPUTNetworkParamsWithTimeout creates a new PUTNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPUTNetworkParamsWithTimeout(timeout time.Duration) *PUTNetworkParams {
	var ()
	return &PUTNetworkParams{

		timeout: timeout,
	}
}

// NewPUTNetworkParamsWithContext creates a new PUTNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPUTNetworkParamsWithContext(ctx context.Context) *PUTNetworkParams {
	var ()
	return &PUTNetworkParams{

		Context: ctx,
	}
}

/*PUTNetworkParams contains all the parameters to send to the API endpoint
for the p u t network operation typically these are written to a http.Request
*/
type PUTNetworkParams struct {

	/*Body*/
	Body *models.NetworkInput
	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p u t network params
func (o *PUTNetworkParams) WithTimeout(timeout time.Duration) *PUTNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p u t network params
func (o *PUTNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p u t network params
func (o *PUTNetworkParams) WithContext(ctx context.Context) *PUTNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p u t network params
func (o *PUTNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p u t network params
func (o *PUTNetworkParams) WithBody(body *models.NetworkInput) *PUTNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p u t network params
func (o *PUTNetworkParams) SetBody(body *models.NetworkInput) {
	o.Body = body
}

// WithID adds the id to the p u t network params
func (o *PUTNetworkParams) WithID(id string) *PUTNetworkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the p u t network params
func (o *PUTNetworkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PUTNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.NetworkInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
