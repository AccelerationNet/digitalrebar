package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*GETDeploymentsOK g e t deployments o k

swagger:response gETDeploymentsOK
*/
type GETDeploymentsOK struct {

	// In: body
	Payload *models.DeploymentsOutput `json:"body,omitempty"`
}

// NewGETDeploymentsOK creates GETDeploymentsOK with default headers values
func NewGETDeploymentsOK() *GETDeploymentsOK {
	return &GETDeploymentsOK{}
}

// WithPayload adds the payload to the g e t deployments o k response
func (o *GETDeploymentsOK) WithPayload(payload *models.DeploymentsOutput) *GETDeploymentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the g e t deployments o k response
func (o *GETDeploymentsOK) SetPayload(payload *models.DeploymentsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GETDeploymentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
