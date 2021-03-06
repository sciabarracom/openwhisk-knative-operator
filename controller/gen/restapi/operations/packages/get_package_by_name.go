// Code generated by go-swagger; DO NOT EDIT.

package packages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/sciabarracom/openwhisk-knative/controller/gen/models"
)

// GetPackageByNameHandlerFunc turns a function with the right signature into a get package by name handler
type GetPackageByNameHandlerFunc func(GetPackageByNameParams, *models.Auth) middleware.Responder

// Handle executing the request and returning a response
func (fn GetPackageByNameHandlerFunc) Handle(params GetPackageByNameParams, principal *models.Auth) middleware.Responder {
	return fn(params, principal)
}

// GetPackageByNameHandler interface for that can handle valid get package by name params
type GetPackageByNameHandler interface {
	Handle(GetPackageByNameParams, *models.Auth) middleware.Responder
}

// NewGetPackageByName creates a new http.Handler for the get package by name operation
func NewGetPackageByName(ctx *middleware.Context, handler GetPackageByNameHandler) *GetPackageByName {
	return &GetPackageByName{Context: ctx, Handler: handler}
}

/*GetPackageByName swagger:route GET /namespaces/{namespace}/packages/{packageName} Packages getPackageByName

Get package information

Get package information.

*/
type GetPackageByName struct {
	Context *middleware.Context
	Handler GetPackageByNameHandler
}

func (o *GetPackageByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetPackageByNameParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Auth
	if uprinc != nil {
		principal = uprinc.(*models.Auth) // this is really a models.Auth, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
