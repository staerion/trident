// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsFilterRuleCollectionGetParams creates a new EmsFilterRuleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsFilterRuleCollectionGetParams() *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithTimeout creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsFilterRuleCollectionGetParamsWithTimeout(timeout time.Duration) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithContext creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsFilterRuleCollectionGetParamsWithContext(ctx context.Context) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsFilterRuleCollectionGetParamsWithHTTPClient creates a new EmsFilterRuleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsFilterRuleCollectionGetParamsWithHTTPClient(client *http.Client) *EmsFilterRuleCollectionGetParams {
	return &EmsFilterRuleCollectionGetParams{
		HTTPClient: client,
	}
}

/* EmsFilterRuleCollectionGetParams contains all the parameters to send to the API endpoint
   for the ems filter rule collection get operation.

   Typically these are written to a http.Request.
*/
type EmsFilterRuleCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* MessageCriteriaNamePattern.

	   Filter by message_criteria.name_pattern
	*/
	MessageCriteriaNamePatternQueryParameter *string

	/* MessageCriteriaSeverities.

	   Filter by message_criteria.severities
	*/
	MessageCriteriaSeveritiesQueryParameter *string

	/* MessageCriteriaSnmpTrapTypes.

	   Filter by message_criteria.snmp_trap_types
	*/
	MessageCriteriaSnmpTrapTypesQueryParameter *string

	/* Name.

	   Filter Name
	*/
	NamePathParameter string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems filter rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleCollectionGetParams) WithDefaults() *EmsFilterRuleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems filter rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsFilterRuleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := EmsFilterRuleCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithTimeout(timeout time.Duration) *EmsFilterRuleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithContext(ctx context.Context) *EmsFilterRuleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithHTTPClient(client *http.Client) *EmsFilterRuleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithFieldsQueryParameter(fields []string) *EmsFilterRuleCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithIndexQueryParameter(index *int64) *EmsFilterRuleCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *EmsFilterRuleCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithMessageCriteriaNamePatternQueryParameter adds the messageCriteriaNamePattern to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern)
	return o
}

// SetMessageCriteriaNamePatternQueryParameter adds the messageCriteriaNamePattern to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaNamePatternQueryParameter(messageCriteriaNamePattern *string) {
	o.MessageCriteriaNamePatternQueryParameter = messageCriteriaNamePattern
}

// WithMessageCriteriaSeveritiesQueryParameter adds the messageCriteriaSeverities to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities)
	return o
}

// SetMessageCriteriaSeveritiesQueryParameter adds the messageCriteriaSeverities to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaSeveritiesQueryParameter(messageCriteriaSeverities *string) {
	o.MessageCriteriaSeveritiesQueryParameter = messageCriteriaSeverities
}

// WithMessageCriteriaSnmpTrapTypesQueryParameter adds the messageCriteriaSnmpTrapTypes to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes *string) *EmsFilterRuleCollectionGetParams {
	o.SetMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes)
	return o
}

// SetMessageCriteriaSnmpTrapTypesQueryParameter adds the messageCriteriaSnmpTrapTypes to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetMessageCriteriaSnmpTrapTypesQueryParameter(messageCriteriaSnmpTrapTypes *string) {
	o.MessageCriteriaSnmpTrapTypesQueryParameter = messageCriteriaSnmpTrapTypes
}

// WithNamePathParameter adds the name to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithNamePathParameter(name string) *EmsFilterRuleCollectionGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *EmsFilterRuleCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *EmsFilterRuleCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *EmsFilterRuleCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithTypeQueryParameter adds the typeVar to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) WithTypeQueryParameter(typeVar *string) *EmsFilterRuleCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the ems filter rule collection get params
func (o *EmsFilterRuleCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *EmsFilterRuleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaNamePatternQueryParameter != nil {

		// query param message_criteria.name_pattern
		var qrMessageCriteriaNamePattern string

		if o.MessageCriteriaNamePatternQueryParameter != nil {
			qrMessageCriteriaNamePattern = *o.MessageCriteriaNamePatternQueryParameter
		}
		qMessageCriteriaNamePattern := qrMessageCriteriaNamePattern
		if qMessageCriteriaNamePattern != "" {

			if err := r.SetQueryParam("message_criteria.name_pattern", qMessageCriteriaNamePattern); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaSeveritiesQueryParameter != nil {

		// query param message_criteria.severities
		var qrMessageCriteriaSeverities string

		if o.MessageCriteriaSeveritiesQueryParameter != nil {
			qrMessageCriteriaSeverities = *o.MessageCriteriaSeveritiesQueryParameter
		}
		qMessageCriteriaSeverities := qrMessageCriteriaSeverities
		if qMessageCriteriaSeverities != "" {

			if err := r.SetQueryParam("message_criteria.severities", qMessageCriteriaSeverities); err != nil {
				return err
			}
		}
	}

	if o.MessageCriteriaSnmpTrapTypesQueryParameter != nil {

		// query param message_criteria.snmp_trap_types
		var qrMessageCriteriaSnmpTrapTypes string

		if o.MessageCriteriaSnmpTrapTypesQueryParameter != nil {
			qrMessageCriteriaSnmpTrapTypes = *o.MessageCriteriaSnmpTrapTypesQueryParameter
		}
		qMessageCriteriaSnmpTrapTypes := qrMessageCriteriaSnmpTrapTypes
		if qMessageCriteriaSnmpTrapTypes != "" {

			if err := r.SetQueryParam("message_criteria.snmp_trap_types", qMessageCriteriaSnmpTrapTypes); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsFilterRuleCollectionGet binds the parameter fields
func (o *EmsFilterRuleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamEmsFilterRuleCollectionGet binds the parameter order_by
func (o *EmsFilterRuleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}