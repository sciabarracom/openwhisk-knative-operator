// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAllTriggersHandlerFunc turns a function with the right signature into a get all triggers handler
type GetAllTriggersHandlerFunc func(GetAllTriggersParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllTriggersHandlerFunc) Handle(params GetAllTriggersParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllTriggersHandler interface for that can handle valid get all triggers params
type GetAllTriggersHandler interface {
	Handle(GetAllTriggersParams, interface{}) middleware.Responder
}

// NewGetAllTriggers creates a new http.Handler for the get all triggers operation
func NewGetAllTriggers(ctx *middleware.Context, handler GetAllTriggersHandler) *GetAllTriggers {
	return &GetAllTriggers{Context: ctx, Handler: handler}
}

/*GetAllTriggers swagger:route GET /namespaces/{namespace}/triggers Triggers getAllTriggers

Get all triggers

Get all triggers

*/
type GetAllTriggers struct {
	Context *middleware.Context
	Handler GetAllTriggersHandler
}

func (o *GetAllTriggers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllTriggersParams()

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
