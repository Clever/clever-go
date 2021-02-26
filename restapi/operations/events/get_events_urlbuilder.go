// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetEventsURL generates an URL for the get events operation
type GetEventsURL struct {
	EndingBefore  *string
	Limit         *int64
	RecordType    []string
	School        *string
	StartingAfter *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEventsURL) WithBasePath(bp string) *GetEventsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetEventsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetEventsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/events"

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

	var recordTypeIR []string
	for _, recordTypeI := range o.RecordType {
		recordTypeIS := recordTypeI
		if recordTypeIS != "" {
			recordTypeIR = append(recordTypeIR, recordTypeIS)
		}
	}

	recordType := swag.JoinByFormat(recordTypeIR, "multi")

	for _, qsv := range recordType {
		qs.Add("record_type", qsv)
	}

	var schoolQ string
	if o.School != nil {
		schoolQ = *o.School
	}
	if schoolQ != "" {
		qs.Set("school", schoolQ)
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
func (o *GetEventsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetEventsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetEventsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetEventsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetEventsURL")
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
func (o *GetEventsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
