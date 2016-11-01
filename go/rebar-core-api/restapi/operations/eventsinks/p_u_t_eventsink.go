package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTEventsinkHandlerFunc turns a function with the right signature into a p u t eventsink handler
type PUTEventsinkHandlerFunc func(PUTEventsinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTEventsinkHandlerFunc) Handle(params PUTEventsinkParams) middleware.Responder {
	return fn(params)
}

// PUTEventsinkHandler interface for that can handle valid p u t eventsink params
type PUTEventsinkHandler interface {
	Handle(PUTEventsinkParams) middleware.Responder
}

// NewPUTEventsink creates a new http.Handler for the p u t eventsink operation
func NewPUTEventsink(ctx *middleware.Context, handler PUTEventsinkHandler) *PUTEventsink {
	return &PUTEventsink{Context: ctx, Handler: handler}
}

/*PUTEventsink swagger:route PUT /eventsinks/{id} Eventsinks pUTEventsink

Update EventSink

*/
type PUTEventsink struct {
	Context *middleware.Context
	Handler PUTEventsinkHandler
}

func (o *PUTEventsink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTEventsinkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
