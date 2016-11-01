package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LISTDnsnamefiltersHandlerFunc turns a function with the right signature into a l i s t dnsnamefilters handler
type LISTDnsnamefiltersHandlerFunc func(LISTDnsnamefiltersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LISTDnsnamefiltersHandlerFunc) Handle(params LISTDnsnamefiltersParams) middleware.Responder {
	return fn(params)
}

// LISTDnsnamefiltersHandler interface for that can handle valid l i s t dnsnamefilters params
type LISTDnsnamefiltersHandler interface {
	Handle(LISTDnsnamefiltersParams) middleware.Responder
}

// NewLISTDnsnamefilters creates a new http.Handler for the l i s t dnsnamefilters operation
func NewLISTDnsnamefilters(ctx *middleware.Context, handler LISTDnsnamefiltersHandler) *LISTDnsnamefilters {
	return &LISTDnsnamefilters{Context: ctx, Handler: handler}
}

/*LISTDnsnamefilters swagger:route GET /dnsnamefilters Dnsnamefilters lISTDnsnamefilters

List Dnsnamefilters

*/
type LISTDnsnamefilters struct {
	Context *middleware.Context
	Handler LISTDnsnamefiltersHandler
}

func (o *LISTDnsnamefilters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewLISTDnsnamefiltersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
