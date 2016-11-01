package networkranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTNetworkrangeOK p u t networkrange o k

swagger:response pUTNetworkrangeOK
*/
type PUTNetworkrangeOK struct {

	// In: body
	Payload *models.NetworkrangeOutput `json:"body,omitempty"`
}

// NewPUTNetworkrangeOK creates PUTNetworkrangeOK with default headers values
func NewPUTNetworkrangeOK() *PUTNetworkrangeOK {
	return &PUTNetworkrangeOK{}
}

// WithPayload adds the payload to the p u t networkrange o k response
func (o *PUTNetworkrangeOK) WithPayload(payload *models.NetworkrangeOutput) *PUTNetworkrangeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t networkrange o k response
func (o *PUTNetworkrangeOK) SetPayload(payload *models.NetworkrangeOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTNetworkrangeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
