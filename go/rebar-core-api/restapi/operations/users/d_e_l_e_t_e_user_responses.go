package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEUserNoContent d e l e t e user no content

swagger:response dELETEUserNoContent
*/
type DELETEUserNoContent struct {
}

// NewDELETEUserNoContent creates DELETEUserNoContent with default headers values
func NewDELETEUserNoContent() *DELETEUserNoContent {
	return &DELETEUserNoContent{}
}

// WriteResponse to the client
func (o *DELETEUserNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
