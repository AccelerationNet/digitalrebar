package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTDnsnamefiltersHandlerFunc turns a function with the right signature into a p u t dnsnamefilters handler
type PUTDnsnamefiltersHandlerFunc func(PUTDnsnamefiltersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTDnsnamefiltersHandlerFunc) Handle(params PUTDnsnamefiltersParams) middleware.Responder {
	return fn(params)
}

// PUTDnsnamefiltersHandler interface for that can handle valid p u t dnsnamefilters params
type PUTDnsnamefiltersHandler interface {
	Handle(PUTDnsnamefiltersParams) middleware.Responder
}

// NewPUTDnsnamefilters creates a new http.Handler for the p u t dnsnamefilters operation
func NewPUTDnsnamefilters(ctx *middleware.Context, handler PUTDnsnamefiltersHandler) *PUTDnsnamefilters {
	return &PUTDnsnamefilters{Context: ctx, Handler: handler}
}

/*PUTDnsnamefilters swagger:route PUT /dnsnamefilters/{id} Dnsnamefilters pUTDnsnamefilters

Update DNSNameFilters

*/
type PUTDnsnamefilters struct {
	Context *middleware.Context
	Handler PUTDnsnamefiltersHandler
}

func (o *PUTDnsnamefilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTDnsnamefiltersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
