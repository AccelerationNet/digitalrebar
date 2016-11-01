package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

/*LISTEventsinksOK l i s t eventsinks o k

swagger:response lISTEventsinksOK
*/
type LISTEventsinksOK struct {

	// In: body
	Payload []*models.EventsinkOutput `json:"body,omitempty"`
}

// NewLISTEventsinksOK creates LISTEventsinksOK with default headers values
func NewLISTEventsinksOK() *LISTEventsinksOK {
	return &LISTEventsinksOK{}
}

// WithPayload adds the payload to the l i s t eventsinks o k response
func (o *LISTEventsinksOK) WithPayload(payload []*models.EventsinkOutput) *LISTEventsinksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the l i s t eventsinks o k response
func (o *LISTEventsinksOK) SetPayload(payload []*models.EventsinkOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LISTEventsinksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
