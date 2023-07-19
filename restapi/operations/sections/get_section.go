// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSectionHandlerFunc turns a function with the right signature into a get section handler
type GetSectionHandlerFunc func(GetSectionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSectionHandlerFunc) Handle(params GetSectionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSectionHandler interface for that can handle valid get section params
type GetSectionHandler interface {
	Handle(GetSectionParams, interface{}) middleware.Responder
}

// NewGetSection creates a new http.Handler for the get section operation
func NewGetSection(ctx *middleware.Context, handler GetSectionHandler) *GetSection {
	return &GetSection{Context: ctx, Handler: handler}
}

/*
	GetSection swagger:route GET /sections/{id} Sections getSection

Returns a specific section
*/
type GetSection struct {
	Context *middleware.Context
	Handler GetSectionHandler
}

func (o *GetSection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetSectionParams()
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