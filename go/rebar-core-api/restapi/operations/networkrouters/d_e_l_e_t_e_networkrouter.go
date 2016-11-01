package networkrouters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DELETENetworkrouterHandlerFunc turns a function with the right signature into a d e l e t e networkrouter handler
type DELETENetworkrouterHandlerFunc func(DELETENetworkrouterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETENetworkrouterHandlerFunc) Handle(params DELETENetworkrouterParams) middleware.Responder {
	return fn(params)
}

// DELETENetworkrouterHandler interface for that can handle valid d e l e t e networkrouter params
type DELETENetworkrouterHandler interface {
	Handle(DELETENetworkrouterParams) middleware.Responder
}

// NewDELETENetworkrouter creates a new http.Handler for the d e l e t e networkrouter operation
func NewDELETENetworkrouter(ctx *middleware.Context, handler DELETENetworkrouterHandler) *DELETENetworkrouter {
	return &DELETENetworkrouter{Context: ctx, Handler: handler}
}

/*DELETENetworkrouter swagger:route DELETE /networkrouters/{id} Networkrouters dELETENetworkrouter

Delete NetworkRouter

*/
type DELETENetworkrouter struct {
	Context *middleware.Context
	Handler DELETENetworkrouterHandler
}

func (o *DELETENetworkrouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETENetworkrouterParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
