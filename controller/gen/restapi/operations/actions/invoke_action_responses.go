// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// InvokeActionOKCode is the HTTP code returned for type InvokeActionOK
const InvokeActionOKCode int = 200

/*InvokeActionOK Successful activation

swagger:response invokeActionOK
*/
type InvokeActionOK struct {
}

// NewInvokeActionOK creates InvokeActionOK with default headers values
func NewInvokeActionOK() *InvokeActionOK {

	return &InvokeActionOK{}
}

// WriteResponse to the client
func (o *InvokeActionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// InvokeActionAcceptedCode is the HTTP code returned for type InvokeActionAccepted
const InvokeActionAcceptedCode int = 202

/*InvokeActionAccepted Accepted activation request

swagger:response invokeActionAccepted
*/
type InvokeActionAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ActivationID `json:"body,omitempty"`
}

// NewInvokeActionAccepted creates InvokeActionAccepted with default headers values
func NewInvokeActionAccepted() *InvokeActionAccepted {

	return &InvokeActionAccepted{}
}

// WithPayload adds the payload to the invoke action accepted response
func (o *InvokeActionAccepted) WithPayload(payload *models.ActivationID) *InvokeActionAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the invoke action accepted response
func (o *InvokeActionAccepted) SetPayload(payload *models.ActivationID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InvokeActionAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InvokeActionUnauthorizedCode is the HTTP code returned for type InvokeActionUnauthorized
const InvokeActionUnauthorizedCode int = 401

/*InvokeActionUnauthorized Unauthorized request

swagger:response invokeActionUnauthorized
*/
type InvokeActionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewInvokeActionUnauthorized creates InvokeActionUnauthorized with default headers values
func NewInvokeActionUnauthorized() *InvokeActionUnauthorized {

	return &InvokeActionUnauthorized{}
}

// WithPayload adds the payload to the invoke action unauthorized response
func (o *InvokeActionUnauthorized) WithPayload(payload *models.ErrorMessage) *InvokeActionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the invoke action unauthorized response
func (o *InvokeActionUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InvokeActionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InvokeActionForbiddenCode is the HTTP code returned for type InvokeActionForbidden
const InvokeActionForbiddenCode int = 403

/*InvokeActionForbidden Unauthorized request

swagger:response invokeActionForbidden
*/
type InvokeActionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewInvokeActionForbidden creates InvokeActionForbidden with default headers values
func NewInvokeActionForbidden() *InvokeActionForbidden {

	return &InvokeActionForbidden{}
}

// WithPayload adds the payload to the invoke action forbidden response
func (o *InvokeActionForbidden) WithPayload(payload *models.ErrorMessage) *InvokeActionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the invoke action forbidden response
func (o *InvokeActionForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InvokeActionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InvokeActionNotFoundCode is the HTTP code returned for type InvokeActionNotFound
const InvokeActionNotFoundCode int = 404

/*InvokeActionNotFound Item not found

swagger:response invokeActionNotFound
*/
type InvokeActionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewInvokeActionNotFound creates InvokeActionNotFound with default headers values
func NewInvokeActionNotFound() *InvokeActionNotFound {

	return &InvokeActionNotFound{}
}

// WithPayload adds the payload to the invoke action not found response
func (o *InvokeActionNotFound) WithPayload(payload *models.ErrorMessage) *InvokeActionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the invoke action not found response
func (o *InvokeActionNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InvokeActionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InvokeActionRequestTimeoutCode is the HTTP code returned for type InvokeActionRequestTimeout
const InvokeActionRequestTimeoutCode int = 408

/*InvokeActionRequestTimeout Request timed out

swagger:response invokeActionRequestTimeout
*/
type InvokeActionRequestTimeout struct {
}

// NewInvokeActionRequestTimeout creates InvokeActionRequestTimeout with default headers values
func NewInvokeActionRequestTimeout() *InvokeActionRequestTimeout {

	return &InvokeActionRequestTimeout{}
}

// WriteResponse to the client
func (o *InvokeActionRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(408)
}

// InvokeActionTooManyRequestsCode is the HTTP code returned for type InvokeActionTooManyRequests
const InvokeActionTooManyRequestsCode int = 429

/*InvokeActionTooManyRequests Too many requests in a given time period

swagger:response invokeActionTooManyRequests
*/
type InvokeActionTooManyRequests struct {
}

// NewInvokeActionTooManyRequests creates InvokeActionTooManyRequests with default headers values
func NewInvokeActionTooManyRequests() *InvokeActionTooManyRequests {

	return &InvokeActionTooManyRequests{}
}

// WriteResponse to the client
func (o *InvokeActionTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// InvokeActionInternalServerErrorCode is the HTTP code returned for type InvokeActionInternalServerError
const InvokeActionInternalServerErrorCode int = 500

/*InvokeActionInternalServerError Server error

swagger:response invokeActionInternalServerError
*/
type InvokeActionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewInvokeActionInternalServerError creates InvokeActionInternalServerError with default headers values
func NewInvokeActionInternalServerError() *InvokeActionInternalServerError {

	return &InvokeActionInternalServerError{}
}

// WithPayload adds the payload to the invoke action internal server error response
func (o *InvokeActionInternalServerError) WithPayload(payload *models.ErrorMessage) *InvokeActionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the invoke action internal server error response
func (o *InvokeActionInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InvokeActionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// InvokeActionBadGatewayCode is the HTTP code returned for type InvokeActionBadGateway
const InvokeActionBadGatewayCode int = 502

/*InvokeActionBadGateway Activation produced an application error

swagger:response invokeActionBadGateway
*/
type InvokeActionBadGateway struct {
}

// NewInvokeActionBadGateway creates InvokeActionBadGateway with default headers values
func NewInvokeActionBadGateway() *InvokeActionBadGateway {

	return &InvokeActionBadGateway{}
}

// WriteResponse to the client
func (o *InvokeActionBadGateway) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(502)
}
