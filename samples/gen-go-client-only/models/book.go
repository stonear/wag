// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Book book
//
// swagger:model Book
type Book struct {

	// author
	Author string `json:"author,omitempty"`

	// genre
	// Enum: [scifi mystery horror]
	Genre string `json:"genre,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// other
	Other map[string]string `json:"other,omitempty"`

	// other array
	OtherArray map[string][]string `json:"otherArray,omitempty"`
}

// Validate validates this book
func (m *Book) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenre(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bookTypeGenrePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scifi","mystery","horror"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bookTypeGenrePropEnum = append(bookTypeGenrePropEnum, v)
	}
}

const (

	// BookGenreScifi captures enum value "scifi"
	BookGenreScifi string = "scifi"

	// BookGenreMystery captures enum value "mystery"
	BookGenreMystery string = "mystery"

	// BookGenreHorror captures enum value "horror"
	BookGenreHorror string = "horror"
)

// prop value enum
func (m *Book) validateGenreEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bookTypeGenrePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Book) validateGenre(formats strfmt.Registry) error {

	if swag.IsZero(m.Genre) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenreEnum("genre", "body", m.Genre); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Book) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Book) UnmarshalBinary(b []byte) error {
	var res Book
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
