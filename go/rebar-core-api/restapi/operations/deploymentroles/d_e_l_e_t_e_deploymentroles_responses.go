package deploymentroles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*DELETEDeploymentrolesNoContent d e l e t e deploymentroles no content

swagger:response dELETEDeploymentrolesNoContent
*/
type DELETEDeploymentrolesNoContent struct {
}

// NewDELETEDeploymentrolesNoContent creates DELETEDeploymentrolesNoContent with default headers values
func NewDELETEDeploymentrolesNoContent() *DELETEDeploymentrolesNoContent {
	return &DELETEDeploymentrolesNoContent{}
}

// WriteResponse to the client
func (o *DELETEDeploymentrolesNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}
