// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDistrictForCourseHandlerFunc turns a function with the right signature into a get district for course handler
type GetDistrictForCourseHandlerFunc func(GetDistrictForCourseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDistrictForCourseHandlerFunc) Handle(params GetDistrictForCourseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDistrictForCourseHandler interface for that can handle valid get district for course params
type GetDistrictForCourseHandler interface {
	Handle(GetDistrictForCourseParams, interface{}) middleware.Responder
}

// NewGetDistrictForCourse creates a new http.Handler for the get district for course operation
func NewGetDistrictForCourse(ctx *middleware.Context, handler GetDistrictForCourseHandler) *GetDistrictForCourse {
	return &GetDistrictForCourse{Context: ctx, Handler: handler}
}

/*
	GetDistrictForCourse swagger:route GET /courses/{id}/district Courses getDistrictForCourse

Returns the district for a course
*/
type GetDistrictForCourse struct {
	Context *middleware.Context
	Handler GetDistrictForCourseHandler
}

func (o *GetDistrictForCourse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetDistrictForCourseParams()
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
