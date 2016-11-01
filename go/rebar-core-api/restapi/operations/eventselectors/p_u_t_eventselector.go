package eventselectors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTEventselectorHandlerFunc turns a function with the right signature into a p u t eventselector handler
type PUTEventselectorHandlerFunc func(PUTEventselectorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTEventselectorHandlerFunc) Handle(params PUTEventselectorParams) middleware.Responder {
	return fn(params)
}

// PUTEventselectorHandler interface for that can handle valid p u t eventselector params
type PUTEventselectorHandler interface {
	Handle(PUTEventselectorParams) middleware.Responder
}

// NewPUTEventselector creates a new http.Handler for the p u t eventselector operation
func NewPUTEventselector(ctx *middleware.Context, handler PUTEventselectorHandler) *PUTEventselector {
	return &PUTEventselector{Context: ctx, Handler: handler}
}

/*PUTEventselector swagger:route PUT /eventselectors/{id} Eventselectors pUTEventselector

Update EventSelector

*/
type PUTEventselector struct {
	Context *middleware.Context
	Handler PUTEventselectorHandler
}

func (o *PUTEventselector) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTEventselectorParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
