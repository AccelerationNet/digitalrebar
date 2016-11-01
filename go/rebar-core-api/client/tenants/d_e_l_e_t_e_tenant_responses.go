package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DELETETenantReader is a Reader for the DELETETenant structure.
type DELETETenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DELETETenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDELETETenantNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDELETETenantNoContent creates a DELETETenantNoContent with default headers values
func NewDELETETenantNoContent() *DELETETenantNoContent {
	return &DELETETenantNoContent{}
}

/*DELETETenantNoContent handles this case with default header values.

DELETETenantNoContent d e l e t e tenant no content
*/
type DELETETenantNoContent struct {
}

func (o *DELETETenantNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenants/{id}][%d] dELETETenantNoContent ", 204)
}

func (o *DELETETenantNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
