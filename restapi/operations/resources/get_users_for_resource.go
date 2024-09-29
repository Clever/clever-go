// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUsersForResourceHandlerFunc turns a function with the right signature into a get users for resource handler
type GetUsersForResourceHandlerFunc func(GetUsersForResourceParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersForResourceHandlerFunc) Handle(params GetUsersForResourceParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetUsersForResourceHandler interface for that can handle valid get users for resource params
type GetUsersForResourceHandler interface {
	Handle(GetUsersForResourceParams, interface{}) middleware.Responder
}

// NewGetUsersForResource creates a new http.Handler for the get users for resource operation
func NewGetUsersForResource(ctx *middleware.Context, handler GetUsersForResourceHandler) *GetUsersForResource {
	return &GetUsersForResource{Context: ctx, Handler: handler}
}

/*
	GetUsersForResource swagger:route GET /resources/{id}/users Resources getUsersForResource

Returns the student and/or teacher users for a resource
*/
type GetUsersForResource struct {
	Context *middleware.Context
	Handler GetUsersForResourceHandler
}

func (o *GetUsersForResource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersForResourceParams()
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
