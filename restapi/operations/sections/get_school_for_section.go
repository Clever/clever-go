// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetSchoolForSectionHandlerFunc turns a function with the right signature into a get school for section handler
type GetSchoolForSectionHandlerFunc func(GetSchoolForSectionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetSchoolForSectionHandlerFunc) Handle(params GetSchoolForSectionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetSchoolForSectionHandler interface for that can handle valid get school for section params
type GetSchoolForSectionHandler interface {
	Handle(GetSchoolForSectionParams, interface{}) middleware.Responder
}

// NewGetSchoolForSection creates a new http.Handler for the get school for section operation
func NewGetSchoolForSection(ctx *middleware.Context, handler GetSchoolForSectionHandler) *GetSchoolForSection {
	return &GetSchoolForSection{Context: ctx, Handler: handler}
}

/* GetSchoolForSection swagger:route GET /sections/{id}/school Sections getSchoolForSection

Returns the school for a section

*/
type GetSchoolForSection struct {
	Context *middleware.Context
	Handler GetSchoolForSectionHandler
}

func (o *GetSchoolForSection) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetSchoolForSectionParams()
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
