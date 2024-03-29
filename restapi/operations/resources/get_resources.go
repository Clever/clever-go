// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetResourcesHandlerFunc turns a function with the right signature into a get resources handler
type GetResourcesHandlerFunc func(GetResourcesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourcesHandlerFunc) Handle(params GetResourcesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetResourcesHandler interface for that can handle valid get resources params
type GetResourcesHandler interface {
	Handle(GetResourcesParams, interface{}) middleware.Responder
}

// NewGetResources creates a new http.Handler for the get resources operation
func NewGetResources(ctx *middleware.Context, handler GetResourcesHandler) *GetResources {
	return &GetResources{Context: ctx, Handler: handler}
}

/* GetResources swagger:route GET /resources Resources getResources

Returns a list of resources

*/
type GetResources struct {
	Context *middleware.Context
	Handler GetResourcesHandler
}

func (o *GetResources) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResourcesParams()
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
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
