// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTeachersForUserHandlerFunc turns a function with the right signature into a get teachers for user handler
type GetTeachersForUserHandlerFunc func(GetTeachersForUserParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeachersForUserHandlerFunc) Handle(params GetTeachersForUserParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTeachersForUserHandler interface for that can handle valid get teachers for user params
type GetTeachersForUserHandler interface {
	Handle(GetTeachersForUserParams, interface{}) middleware.Responder
}

// NewGetTeachersForUser creates a new http.Handler for the get teachers for user operation
func NewGetTeachersForUser(ctx *middleware.Context, handler GetTeachersForUserHandler) *GetTeachersForUser {
	return &GetTeachersForUser{Context: ctx, Handler: handler}
}

/*
	GetTeachersForUser swagger:route GET /users/{id}/myteachers Users getTeachersForUser

Returns the teacher users for a student user
*/
type GetTeachersForUser struct {
	Context *middleware.Context
	Handler GetTeachersForUserHandler
}

func (o *GetTeachersForUser) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTeachersForUserParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
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
