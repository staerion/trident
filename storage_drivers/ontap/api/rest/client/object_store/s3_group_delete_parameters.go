// Code generated by go-swagger; DO NOT EDIT.

package object_store

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
)

// NewS3GroupDeleteParams creates a new S3GroupDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3GroupDeleteParams() *S3GroupDeleteParams {
	return &S3GroupDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3GroupDeleteParamsWithTimeout creates a new S3GroupDeleteParams object
// with the ability to set a timeout on a request.
func NewS3GroupDeleteParamsWithTimeout(timeout time.Duration) *S3GroupDeleteParams {
	return &S3GroupDeleteParams{
		timeout: timeout,
	}
}

// NewS3GroupDeleteParamsWithContext creates a new S3GroupDeleteParams object
// with the ability to set a context for a request.
func NewS3GroupDeleteParamsWithContext(ctx context.Context) *S3GroupDeleteParams {
	return &S3GroupDeleteParams{
		Context: ctx,
	}
}

// NewS3GroupDeleteParamsWithHTTPClient creates a new S3GroupDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3GroupDeleteParamsWithHTTPClient(client *http.Client) *S3GroupDeleteParams {
	return &S3GroupDeleteParams{
		HTTPClient: client,
	}
}

/* S3GroupDeleteParams contains all the parameters to send to the API endpoint
   for the s3 group delete operation.

   Typically these are written to a http.Request.
*/
type S3GroupDeleteParams struct {

	/* ID.

	   Group identifier that identifies the unique group.
	*/
	IDPathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupDeleteParams) WithDefaults() *S3GroupDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3GroupDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 group delete params
func (o *S3GroupDeleteParams) WithTimeout(timeout time.Duration) *S3GroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 group delete params
func (o *S3GroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 group delete params
func (o *S3GroupDeleteParams) WithContext(ctx context.Context) *S3GroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 group delete params
func (o *S3GroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 group delete params
func (o *S3GroupDeleteParams) WithHTTPClient(client *http.Client) *S3GroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 group delete params
func (o *S3GroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDPathParameter adds the id to the s3 group delete params
func (o *S3GroupDeleteParams) WithIDPathParameter(id string) *S3GroupDeleteParams {
	o.SetIDPathParameter(id)
	return o
}

// SetIDPathParameter adds the id to the s3 group delete params
func (o *S3GroupDeleteParams) SetIDPathParameter(id string) {
	o.IDPathParameter = id
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 group delete params
func (o *S3GroupDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *S3GroupDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 group delete params
func (o *S3GroupDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3GroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.IDPathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}