// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostTestParamsBody post test params body
// swagger:model postTestParamsBody
type PostTestParamsBody struct {

	// bam
	Bam string `json:"bam,omitempty"`

	// bar
	Bar int64 `json:"bar,omitempty"`

	// foo
	Foo bool `json:"foo,omitempty"`
}

func (m *PostTestParamsBody) UnmarshalJSON(b []byte) error {
	type PostTestParamsBodyAlias PostTestParamsBody
	var t PostTestParamsBodyAlias
	if err := json.Unmarshal([]byte("{\"bam\":\"boom\",\"bar\":42,\"foo\":true}"), &t); err != nil {
		return err
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	*m = PostTestParamsBody(t)
	return nil
}

// Validate validates this post test params body
func (m *PostTestParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostTestParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTestParamsBody) UnmarshalBinary(b []byte) error {
	var res PostTestParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
