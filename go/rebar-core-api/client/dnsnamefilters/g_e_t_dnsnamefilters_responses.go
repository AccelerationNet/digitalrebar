package dnsnamefilters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// GETDnsnamefiltersReader is a Reader for the GETDnsnamefilters structure.
type GETDnsnamefiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETDnsnamefiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETDnsnamefiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETDnsnamefiltersOK creates a GETDnsnamefiltersOK with default headers values
func NewGETDnsnamefiltersOK() *GETDnsnamefiltersOK {
	return &GETDnsnamefiltersOK{}
}

/*GETDnsnamefiltersOK handles this case with default header values.

GETDnsnamefiltersOK g e t dnsnamefilters o k
*/
type GETDnsnamefiltersOK struct {
	Payload *models.DnsnamefiltersOutput
}

func (o *GETDnsnamefiltersOK) Error() string {
	return fmt.Sprintf("[GET /dnsnamefilters/{id}][%d] gETDnsnamefiltersOK  %+v", 200, o.Payload)
}

func (o *GETDnsnamefiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DnsnamefiltersOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
