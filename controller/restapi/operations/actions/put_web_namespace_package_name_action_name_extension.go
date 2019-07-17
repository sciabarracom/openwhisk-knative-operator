// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PutWebNamespacePackageNameActionNameExtensionHandlerFunc turns a function with the right signature into a put web namespace package name action name extension handler
type PutWebNamespacePackageNameActionNameExtensionHandlerFunc func(PutWebNamespacePackageNameActionNameExtensionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PutWebNamespacePackageNameActionNameExtensionHandlerFunc) Handle(params PutWebNamespacePackageNameActionNameExtensionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PutWebNamespacePackageNameActionNameExtensionHandler interface for that can handle valid put web namespace package name action name extension params
type PutWebNamespacePackageNameActionNameExtensionHandler interface {
	Handle(PutWebNamespacePackageNameActionNameExtensionParams, interface{}) middleware.Responder
}

// NewPutWebNamespacePackageNameActionNameExtension creates a new http.Handler for the put web namespace package name action name extension operation
func NewPutWebNamespacePackageNameActionNameExtension(ctx *middleware.Context, handler PutWebNamespacePackageNameActionNameExtensionHandler) *PutWebNamespacePackageNameActionNameExtension {
	return &PutWebNamespacePackageNameActionNameExtension{Context: ctx, Handler: handler}
}

/*PutWebNamespacePackageNameActionNameExtension swagger:route PUT /web/{namespace}/{packageName}/{actionName}.{extension} Actions putWebNamespacePackageNameActionNameExtension

PutWebNamespacePackageNameActionNameExtension put web namespace package name action name extension API

*/
type PutWebNamespacePackageNameActionNameExtension struct {
	Context *middleware.Context
	Handler PutWebNamespacePackageNameActionNameExtensionHandler
}

func (o *PutWebNamespacePackageNameActionNameExtension) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutWebNamespacePackageNameActionNameExtensionParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
