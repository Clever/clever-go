// Code generated by go-swagger; DO NOT EDIT.

package courses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetResourcesForCourseHandlerFunc turns a function with the right signature into a get resources for course handler
type GetResourcesForCourseHandlerFunc func(GetResourcesForCourseParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourcesForCourseHandlerFunc) Handle(params GetResourcesForCourseParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetResourcesForCourseHandler interface for that can handle valid get resources for course params
type GetResourcesForCourseHandler interface {
	Handle(GetResourcesForCourseParams, interface{}) middleware.Responder
}

// NewGetResourcesForCourse creates a new http.Handler for the get resources for course operation
func NewGetResourcesForCourse(ctx *middleware.Context, handler GetResourcesForCourseHandler) *GetResourcesForCourse {
	return &GetResourcesForCourse{Context: ctx, Handler: handler}
}

/* GetResourcesForCourse swagger:route GET /courses/{id}/resources Courses getResourcesForCourse

Returns the resources for a course

*/
type GetResourcesForCourse struct {
	Context *middleware.Context
	Handler GetResourcesForCourseHandler
}

func (o *GetResourcesForCourse) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResourcesForCourseParams()
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
