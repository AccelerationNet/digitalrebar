package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// LISTTenantsHandlerFunc turns a function with the right signature into a l i s t tenants handler
type LISTTenantsHandlerFunc func(LISTTenantsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn LISTTenantsHandlerFunc) Handle(params LISTTenantsParams) middleware.Responder {
	return fn(params)
}

// LISTTenantsHandler interface for that can handle valid l i s t tenants params
type LISTTenantsHandler interface {
	Handle(LISTTenantsParams) middleware.Responder
}

// NewLISTTenants creates a new http.Handler for the l i s t tenants operation
func NewLISTTenants(ctx *middleware.Context, handler LISTTenantsHandler) *LISTTenants {
	return &LISTTenants{Context: ctx, Handler: handler}
}

/*LISTTenants swagger:route GET /tenants Tenants lISTTenants

List Tenants

*/
type LISTTenants struct {
	Context *middleware.Context
	Handler LISTTenantsHandler
}

func (o *LISTTenants) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewLISTTenantsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
