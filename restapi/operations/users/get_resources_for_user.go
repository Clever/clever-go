// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetResourcesForUserHandlerFunc turns a function with the right signature into a get resources for user handler
type GetResourcesForUserHandlerFunc func(GetResourcesForUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourcesForUserHandlerFunc) Handle(params GetResourcesForUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetResourcesForUserHandler interface for that can handle valid get resources for user params
type GetResourcesForUserHandler interface {
	Handle(GetResourcesForUserParams, interface{}) middleware.Responder
}

// NewGetResourcesForUser creates a new http.Handler for the get resources for user operation
func NewGetResourcesForUser(ctx *middleware.Context, handler GetResourcesForUserHandler) *GetResourcesForUser {
	return &GetResourcesForUser{Context: ctx, Handler: handler}
}

/* GetResourcesForUser swagger:route GET /users/{id}/resources Users getResourcesForUser

Returns the resources for a user

*/
type GetResourcesForUser struct {
	Context *middleware.Context
	Handler GetResourcesForUserHandler
}

func (o *GetResourcesForUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResourcesForUserParams()
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
