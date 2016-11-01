package availablehammers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PUTAvailablehammersHandlerFunc turns a function with the right signature into a p u t availablehammers handler
type PUTAvailablehammersHandlerFunc func(PUTAvailablehammersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PUTAvailablehammersHandlerFunc) Handle(params PUTAvailablehammersParams) middleware.Responder {
	return fn(params)
}

// PUTAvailablehammersHandler interface for that can handle valid p u t availablehammers params
type PUTAvailablehammersHandler interface {
	Handle(PUTAvailablehammersParams) middleware.Responder
}

// NewPUTAvailablehammers creates a new http.Handler for the p u t availablehammers operation
func NewPUTAvailablehammers(ctx *middleware.Context, handler PUTAvailablehammersHandler) *PUTAvailablehammers {
	return &PUTAvailablehammers{Context: ctx, Handler: handler}
}

/*PUTAvailablehammers swagger:route PUT /availablehammers/{id} Availablehammers pUTAvailablehammers

Update AvailableHammers

*/
type PUTAvailablehammers struct {
	Context *middleware.Context
	Handler PUTAvailablehammersHandler
}

func (o *PUTAvailablehammers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPUTAvailablehammersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
