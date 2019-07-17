// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteTriggerParams creates a new DeleteTriggerParams object
// no default values defined in spec.
func NewDeleteTriggerParams() DeleteTriggerParams {

	return DeleteTriggerParams{}
}

// DeleteTriggerParams contains all the bound params for the delete trigger operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteTrigger
type DeleteTriggerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The entity namespace
	  Required: true
	  In: path
	*/
	Namespace string
	/*Name of trigger to delete
	  Required: true
	  In: path
	*/
	TriggerName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteTriggerParams() beforehand.
func (o *DeleteTriggerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rNamespace, rhkNamespace, _ := route.Params.GetOK("namespace")
	if err := o.bindNamespace(rNamespace, rhkNamespace, route.Formats); err != nil {
		res = append(res, err)
	}

	rTriggerName, rhkTriggerName, _ := route.Params.GetOK("triggerName")
	if err := o.bindTriggerName(rTriggerName, rhkTriggerName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindNamespace binds and validates parameter Namespace from path.
func (o *DeleteTriggerParams) bindNamespace(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Namespace = raw

	return nil
}

// bindTriggerName binds and validates parameter TriggerName from path.
func (o *DeleteTriggerParams) bindTriggerName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.TriggerName = raw

	return nil
}
