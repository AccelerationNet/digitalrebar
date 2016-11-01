package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETENetworkNoContent d e l e t e network no content

swagger:response dELETENetworkNoContent
*/
type DELETENetworkNoContent struct {
}

// NewDELETENetworkNoContent creates DELETENetworkNoContent with default headers values
func NewDELETENetworkNoContent() *DELETENetworkNoContent {
	return &DELETENetworkNoContent{}
}

// WriteResponse to the client
func (o *DELETENetworkNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
