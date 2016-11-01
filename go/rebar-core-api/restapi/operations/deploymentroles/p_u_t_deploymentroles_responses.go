package deploymentroles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTDeploymentrolesOK p u t deploymentroles o k

swagger:response pUTDeploymentrolesOK
*/
type PUTDeploymentrolesOK struct {

	// In: body
	Payload *models.DeploymentrolesOutput `json:"body,omitempty"`
}

// NewPUTDeploymentrolesOK creates PUTDeploymentrolesOK with default headers values
func NewPUTDeploymentrolesOK() *PUTDeploymentrolesOK {
	return &PUTDeploymentrolesOK{}
}

// WithPayload adds the payload to the p u t deploymentroles o k response
func (o *PUTDeploymentrolesOK) WithPayload(payload *models.DeploymentrolesOutput) *PUTDeploymentrolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t deploymentroles o k response
func (o *PUTDeploymentrolesOK) SetPayload(payload *models.DeploymentrolesOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTDeploymentrolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
