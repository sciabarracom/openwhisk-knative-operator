// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/models"
)

// DeleteActionOKCode is the HTTP code returned for type DeleteActionOK
const DeleteActionOKCode int = 200

/*DeleteActionOK Deleted Item

swagger:response deleteActionOK
*/
type DeleteActionOK struct {
}

// NewDeleteActionOK creates DeleteActionOK with default headers values
func NewDeleteActionOK() *DeleteActionOK {

	return &DeleteActionOK{}
}

// WriteResponse to the client
func (o *DeleteActionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteActionBadRequestCode is the HTTP code returned for type DeleteActionBadRequest
const DeleteActionBadRequestCode int = 400

/*DeleteActionBadRequest Bad request

swagger:response deleteActionBadRequest
*/
type DeleteActionBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionBadRequest creates DeleteActionBadRequest with default headers values
func NewDeleteActionBadRequest() *DeleteActionBadRequest {

	return &DeleteActionBadRequest{}
}

// WithPayload adds the payload to the delete action bad request response
func (o *DeleteActionBadRequest) WithPayload(payload *models.ErrorMessage) *DeleteActionBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action bad request response
func (o *DeleteActionBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionUnauthorizedCode is the HTTP code returned for type DeleteActionUnauthorized
const DeleteActionUnauthorizedCode int = 401

/*DeleteActionUnauthorized Unauthorized request

swagger:response deleteActionUnauthorized
*/
type DeleteActionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionUnauthorized creates DeleteActionUnauthorized with default headers values
func NewDeleteActionUnauthorized() *DeleteActionUnauthorized {

	return &DeleteActionUnauthorized{}
}

// WithPayload adds the payload to the delete action unauthorized response
func (o *DeleteActionUnauthorized) WithPayload(payload *models.ErrorMessage) *DeleteActionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action unauthorized response
func (o *DeleteActionUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionForbiddenCode is the HTTP code returned for type DeleteActionForbidden
const DeleteActionForbiddenCode int = 403

/*DeleteActionForbidden Unauthorized request

swagger:response deleteActionForbidden
*/
type DeleteActionForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionForbidden creates DeleteActionForbidden with default headers values
func NewDeleteActionForbidden() *DeleteActionForbidden {

	return &DeleteActionForbidden{}
}

// WithPayload adds the payload to the delete action forbidden response
func (o *DeleteActionForbidden) WithPayload(payload *models.ErrorMessage) *DeleteActionForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action forbidden response
func (o *DeleteActionForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionNotFoundCode is the HTTP code returned for type DeleteActionNotFound
const DeleteActionNotFoundCode int = 404

/*DeleteActionNotFound Item not found

swagger:response deleteActionNotFound
*/
type DeleteActionNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionNotFound creates DeleteActionNotFound with default headers values
func NewDeleteActionNotFound() *DeleteActionNotFound {

	return &DeleteActionNotFound{}
}

// WithPayload adds the payload to the delete action not found response
func (o *DeleteActionNotFound) WithPayload(payload *models.ErrorMessage) *DeleteActionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action not found response
func (o *DeleteActionNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionConflictCode is the HTTP code returned for type DeleteActionConflict
const DeleteActionConflictCode int = 409

/*DeleteActionConflict Conflicting item already exists

swagger:response deleteActionConflict
*/
type DeleteActionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionConflict creates DeleteActionConflict with default headers values
func NewDeleteActionConflict() *DeleteActionConflict {

	return &DeleteActionConflict{}
}

// WithPayload adds the payload to the delete action conflict response
func (o *DeleteActionConflict) WithPayload(payload *models.ErrorMessage) *DeleteActionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action conflict response
func (o *DeleteActionConflict) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteActionInternalServerErrorCode is the HTTP code returned for type DeleteActionInternalServerError
const DeleteActionInternalServerErrorCode int = 500

/*DeleteActionInternalServerError Server error

swagger:response deleteActionInternalServerError
*/
type DeleteActionInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewDeleteActionInternalServerError creates DeleteActionInternalServerError with default headers values
func NewDeleteActionInternalServerError() *DeleteActionInternalServerError {

	return &DeleteActionInternalServerError{}
}

// WithPayload adds the payload to the delete action internal server error response
func (o *DeleteActionInternalServerError) WithPayload(payload *models.ErrorMessage) *DeleteActionInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete action internal server error response
func (o *DeleteActionInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteActionInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
