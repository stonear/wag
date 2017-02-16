package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// AuthorsResponse authors response
// swagger:model AuthorsResponse
type AuthorsResponse struct {

	// author set
	AuthorSet *AuthorSet `json:"authorSet,omitempty"`

	// metadata
	Metadata *AuthorsResponseMetadata `json:"metadata,omitempty"`
}

// Validate validates this authors response
func (m *AuthorsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorSet(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorsResponse) validateAuthorSet(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorSet) { // not required
		return nil
	}

	if m.AuthorSet != nil {

		if err := m.AuthorSet.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *AuthorsResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {

		if err := m.Metadata.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
