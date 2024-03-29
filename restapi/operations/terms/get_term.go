// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTermHandlerFunc turns a function with the right signature into a get term handler
type GetTermHandlerFunc func(GetTermParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTermHandlerFunc) Handle(params GetTermParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTermHandler interface for that can handle valid get term params
type GetTermHandler interface {
	Handle(GetTermParams, interface{}) middleware.Responder
}

// NewGetTerm creates a new http.Handler for the get term operation
func NewGetTerm(ctx *middleware.Context, handler GetTermHandler) *GetTerm {
	return &GetTerm{Context: ctx, Handler: handler}
}

/* GetTerm swagger:route GET /terms/{id} Terms getTerm

Returns a specific term

*/
type GetTerm struct {
	Context *middleware.Context
	Handler GetTermHandler
}

func (o *GetTerm) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTermParams()
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
