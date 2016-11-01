package dnsnameentries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// PUTDnsnameentriesReader is a Reader for the PUTDnsnameentries structure.
type PUTDnsnameentriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PUTDnsnameentriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPUTDnsnameentriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPUTDnsnameentriesOK creates a PUTDnsnameentriesOK with default headers values
func NewPUTDnsnameentriesOK() *PUTDnsnameentriesOK {
	return &PUTDnsnameentriesOK{}
}

/*PUTDnsnameentriesOK handles this case with default header values.

PUTDnsnameentriesOK p u t dnsnameentries o k
*/
type PUTDnsnameentriesOK struct {
	Payload *models.DnsnameentriesOutput
}

func (o *PUTDnsnameentriesOK) Error() string {
	return fmt.Sprintf("[PUT /dnsnameentries/{id}][%d] pUTDnsnameentriesOK  %+v", 200, o.Payload)
}

func (o *PUTDnsnameentriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DnsnameentriesOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
