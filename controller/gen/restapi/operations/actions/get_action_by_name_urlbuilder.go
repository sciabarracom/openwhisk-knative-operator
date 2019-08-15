// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetActionByNameURL generates an URL for the get action by name operation
type GetActionByNameURL struct {
	ActionName string
	Namespace  string

	Code *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetActionByNameURL) WithBasePath(bp string) *GetActionByNameURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetActionByNameURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetActionByNameURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/namespaces/{namespace}/actions/{actionName}"

	actionName := o.ActionName
	if actionName != "" {
		_path = strings.Replace(_path, "{actionName}", actionName, -1)
	} else {
		return nil, errors.New("actionName is required on GetActionByNameURL")
	}

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on GetActionByNameURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var code string
	if o.Code != nil {
		code = swag.FormatBool(*o.Code)
	}
	if code != "" {
		qs.Set("code", code)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetActionByNameURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetActionByNameURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetActionByNameURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetActionByNameURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetActionByNameURL")
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
func (o *GetActionByNameURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}