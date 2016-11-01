package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*PUTRolesOK p u t roles o k

swagger:response pUTRolesOK
*/
type PUTRolesOK struct {

	// In: body
	Payload *models.RolesOutput `json:"body,omitempty"`
}

// NewPUTRolesOK creates PUTRolesOK with default headers values
func NewPUTRolesOK() *PUTRolesOK {
	return &PUTRolesOK{}
}

// WithPayload adds the payload to the p u t roles o k response
func (o *PUTRolesOK) WithPayload(payload *models.RolesOutput) *PUTRolesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p u t roles o k response
func (o *PUTRolesOK) SetPayload(payload *models.RolesOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PUTRolesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
