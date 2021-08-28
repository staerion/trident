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

// AggregateGetReader is a Reader for the AggregateGet structure.
type AggregateGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateGetOK creates a AggregateGetOK with default headers values
func NewAggregateGetOK() *AggregateGetOK {
	return &AggregateGetOK{}
}

/* AggregateGetOK describes a response with status code 200, with default header values.

OK
*/
type AggregateGetOK struct {
	Payload *models.Aggregate
}

func (o *AggregateGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregateGetOK  %+v", 200, o.Payload)
}
func (o *AggregateGetOK) GetPayload() *models.Aggregate {
	return o.Payload
}

func (o *AggregateGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Aggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateGetDefault creates a AggregateGetDefault with default headers values
func NewAggregateGetDefault(code int) *AggregateGetDefault {
	return &AggregateGetDefault{
		_statusCode: code,
	}
}

/* AggregateGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 787092 | The target field cannot be specified for this operation. |
| 7209049 | Cannot perform the operation because the aggregate is currently expanding. |
| 8586225 | Unexpected error encountered when retrieving metrics and statistics for this aggregate. |
| 19726382 | Another provisioning operation is in progress on this cluster. Wait a few minutes, and try the operation again. |
| 19726390 | Unable to provide a recommmendation to expand the aggregate. |
| 19726391 | Too many unassigned disks visible to the node that owns this aggregate. |
| 19726392 | Layout of this aggregate is not a supported configuration. |
| 19726393 | Failed to expand the aggregate. Aggregate expansion is not supported on this system. |
| 19726394 | Automatic aggregate expansion is not supported on systems with multiple data aggregates. |
| 19726395 | Automatic aggregate expansion is not supported when MetroCluster is not configured |
| 19726396 | Automatic aggregate expansion is not supported when the DR group is not in a normal state |
| 19726397 | Aggregates must contain disks with identical disk-types and disk-sizes. |
| 19726402 | Internal error. Unable to determine the MetroCluster configuration state. |
| 19726538 | Cannot perform the operation because the aggregate is not in a healthy state. |

*/
type AggregateGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the aggregate get default response
func (o *AggregateGetDefault) Code() int {
	return o._statusCode
}

func (o *AggregateGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/aggregates/{uuid}][%d] aggregate_get default  %+v", o._statusCode, o.Payload)
}
func (o *AggregateGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregateGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}