// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// PlexGetReader is a Reader for the PlexGet structure.
type PlexGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlexGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlexGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlexGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlexGetOK creates a PlexGetOK with default headers values
func NewPlexGetOK() *PlexGetOK {
	return &PlexGetOK{}
}

/* PlexGetOK describes a response with status code 200, with default header values.

OK
*/
type PlexGetOK struct {
	Payload *models.Plex
}

func (o *PlexGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes/{name}][%d] plexGetOK  %+v", 200, o.Payload)
}
func (o *PlexGetOK) GetPayload() *models.Plex {
	return o.Payload
}

func (o *PlexGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Plex)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlexGetDefault creates a PlexGetDefault with default headers values
func NewPlexGetDefault(code int) *PlexGetDefault {
	return &PlexGetDefault{
		_statusCode: code,
	}
}

/* PlexGetDefault describes a response with status code -1, with default header values.

Error
*/
type PlexGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the plex get default response
func (o *PlexGetDefault) Code() int {
	return o._statusCode
}

func (o *PlexGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{aggregate.uuid}/plexes/{name}][%d] plex_get default  %+v", o._statusCode, o.Payload)
}
func (o *PlexGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PlexGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}