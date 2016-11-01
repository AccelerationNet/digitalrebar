package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GETRolesHandlerFunc turns a function with the right signature into a g e t roles handler
type GETRolesHandlerFunc func(GETRolesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GETRolesHandlerFunc) Handle(params GETRolesParams) middleware.Responder {
	return fn(params)
}

// GETRolesHandler interface for that can handle valid g e t roles params
type GETRolesHandler interface {
	Handle(GETRolesParams) middleware.Responder
}

// NewGETRoles creates a new http.Handler for the g e t roles operation
func NewGETRoles(ctx *middleware.Context, handler GETRolesHandler) *GETRoles {
	return &GETRoles{Context: ctx, Handler: handler}
}

/*GETRoles swagger:route GET /roles/{id} Roles gETRoles

Get Roles

*/
type GETRoles struct {
	Context *middleware.Context
	Handler GETRolesHandler
}

func (o *GETRoles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGETRolesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
