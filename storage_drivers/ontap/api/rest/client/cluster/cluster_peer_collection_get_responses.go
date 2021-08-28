// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ClusterPeerCollectionGetReader is a Reader for the ClusterPeerCollectionGet structure.
type ClusterPeerCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterPeerCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterPeerCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterPeerCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterPeerCollectionGetOK creates a ClusterPeerCollectionGetOK with default headers values
func NewClusterPeerCollectionGetOK() *ClusterPeerCollectionGetOK {
	return &ClusterPeerCollectionGetOK{}
}

/* ClusterPeerCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type ClusterPeerCollectionGetOK struct {
	Payload *models.ClusterPeerResponse
}

func (o *ClusterPeerCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/peers][%d] clusterPeerCollectionGetOK  %+v", 200, o.Payload)
}
func (o *ClusterPeerCollectionGetOK) GetPayload() *models.ClusterPeerResponse {
	return o.Payload
}

func (o *ClusterPeerCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPeerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterPeerCollectionGetDefault creates a ClusterPeerCollectionGetDefault with default headers values
func NewClusterPeerCollectionGetDefault(code int) *ClusterPeerCollectionGetDefault {
	return &ClusterPeerCollectionGetDefault{
		_statusCode: code,
	}
}

/* ClusterPeerCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterPeerCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster peer collection get default response
func (o *ClusterPeerCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *ClusterPeerCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/peers][%d] cluster_peer_collection_get default  %+v", o._statusCode, o.Payload)
}
func (o *ClusterPeerCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterPeerCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}