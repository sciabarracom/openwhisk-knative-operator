// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/models"
)

// FireTriggerAcceptedCode is the HTTP code returned for type FireTriggerAccepted
const FireTriggerAcceptedCode int = 202

/*FireTriggerAccepted Activation id

swagger:response fireTriggerAccepted
*/
type FireTriggerAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ActivationID `json:"body,omitempty"`
}

// NewFireTriggerAccepted creates FireTriggerAccepted with default headers values
func NewFireTriggerAccepted() *FireTriggerAccepted {

	return &FireTriggerAccepted{}
}

// WithPayload adds the payload to the fire trigger accepted response
func (o *FireTriggerAccepted) WithPayload(payload *models.ActivationID) *FireTriggerAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fire trigger accepted response
func (o *FireTriggerAccepted) SetPayload(payload *models.ActivationID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FireTriggerAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FireTriggerNoContentCode is the HTTP code returned for type FireTriggerNoContent
const FireTriggerNoContentCode int = 204

/*FireTriggerNoContent Trigger has no active rules

swagger:response fireTriggerNoContent
*/
type FireTriggerNoContent struct {
}

// NewFireTriggerNoContent creates FireTriggerNoContent with default headers values
func NewFireTriggerNoContent() *FireTriggerNoContent {

	return &FireTriggerNoContent{}
}

// WriteResponse to the client
func (o *FireTriggerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// FireTriggerUnauthorizedCode is the HTTP code returned for type FireTriggerUnauthorized
const FireTriggerUnauthorizedCode int = 401

/*FireTriggerUnauthorized Unauthorized request

swagger:response fireTriggerUnauthorized
*/
type FireTriggerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewFireTriggerUnauthorized creates FireTriggerUnauthorized with default headers values
func NewFireTriggerUnauthorized() *FireTriggerUnauthorized {

	return &FireTriggerUnauthorized{}
}

// WithPayload adds the payload to the fire trigger unauthorized response
func (o *FireTriggerUnauthorized) WithPayload(payload *models.ErrorMessage) *FireTriggerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fire trigger unauthorized response
func (o *FireTriggerUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FireTriggerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FireTriggerNotFoundCode is the HTTP code returned for type FireTriggerNotFound
const FireTriggerNotFoundCode int = 404

/*FireTriggerNotFound Item not found

swagger:response fireTriggerNotFound
*/
type FireTriggerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewFireTriggerNotFound creates FireTriggerNotFound with default headers values
func NewFireTriggerNotFound() *FireTriggerNotFound {

	return &FireTriggerNotFound{}
}

// WithPayload adds the payload to the fire trigger not found response
func (o *FireTriggerNotFound) WithPayload(payload *models.ErrorMessage) *FireTriggerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fire trigger not found response
func (o *FireTriggerNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FireTriggerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// FireTriggerRequestTimeoutCode is the HTTP code returned for type FireTriggerRequestTimeout
const FireTriggerRequestTimeoutCode int = 408

/*FireTriggerRequestTimeout Request timed out

swagger:response fireTriggerRequestTimeout
*/
type FireTriggerRequestTimeout struct {
}

// NewFireTriggerRequestTimeout creates FireTriggerRequestTimeout with default headers values
func NewFireTriggerRequestTimeout() *FireTriggerRequestTimeout {

	return &FireTriggerRequestTimeout{}
}

// WriteResponse to the client
func (o *FireTriggerRequestTimeout) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(408)
}

// FireTriggerTooManyRequestsCode is the HTTP code returned for type FireTriggerTooManyRequests
const FireTriggerTooManyRequestsCode int = 429

/*FireTriggerTooManyRequests Too many requests in a given time period

swagger:response fireTriggerTooManyRequests
*/
type FireTriggerTooManyRequests struct {
}

// NewFireTriggerTooManyRequests creates FireTriggerTooManyRequests with default headers values
func NewFireTriggerTooManyRequests() *FireTriggerTooManyRequests {

	return &FireTriggerTooManyRequests{}
}

// WriteResponse to the client
func (o *FireTriggerTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// FireTriggerInternalServerErrorCode is the HTTP code returned for type FireTriggerInternalServerError
const FireTriggerInternalServerErrorCode int = 500

/*FireTriggerInternalServerError Server error

swagger:response fireTriggerInternalServerError
*/
type FireTriggerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewFireTriggerInternalServerError creates FireTriggerInternalServerError with default headers values
func NewFireTriggerInternalServerError() *FireTriggerInternalServerError {

	return &FireTriggerInternalServerError{}
}

// WithPayload adds the payload to the fire trigger internal server error response
func (o *FireTriggerInternalServerError) WithPayload(payload *models.ErrorMessage) *FireTriggerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the fire trigger internal server error response
func (o *FireTriggerInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FireTriggerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
