package noderoles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETENoderoleNoContent d e l e t e noderole no content

swagger:response dELETENoderoleNoContent
*/
type DELETENoderoleNoContent struct {
}

// NewDELETENoderoleNoContent creates DELETENoderoleNoContent with default headers values
func NewDELETENoderoleNoContent() *DELETENoderoleNoContent {
	return &DELETENoderoleNoContent{}
}

// WriteResponse to the client
func (o *DELETENoderoleNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
