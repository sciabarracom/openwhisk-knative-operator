// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

/*PutWebNamespacePackageNameActionNameExtensionDefault any response

swagger:response putWebNamespacePackageNameActionNameExtensionDefault
*/
type PutWebNamespacePackageNameActionNameExtensionDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPutWebNamespacePackageNameActionNameExtensionDefault creates PutWebNamespacePackageNameActionNameExtensionDefault with default headers values
func NewPutWebNamespacePackageNameActionNameExtensionDefault(code int) *PutWebNamespacePackageNameActionNameExtensionDefault {
	if code <= 0 {
		code = 500
	}

	return &PutWebNamespacePackageNameActionNameExtensionDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put web namespace package name action name extension default response
func (o *PutWebNamespacePackageNameActionNameExtensionDefault) WithStatusCode(code int) *PutWebNamespacePackageNameActionNameExtensionDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put web namespace package name action name extension default response
func (o *PutWebNamespacePackageNameActionNameExtensionDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put web namespace package name action name extension default response
func (o *PutWebNamespacePackageNameActionNameExtensionDefault) WithPayload(payload interface{}) *PutWebNamespacePackageNameActionNameExtensionDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put web namespace package name action name extension default response
func (o *PutWebNamespacePackageNameActionNameExtensionDefault) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutWebNamespacePackageNameActionNameExtensionDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
