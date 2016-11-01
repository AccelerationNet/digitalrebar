package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalrebar/digitalrebar/go/rebar-core-api/models"
)

// LISTProfilesReader is a Reader for the LISTProfiles structure.
type LISTProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LISTProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLISTProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLISTProfilesOK creates a LISTProfilesOK with default headers values
func NewLISTProfilesOK() *LISTProfilesOK {
	return &LISTProfilesOK{}
}

/*LISTProfilesOK handles this case with default header values.

LISTProfilesOK l i s t profiles o k
*/
type LISTProfilesOK struct {
	Payload []*models.ProfileOutput
}

func (o *LISTProfilesOK) Error() string {
	return fmt.Sprintf("[GET /profiles][%d] lISTProfilesOK  %+v", 200, o.Payload)
}

func (o *LISTProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
