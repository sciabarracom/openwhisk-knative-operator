// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PutWebNamespacePackageNameActionNameExtensionURL generates an URL for the put web namespace package name action name extension operation
type PutWebNamespacePackageNameActionNameExtensionURL struct {
	ActionName  string
	Extension   string
	Namespace   string
	PackageName string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutWebNamespacePackageNameActionNameExtensionURL) WithBasePath(bp string) *PutWebNamespacePackageNameActionNameExtensionURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PutWebNamespacePackageNameActionNameExtensionURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PutWebNamespacePackageNameActionNameExtensionURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/web/{namespace}/{packageName}/{actionName}.{extension}"

	actionName := o.ActionName
	if actionName != "" {
		_path = strings.Replace(_path, "{actionName}", actionName, -1)
	} else {
		return nil, errors.New("actionName is required on PutWebNamespacePackageNameActionNameExtensionURL")
	}

	extension := o.Extension
	if extension != "" {
		_path = strings.Replace(_path, "{extension}", extension, -1)
	} else {
		return nil, errors.New("extension is required on PutWebNamespacePackageNameActionNameExtensionURL")
	}

	namespace := o.Namespace
	if namespace != "" {
		_path = strings.Replace(_path, "{namespace}", namespace, -1)
	} else {
		return nil, errors.New("namespace is required on PutWebNamespacePackageNameActionNameExtensionURL")
	}

	packageName := o.PackageName
	if packageName != "" {
		_path = strings.Replace(_path, "{packageName}", packageName, -1)
	} else {
		return nil, errors.New("packageName is required on PutWebNamespacePackageNameActionNameExtensionURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PutWebNamespacePackageNameActionNameExtensionURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PutWebNamespacePackageNameActionNameExtensionURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PutWebNamespacePackageNameActionNameExtensionURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PutWebNamespacePackageNameActionNameExtensionURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PutWebNamespacePackageNameActionNameExtensionURL")
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
func (o *PutWebNamespacePackageNameActionNameExtensionURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
