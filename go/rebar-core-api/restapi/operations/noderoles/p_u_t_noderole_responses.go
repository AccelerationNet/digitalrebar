package noderoles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTNoderoleOK p u t noderole o k

swagger:response pUTNoderoleOK
*/
type PUTNoderoleOK struct {

	// In: body
	Payload *models.NoderoleOutput `json:"body,omitempty"`
}

// NewPUTNoderoleOK creates PUTNoderoleOK with default headers values
func NewPUTNoderoleOK() *PUTNoderoleOK {
	return &PUTNoderoleOK{}
}

// WithPayload adds the payload to the p u t noderole o k response
func (o *PUTNoderoleOK) WithPayload(payload *models.NoderoleOutput) *PUTNoderoleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t noderole o k response
func (o *PUTNoderoleOK) SetPayload(payload *models.NoderoleOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTNoderoleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
