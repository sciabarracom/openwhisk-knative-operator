// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// GetPackageByNameOKCode is the HTTP code returned for type GetPackageByNameOK
const GetPackageByNameOKCode int = 200

/*GetPackageByNameOK Returned package

swagger:response getPackageByNameOK
*/
type GetPackageByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Package `json:"body,omitempty"`
}

// NewGetPackageByNameOK creates GetPackageByNameOK with default headers values
func NewGetPackageByNameOK() *GetPackageByNameOK {

	return &GetPackageByNameOK{}
}

// WithPayload adds the payload to the get package by name o k response
func (o *GetPackageByNameOK) WithPayload(payload *models.Package) *GetPackageByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name o k response
func (o *GetPackageByNameOK) SetPayload(payload *models.Package) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPackageByNameUnauthorizedCode is the HTTP code returned for type GetPackageByNameUnauthorized
const GetPackageByNameUnauthorizedCode int = 401

/*GetPackageByNameUnauthorized Unauthorized request

swagger:response getPackageByNameUnauthorized
*/
type GetPackageByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetPackageByNameUnauthorized creates GetPackageByNameUnauthorized with default headers values
func NewGetPackageByNameUnauthorized() *GetPackageByNameUnauthorized {

	return &GetPackageByNameUnauthorized{}
}

// WithPayload adds the payload to the get package by name unauthorized response
func (o *GetPackageByNameUnauthorized) WithPayload(payload *models.ErrorMessage) *GetPackageByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name unauthorized response
func (o *GetPackageByNameUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPackageByNameForbiddenCode is the HTTP code returned for type GetPackageByNameForbidden
const GetPackageByNameForbiddenCode int = 403

/*GetPackageByNameForbidden Unauthorized request

swagger:response getPackageByNameForbidden
*/
type GetPackageByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetPackageByNameForbidden creates GetPackageByNameForbidden with default headers values
func NewGetPackageByNameForbidden() *GetPackageByNameForbidden {

	return &GetPackageByNameForbidden{}
}

// WithPayload adds the payload to the get package by name forbidden response
func (o *GetPackageByNameForbidden) WithPayload(payload *models.ErrorMessage) *GetPackageByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name forbidden response
func (o *GetPackageByNameForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPackageByNameNotFoundCode is the HTTP code returned for type GetPackageByNameNotFound
const GetPackageByNameNotFoundCode int = 404

/*GetPackageByNameNotFound Item not found

swagger:response getPackageByNameNotFound
*/
type GetPackageByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetPackageByNameNotFound creates GetPackageByNameNotFound with default headers values
func NewGetPackageByNameNotFound() *GetPackageByNameNotFound {

	return &GetPackageByNameNotFound{}
}

// WithPayload adds the payload to the get package by name not found response
func (o *GetPackageByNameNotFound) WithPayload(payload *models.ErrorMessage) *GetPackageByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name not found response
func (o *GetPackageByNameNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPackageByNameConflictCode is the HTTP code returned for type GetPackageByNameConflict
const GetPackageByNameConflictCode int = 409

/*GetPackageByNameConflict Conflicting item already exists

swagger:response getPackageByNameConflict
*/
type GetPackageByNameConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetPackageByNameConflict creates GetPackageByNameConflict with default headers values
func NewGetPackageByNameConflict() *GetPackageByNameConflict {

	return &GetPackageByNameConflict{}
}

// WithPayload adds the payload to the get package by name conflict response
func (o *GetPackageByNameConflict) WithPayload(payload *models.ErrorMessage) *GetPackageByNameConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name conflict response
func (o *GetPackageByNameConflict) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPackageByNameInternalServerErrorCode is the HTTP code returned for type GetPackageByNameInternalServerError
const GetPackageByNameInternalServerErrorCode int = 500

/*GetPackageByNameInternalServerError Server error

swagger:response getPackageByNameInternalServerError
*/
type GetPackageByNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetPackageByNameInternalServerError creates GetPackageByNameInternalServerError with default headers values
func NewGetPackageByNameInternalServerError() *GetPackageByNameInternalServerError {

	return &GetPackageByNameInternalServerError{}
}

// WithPayload adds the payload to the get package by name internal server error response
func (o *GetPackageByNameInternalServerError) WithPayload(payload *models.ErrorMessage) *GetPackageByNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get package by name internal server error response
func (o *GetPackageByNameInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPackageByNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}