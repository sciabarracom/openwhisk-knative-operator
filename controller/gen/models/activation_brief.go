// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActivationBrief activation brief
// swagger:model ActivationBrief
type ActivationBrief struct {

	// Id of the activation
	// Required: true
	ActivationID *string `json:"activationId"`

	// annotations on the item
	Annotations []*KeyValue `json:"annotations"`

	// the activation id that caused this activation
	Cause string `json:"cause,omitempty"`

	// How long the invocation took, in millisecnods
	Duration int64 `json:"duration,omitempty"`

	// Time when the activation completed
	End int64 `json:"end,omitempty"`

	// Name of the item
	// Required: true
	Name *string `json:"name"`

	// Namespace of the associated item
	// Required: true
	Namespace *string `json:"namespace"`

	// Whether to publish the item or not
	// Required: true
	Publish *bool `json:"publish"`

	// Time when the activation began
	// Required: true
	Start *int64 `json:"start"`

	// The status code
	// Enum: [0 1 2]
	StatusCode int32 `json:"statusCode,omitempty"`

	// Semantic version of the item
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this activation brief
func (m *ActivationBrief) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAnnotations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublish(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivationBrief) validateActivationID(formats strfmt.Registry) error {

	if err := validate.Required("activationId", "body", m.ActivationID); err != nil {
		return err
	}

	return nil
}

func (m *ActivationBrief) validateAnnotations(formats strfmt.Registry) error {

	if swag.IsZero(m.Annotations) { // not required
		return nil
	}

	for i := 0; i < len(m.Annotations); i++ {
		if swag.IsZero(m.Annotations[i]) { // not required
			continue
		}

		if m.Annotations[i] != nil {
			if err := m.Annotations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("annotations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ActivationBrief) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ActivationBrief) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *ActivationBrief) validatePublish(formats strfmt.Registry) error {

	if err := validate.Required("publish", "body", m.Publish); err != nil {
		return err
	}

	return nil
}

func (m *ActivationBrief) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", m.Start); err != nil {
		return err
	}

	return nil
}

var activationBriefTypeStatusCodePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[0,1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activationBriefTypeStatusCodePropEnum = append(activationBriefTypeStatusCodePropEnum, v)
	}
}

// prop value enum
func (m *ActivationBrief) validateStatusCodeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, activationBriefTypeStatusCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ActivationBrief) validateStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusCodeEnum("statusCode", "body", m.StatusCode); err != nil {
		return err
	}

	return nil
}

func (m *ActivationBrief) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivationBrief) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivationBrief) UnmarshalBinary(b []byte) error {
	var res ActivationBrief
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
