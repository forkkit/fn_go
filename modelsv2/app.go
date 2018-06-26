// Code generated by go-swagger; DO NOT EDIT.

package modelsv2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// App app
// swagger:model App
type App struct {

	// Application annotations - this is a map of annotations attached to this app, keys must not exceed 128 bytes and must consist of non-whitespace printable ascii characters, and the seralized representation of individual values must not exeed 512 bytes.
	Annotations map[string]interface{} `json:"annotations,omitempty"`

	// Application function configuration, applied to all routes.
	Config map[string]string `json:"config,omitempty"`

	// Time when app was created. Always in UTC.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// App ID
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Name of this app. Must be different than the image name. Can ony contain alphanumeric, -, and _.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// A comma separated list of syslog urls to send all function logs to. supports tls, udp or tcp. e.g. tls://logs.papertrailapp.com:1
	SyslogURL string `json:"syslog_url,omitempty"`

	// Most recent time that app was updated. Always in UTC.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this app
func (m *App) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *App) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *App) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *App) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *App) UnmarshalBinary(b []byte) error {
	var res App
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
