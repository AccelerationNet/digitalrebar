package eventsinks

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

// NewGETEventsinkParams creates a new GETEventsinkParams object
// with the default values initialized.
func NewGETEventsinkParams() *GETEventsinkParams {
	var ()
	return &GETEventsinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGETEventsinkParamsWithTimeout creates a new GETEventsinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETEventsinkParamsWithTimeout(timeout time.Duration) *GETEventsinkParams {
	var ()
	return &GETEventsinkParams{

		timeout: timeout,
	}
}

// NewGETEventsinkParamsWithContext creates a new GETEventsinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETEventsinkParamsWithContext(ctx context.Context) *GETEventsinkParams {
	var ()
	return &GETEventsinkParams{

		Context: ctx,
	}
}

/*GETEventsinkParams contains all the parameters to send to the API endpoint
for the g e t eventsink operation typically these are written to a http.Request
*/
type GETEventsinkParams struct {

	/*ID*/
	ID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the g e t eventsink params
func (o *GETEventsinkParams) WithTimeout(timeout time.Duration) *GETEventsinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the g e t eventsink params
func (o *GETEventsinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the g e t eventsink params
func (o *GETEventsinkParams) WithContext(ctx context.Context) *GETEventsinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the g e t eventsink params
func (o *GETEventsinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the g e t eventsink params
func (o *GETEventsinkParams) WithID(id string) *GETEventsinkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the g e t eventsink params
func (o *GETEventsinkParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GETEventsinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
