package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTProfileOK p u t profile o k

swagger:response pUTProfileOK
*/
type PUTProfileOK struct {

	// In: body
	Payload *models.ProfileOutput `json:"body,omitempty"`
}

// NewPUTProfileOK creates PUTProfileOK with default headers values
func NewPUTProfileOK() *PUTProfileOK {
	return &PUTProfileOK{}
}

// WithPayload adds the payload to the p u t profile o k response
func (o *PUTProfileOK) WithPayload(payload *models.ProfileOutput) *PUTProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t profile o k response
func (o *PUTProfileOK) SetPayload(payload *models.ProfileOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
