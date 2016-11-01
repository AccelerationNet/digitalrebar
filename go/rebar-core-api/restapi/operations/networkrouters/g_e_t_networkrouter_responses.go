package networkrouters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETNetworkrouterOK g e t networkrouter o k

swagger:response gETNetworkrouterOK
*/
type GETNetworkrouterOK struct {

	// In: body
	Payload *models.NetworkrouterOutput `json:"body,omitempty"`
}

// NewGETNetworkrouterOK creates GETNetworkrouterOK with default headers values
func NewGETNetworkrouterOK() *GETNetworkrouterOK {
	return &GETNetworkrouterOK{}
}

// WithPayload adds the payload to the g e t networkrouter o k response
func (o *GETNetworkrouterOK) WithPayload(payload *models.NetworkrouterOutput) *GETNetworkrouterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t networkrouter o k response
func (o *GETNetworkrouterOK) SetPayload(payload *models.NetworkrouterOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETNetworkrouterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
