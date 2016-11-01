package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// POSTDnsnamefiltersHandlerFunc turns a function with the right signature into a p o s t dnsnamefilters handler
type POSTDnsnamefiltersHandlerFunc func(POSTDnsnamefiltersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn POSTDnsnamefiltersHandlerFunc) Handle(params POSTDnsnamefiltersParams) middleware.Responder {
	return fn(params)
}

// POSTDnsnamefiltersHandler interface for that can handle valid p o s t dnsnamefilters params
type POSTDnsnamefiltersHandler interface {
	Handle(POSTDnsnamefiltersParams) middleware.Responder
}

// NewPOSTDnsnamefilters creates a new http.Handler for the p o s t dnsnamefilters operation
func NewPOSTDnsnamefilters(ctx *middleware.Context, handler POSTDnsnamefiltersHandler) *POSTDnsnamefilters {
	return &POSTDnsnamefilters{Context: ctx, Handler: handler}
}

/*POSTDnsnamefilters swagger:route POST /dnsnamefilters Dnsnamefilters pOSTDnsnamefilters

Create DNSNameFilters

*/
type POSTDnsnamefilters struct {
	Context *middleware.Context
	Handler POSTDnsnamefiltersHandler
}

func (o *POSTDnsnamefilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPOSTDnsnamefiltersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
