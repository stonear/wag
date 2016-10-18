package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Book book
// swagger:model Book
type Book struct {

	// author
	Author string `json:"author,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// other
	Other map[string]string `json:"other,omitempty"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOther(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Book) validateOther(formats strfmt.Registry) error {

	if swag.IsZero(m.Other) { // not required
		return nil
	}

	if err := validate.Required("other", "body", m.Other); err != nil {
		return err
	}

	return nil
}
