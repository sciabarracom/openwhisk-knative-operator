// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*GetWebNamespacePackageNameActionNameExtensionDefault any response

swagger:response getWebNamespacePackageNameActionNameExtensionDefault
*/
type GetWebNamespacePackageNameActionNameExtensionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewGetWebNamespacePackageNameActionNameExtensionDefault creates GetWebNamespacePackageNameActionNameExtensionDefault with default headers values
func NewGetWebNamespacePackageNameActionNameExtensionDefault(code int) *GetWebNamespacePackageNameActionNameExtensionDefault {
	if code <= 0 {
		code = 500
	}

	return &GetWebNamespacePackageNameActionNameExtensionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get web namespace package name action name extension default response
func (o *GetWebNamespacePackageNameActionNameExtensionDefault) WithStatusCode(code int) *GetWebNamespacePackageNameActionNameExtensionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get web namespace package name action name extension default response
func (o *GetWebNamespacePackageNameActionNameExtensionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get web namespace package name action name extension default response
func (o *GetWebNamespacePackageNameActionNameExtensionDefault) WithPayload(payload interface{}) *GetWebNamespacePackageNameActionNameExtensionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get web namespace package name action name extension default response
func (o *GetWebNamespacePackageNameActionNameExtensionDefault) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetWebNamespacePackageNameActionNameExtensionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
