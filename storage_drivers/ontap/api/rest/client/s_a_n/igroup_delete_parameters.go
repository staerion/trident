// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIgroupDeleteParams creates a new IgroupDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupDeleteParams() *IgroupDeleteParams {
	return &IgroupDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupDeleteParamsWithTimeout creates a new IgroupDeleteParams object
// with the ability to set a timeout on a request.
func NewIgroupDeleteParamsWithTimeout(timeout time.Duration) *IgroupDeleteParams {
	return &IgroupDeleteParams{
		timeout: timeout,
	}
}

// NewIgroupDeleteParamsWithContext creates a new IgroupDeleteParams object
// with the ability to set a context for a request.
func NewIgroupDeleteParamsWithContext(ctx context.Context) *IgroupDeleteParams {
	return &IgroupDeleteParams{
		Context: ctx,
	}
}

// NewIgroupDeleteParamsWithHTTPClient creates a new IgroupDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupDeleteParamsWithHTTPClient(client *http.Client) *IgroupDeleteParams {
	return &IgroupDeleteParams{
		HTTPClient: client,
	}
}

/*
IgroupDeleteParams contains all the parameters to send to the API endpoint

	for the igroup delete operation.

	Typically these are written to a http.Request.
*/
type IgroupDeleteParams struct {

	/* AllowDeleteWhileMapped.

	     Allows the deletion of a mapped initiator group.<br/>
	Deleting a mapped initiator group means that the LUNs, to which the
	initiator group is mapped, are no longer available to the initiators.
	This might cause a disruption in the availability of data.<br/>
	<b>This parameter should be used with caution.</b>

	*/
	AllowDeleteWhileMapped *bool

	/* UUID.

	   The unique identifier of the initiator group.

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupDeleteParams) WithDefaults() *IgroupDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupDeleteParams) SetDefaults() {
	var (
		allowDeleteWhileMappedDefault = bool(false)
	)

	val := IgroupDeleteParams{
		AllowDeleteWhileMapped: &allowDeleteWhileMappedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the igroup delete params
func (o *IgroupDeleteParams) WithTimeout(timeout time.Duration) *IgroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup delete params
func (o *IgroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup delete params
func (o *IgroupDeleteParams) WithContext(ctx context.Context) *IgroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup delete params
func (o *IgroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup delete params
func (o *IgroupDeleteParams) WithHTTPClient(client *http.Client) *IgroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup delete params
func (o *IgroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the igroup delete params
func (o *IgroupDeleteParams) WithAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) *IgroupDeleteParams {
	o.SetAllowDeleteWhileMapped(allowDeleteWhileMapped)
	return o
}

// SetAllowDeleteWhileMapped adds the allowDeleteWhileMapped to the igroup delete params
func (o *IgroupDeleteParams) SetAllowDeleteWhileMapped(allowDeleteWhileMapped *bool) {
	o.AllowDeleteWhileMapped = allowDeleteWhileMapped
}

// WithUUID adds the uuid to the igroup delete params
func (o *IgroupDeleteParams) WithUUID(uuid string) *IgroupDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the igroup delete params
func (o *IgroupDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowDeleteWhileMapped != nil {

		// query param allow_delete_while_mapped
		var qrAllowDeleteWhileMapped bool

		if o.AllowDeleteWhileMapped != nil {
			qrAllowDeleteWhileMapped = *o.AllowDeleteWhileMapped
		}
		qAllowDeleteWhileMapped := swag.FormatBool(qrAllowDeleteWhileMapped)
		if qAllowDeleteWhileMapped != "" {

			if err := r.SetQueryParam("allow_delete_while_mapped", qAllowDeleteWhileMapped); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
