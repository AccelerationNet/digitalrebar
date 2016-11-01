package barclamps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETBarclampOK g e t barclamp o k

swagger:response gETBarclampOK
*/
type GETBarclampOK struct {

	// In: body
	Payload *models.BarclampOutput `json:"body,omitempty"`
}

// NewGETBarclampOK creates GETBarclampOK with default headers values
func NewGETBarclampOK() *GETBarclampOK {
	return &GETBarclampOK{}
}

// WithPayload adds the payload to the g e t barclamp o k response
func (o *GETBarclampOK) WithPayload(payload *models.BarclampOutput) *GETBarclampOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t barclamp o k response
func (o *GETBarclampOK) SetPayload(payload *models.BarclampOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETBarclampOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
