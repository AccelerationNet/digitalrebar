package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTDnsnamefiltersOK p u t dnsnamefilters o k

swagger:response pUTDnsnamefiltersOK
*/
type PUTDnsnamefiltersOK struct {

	// In: body
	Payload *models.DnsnamefiltersOutput `json:"body,omitempty"`
}

// NewPUTDnsnamefiltersOK creates PUTDnsnamefiltersOK with default headers values
func NewPUTDnsnamefiltersOK() *PUTDnsnamefiltersOK {
	return &PUTDnsnamefiltersOK{}
}

// WithPayload adds the payload to the p u t dnsnamefilters o k response
func (o *PUTDnsnamefiltersOK) WithPayload(payload *models.DnsnamefiltersOutput) *PUTDnsnamefiltersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t dnsnamefilters o k response
func (o *PUTDnsnamefiltersOK) SetPayload(payload *models.DnsnamefiltersOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTDnsnamefiltersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
