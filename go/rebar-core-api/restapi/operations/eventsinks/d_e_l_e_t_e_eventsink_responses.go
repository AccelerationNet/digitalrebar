package eventsinks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEEventsinkNoContent d e l e t e eventsink no content

swagger:response dELETEEventsinkNoContent
*/
type DELETEEventsinkNoContent struct {
}

// NewDELETEEventsinkNoContent creates DELETEEventsinkNoContent with default headers values
func NewDELETEEventsinkNoContent() *DELETEEventsinkNoContent {
	return &DELETEEventsinkNoContent{}
}

// WriteResponse to the client
func (o *DELETEEventsinkNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
