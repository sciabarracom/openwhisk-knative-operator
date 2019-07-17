// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/models"
)

// GetActionInPackageByNameOKCode is the HTTP code returned for type GetActionInPackageByNameOK
const GetActionInPackageByNameOKCode int = 200

/*GetActionInPackageByNameOK Returned action

swagger:response getActionInPackageByNameOK
*/
type GetActionInPackageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Action `json:"body,omitempty"`
}

// NewGetActionInPackageByNameOK creates GetActionInPackageByNameOK with default headers values
func NewGetActionInPackageByNameOK() *GetActionInPackageByNameOK {

	return &GetActionInPackageByNameOK{}
}

// WithPayload adds the payload to the get action in package by name o k response
func (o *GetActionInPackageByNameOK) WithPayload(payload *models.Action) *GetActionInPackageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action in package by name o k response
func (o *GetActionInPackageByNameOK) SetPayload(payload *models.Action) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionInPackageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionInPackageByNameUnauthorizedCode is the HTTP code returned for type GetActionInPackageByNameUnauthorized
const GetActionInPackageByNameUnauthorizedCode int = 401

/*GetActionInPackageByNameUnauthorized Unauthorized request

swagger:response getActionInPackageByNameUnauthorized
*/
type GetActionInPackageByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionInPackageByNameUnauthorized creates GetActionInPackageByNameUnauthorized with default headers values
func NewGetActionInPackageByNameUnauthorized() *GetActionInPackageByNameUnauthorized {

	return &GetActionInPackageByNameUnauthorized{}
}

// WithPayload adds the payload to the get action in package by name unauthorized response
func (o *GetActionInPackageByNameUnauthorized) WithPayload(payload *models.ErrorMessage) *GetActionInPackageByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action in package by name unauthorized response
func (o *GetActionInPackageByNameUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionInPackageByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionInPackageByNameForbiddenCode is the HTTP code returned for type GetActionInPackageByNameForbidden
const GetActionInPackageByNameForbiddenCode int = 403

/*GetActionInPackageByNameForbidden Unauthorized request

swagger:response getActionInPackageByNameForbidden
*/
type GetActionInPackageByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionInPackageByNameForbidden creates GetActionInPackageByNameForbidden with default headers values
func NewGetActionInPackageByNameForbidden() *GetActionInPackageByNameForbidden {

	return &GetActionInPackageByNameForbidden{}
}

// WithPayload adds the payload to the get action in package by name forbidden response
func (o *GetActionInPackageByNameForbidden) WithPayload(payload *models.ErrorMessage) *GetActionInPackageByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action in package by name forbidden response
func (o *GetActionInPackageByNameForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionInPackageByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionInPackageByNameNotFoundCode is the HTTP code returned for type GetActionInPackageByNameNotFound
const GetActionInPackageByNameNotFoundCode int = 404

/*GetActionInPackageByNameNotFound Item not found

swagger:response getActionInPackageByNameNotFound
*/
type GetActionInPackageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionInPackageByNameNotFound creates GetActionInPackageByNameNotFound with default headers values
func NewGetActionInPackageByNameNotFound() *GetActionInPackageByNameNotFound {

	return &GetActionInPackageByNameNotFound{}
}

// WithPayload adds the payload to the get action in package by name not found response
func (o *GetActionInPackageByNameNotFound) WithPayload(payload *models.ErrorMessage) *GetActionInPackageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action in package by name not found response
func (o *GetActionInPackageByNameNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionInPackageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionInPackageByNameInternalServerErrorCode is the HTTP code returned for type GetActionInPackageByNameInternalServerError
const GetActionInPackageByNameInternalServerErrorCode int = 500

/*GetActionInPackageByNameInternalServerError Server error

swagger:response getActionInPackageByNameInternalServerError
*/
type GetActionInPackageByNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionInPackageByNameInternalServerError creates GetActionInPackageByNameInternalServerError with default headers values
func NewGetActionInPackageByNameInternalServerError() *GetActionInPackageByNameInternalServerError {

	return &GetActionInPackageByNameInternalServerError{}
}

// WithPayload adds the payload to the get action in package by name internal server error response
func (o *GetActionInPackageByNameInternalServerError) WithPayload(payload *models.ErrorMessage) *GetActionInPackageByNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action in package by name internal server error response
func (o *GetActionInPackageByNameInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionInPackageByNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
