package attribs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEAttribNoContent d e l e t e attrib no content

swagger:response dELETEAttribNoContent
*/
type DELETEAttribNoContent struct {
}

// NewDELETEAttribNoContent creates DELETEAttribNoContent with default headers values
func NewDELETEAttribNoContent() *DELETEAttribNoContent {
	return &DELETEAttribNoContent{}
}

// WriteResponse to the client
func (o *DELETEAttribNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
