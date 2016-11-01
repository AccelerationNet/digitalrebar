package availablehammers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETAvailablehammersOK g e t availablehammers o k

swagger:response gETAvailablehammersOK
*/
type GETAvailablehammersOK struct {

	// In: body
	Payload *models.AvailablehammersOutput `json:"body,omitempty"`
}

// NewGETAvailablehammersOK creates GETAvailablehammersOK with default headers values
func NewGETAvailablehammersOK() *GETAvailablehammersOK {
	return &GETAvailablehammersOK{}
}

// WithPayload adds the payload to the g e t availablehammers o k response
func (o *GETAvailablehammersOK) WithPayload(payload *models.AvailablehammersOutput) *GETAvailablehammersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t availablehammers o k response
func (o *GETAvailablehammersOK) SetPayload(payload *models.AvailablehammersOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETAvailablehammersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
