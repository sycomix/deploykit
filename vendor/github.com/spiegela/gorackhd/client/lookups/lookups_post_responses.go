package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// LookupsPostReader is a Reader for the LookupsPost structure.
type LookupsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LookupsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewLookupsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewLookupsPostDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLookupsPostCreated creates a LookupsPostCreated with default headers values
func NewLookupsPostCreated() *LookupsPostCreated {
	return &LookupsPostCreated{}
}

/*LookupsPostCreated handles this case with default header values.

Successfully created new lookup
*/
type LookupsPostCreated struct {
	Payload *models.Lookups20LookupBase
}

func (o *LookupsPostCreated) Error() string {
	return fmt.Sprintf("[POST /lookups][%d] lookupsPostCreated  %+v", 201, o.Payload)
}

func (o *LookupsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Lookups20LookupBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLookupsPostDefault creates a LookupsPostDefault with default headers values
func NewLookupsPostDefault(code int) *LookupsPostDefault {
	return &LookupsPostDefault{
		_statusCode: code,
	}
}

/*LookupsPostDefault handles this case with default header values.

Unexpected error
*/
type LookupsPostDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the lookups post default response
func (o *LookupsPostDefault) Code() int {
	return o._statusCode
}

func (o *LookupsPostDefault) Error() string {
	return fmt.Sprintf("[POST /lookups][%d] lookupsPost default  %+v", o._statusCode, o.Payload)
}

func (o *LookupsPostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
