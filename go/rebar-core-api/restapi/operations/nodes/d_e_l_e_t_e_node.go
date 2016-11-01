package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DELETENodeHandlerFunc turns a function with the right signature into a d e l e t e node handler
type DELETENodeHandlerFunc func(DELETENodeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DELETENodeHandlerFunc) Handle(params DELETENodeParams) middleware.Responder {
	return fn(params)
}

// DELETENodeHandler interface for that can handle valid d e l e t e node params
type DELETENodeHandler interface {
	Handle(DELETENodeParams) middleware.Responder
}

// NewDELETENode creates a new http.Handler for the d e l e t e node operation
func NewDELETENode(ctx *middleware.Context, handler DELETENodeHandler) *DELETENode {
	return &DELETENode{Context: ctx, Handler: handler}
}

/*DELETENode swagger:route DELETE /nodes/{id} Nodes dELETENode

Delete Node

*/
type DELETENode struct {
	Context *middleware.Context
	Handler DELETENodeHandler
}

func (o *DELETENode) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDELETENodeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
