// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetContactsForUserURL generates an URL for the get contacts for user operation
type GetContactsForUserURL struct {
	ID string

	EndingBefore  *string
	Limit         *int64
	StartingAfter *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetContactsForUserURL) WithBasePath(bp string) *GetContactsForUserURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetContactsForUserURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetContactsForUserURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/users/{id}/mycontacts"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on GetContactsForUserURL")
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
func (o *GetContactsForUserURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetContactsForUserURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetContactsForUserURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetContactsForUserURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetContactsForUserURL")
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
func (o *GetContactsForUserURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
