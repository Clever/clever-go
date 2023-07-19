// Code generated by go-swagger; DO NOT EDIT.

package terms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTermsHandlerFunc turns a function with the right signature into a get terms handler
type GetTermsHandlerFunc func(GetTermsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTermsHandlerFunc) Handle(params GetTermsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTermsHandler interface for that can handle valid get terms params
type GetTermsHandler interface {
	Handle(GetTermsParams, interface{}) middleware.Responder
}

// NewGetTerms creates a new http.Handler for the get terms operation
func NewGetTerms(ctx *middleware.Context, handler GetTermsHandler) *GetTerms {
	return &GetTerms{Context: ctx, Handler: handler}
}

/*
	GetTerms swagger:route GET /terms Terms getTerms

Returns a list of terms
*/
type GetTerms struct {
	Context *middleware.Context
	Handler GetTermsHandler
}

func (o *GetTerms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTermsParams()
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