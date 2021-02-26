// Code generated by go-swagger; DO NOT EDIT.

package sections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetUsersForSectionURL generates an URL for the get users for section operation
type GetUsersForSectionURL struct {
	ID string

	EndingBefore  *string
	Limit         *int64
	Primary       *string
	Role          *string
	StartingAfter *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersForSectionURL) WithBasePath(bp string) *GetUsersForSectionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersForSectionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUsersForSectionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/sections/{id}/users"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetUsersForSectionURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var endingBeforeQ string
	if o.EndingBefore != nil {
		endingBeforeQ = *o.EndingBefore
	}
	if endingBeforeQ != "" {
		qs.Set("ending_before", endingBeforeQ)
	}

	var limitQ string
	if o.Limit != nil {
		limitQ = swag.FormatInt64(*o.Limit)
	}
	if limitQ != "" {
		qs.Set("limit", limitQ)
	}

	var primaryQ string
	if o.Primary != nil {
		primaryQ = *o.Primary
	}
	if primaryQ != "" {
		qs.Set("primary", primaryQ)
	}

	var roleQ string
	if o.Role != nil {
		roleQ = *o.Role
	}
	if roleQ != "" {
		qs.Set("role", roleQ)
	}

	var startingAfterQ string
	if o.StartingAfter != nil {
		startingAfterQ = *o.StartingAfter
	}
	if startingAfterQ != "" {
		qs.Set("starting_after", startingAfterQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUsersForSectionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUsersForSectionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUsersForSectionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUsersForSectionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUsersForSectionURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetUsersForSectionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
