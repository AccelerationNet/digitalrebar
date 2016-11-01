package networkranges

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// GETNetworkrangeReader is a Reader for the GETNetworkrange structure.
type GETNetworkrangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETNetworkrangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETNetworkrangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETNetworkrangeOK creates a GETNetworkrangeOK with default headers values
func NewGETNetworkrangeOK() *GETNetworkrangeOK {
	return &GETNetworkrangeOK{}
}

/*GETNetworkrangeOK handles this case with default header values.

GETNetworkrangeOK g e t networkrange o k
*/
type GETNetworkrangeOK struct {
	Payload *models.NetworkrangeOutput
}

func (o *GETNetworkrangeOK) Error() string {
	return fmt.Sprintf("[GET /networkranges/{id}][%d] gETNetworkrangeOK  %+v", 200, o.Payload)
}

func (o *GETNetworkrangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkrangeOutput)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
