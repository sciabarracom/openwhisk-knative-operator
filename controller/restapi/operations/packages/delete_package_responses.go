// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/models"
)

// DeletePackageOKCode is the HTTP code returned for type DeletePackageOK
const DeletePackageOKCode int = 200

/*DeletePackageOK Deleted Item

swagger:response deletePackageOK
*/
type DeletePackageOK struct {
}

// NewDeletePackageOK creates DeletePackageOK with default headers values
func NewDeletePackageOK() *DeletePackageOK {

	return &DeletePackageOK{}
}

// WriteResponse to the client
func (o *DeletePackageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeletePackageUnauthorizedCode is the HTTP code returned for type DeletePackageUnauthorized
const DeletePackageUnauthorizedCode int = 401

/*DeletePackageUnauthorized Unauthorized request

swagger:response deletePackageUnauthorized
*/
type DeletePackageUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeletePackageUnauthorized creates DeletePackageUnauthorized with default headers values
func NewDeletePackageUnauthorized() *DeletePackageUnauthorized {

	return &DeletePackageUnauthorized{}
}

// WithPayload adds the payload to the delete package unauthorized response
func (o *DeletePackageUnauthorized) WithPayload(payload *models.ErrorMessage) *DeletePackageUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package unauthorized response
func (o *DeletePackageUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePackageNotFoundCode is the HTTP code returned for type DeletePackageNotFound
const DeletePackageNotFoundCode int = 404

/*DeletePackageNotFound Item not found

swagger:response deletePackageNotFound
*/
type DeletePackageNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeletePackageNotFound creates DeletePackageNotFound with default headers values
func NewDeletePackageNotFound() *DeletePackageNotFound {

	return &DeletePackageNotFound{}
}

// WithPayload adds the payload to the delete package not found response
func (o *DeletePackageNotFound) WithPayload(payload *models.ErrorMessage) *DeletePackageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package not found response
func (o *DeletePackageNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePackageConflictCode is the HTTP code returned for type DeletePackageConflict
const DeletePackageConflictCode int = 409

/*DeletePackageConflict Conflicting item already exists

swagger:response deletePackageConflict
*/
type DeletePackageConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeletePackageConflict creates DeletePackageConflict with default headers values
func NewDeletePackageConflict() *DeletePackageConflict {

	return &DeletePackageConflict{}
}

// WithPayload adds the payload to the delete package conflict response
func (o *DeletePackageConflict) WithPayload(payload *models.ErrorMessage) *DeletePackageConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package conflict response
func (o *DeletePackageConflict) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePackageInternalServerErrorCode is the HTTP code returned for type DeletePackageInternalServerError
const DeletePackageInternalServerErrorCode int = 500

/*DeletePackageInternalServerError Server error

swagger:response deletePackageInternalServerError
*/
type DeletePackageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeletePackageInternalServerError creates DeletePackageInternalServerError with default headers values
func NewDeletePackageInternalServerError() *DeletePackageInternalServerError {

	return &DeletePackageInternalServerError{}
}

// WithPayload adds the payload to the delete package internal server error response
func (o *DeletePackageInternalServerError) WithPayload(payload *models.ErrorMessage) *DeletePackageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete package internal server error response
func (o *DeletePackageInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePackageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
