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

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// NewPOSTEventsinkParams creates a new POSTEventsinkParams object
// with the default values initialized.
func NewPOSTEventsinkParams() *POSTEventsinkParams {
	var ()
	return &POSTEventsinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPOSTEventsinkParamsWithTimeout creates a new POSTEventsinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPOSTEventsinkParamsWithTimeout(timeout time.Duration) *POSTEventsinkParams {
	var ()
	return &POSTEventsinkParams{

		timeout: timeout,
	}
}

// NewPOSTEventsinkParamsWithContext creates a new POSTEventsinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPOSTEventsinkParamsWithContext(ctx context.Context) *POSTEventsinkParams {
	var ()
	return &POSTEventsinkParams{

		Context: ctx,
	}
}

/*POSTEventsinkParams contains all the parameters to send to the API endpoint
for the p o s t eventsink operation typically these are written to a http.Request
*/
type POSTEventsinkParams struct {

	/*Body*/
	Body *models.EventsinkInput

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the p o s t eventsink params
func (o *POSTEventsinkParams) WithTimeout(timeout time.Duration) *POSTEventsinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the p o s t eventsink params
func (o *POSTEventsinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the p o s t eventsink params
func (o *POSTEventsinkParams) WithContext(ctx context.Context) *POSTEventsinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the p o s t eventsink params
func (o *POSTEventsinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the p o s t eventsink params
func (o *POSTEventsinkParams) WithBody(body *models.EventsinkInput) *POSTEventsinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the p o s t eventsink params
func (o *POSTEventsinkParams) SetBody(body *models.EventsinkInput) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *POSTEventsinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.EventsinkInput)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
