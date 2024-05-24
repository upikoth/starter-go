// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s *ErrorResponse) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Success.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "success",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ErrorResponseStatusCode) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Response.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "Response",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ErrorResponseSuccess) Validate() error {
	switch s {
	case false:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *SuccessResponse) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Success.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "success",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s SuccessResponseSuccess) Validate() error {
	switch s {
	case true:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *V1RegistrationsCreateRegistrationRequestBody) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        true,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Email)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *V1RegistrationsCreateRegistrationResponse) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Success.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "success",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Data.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *V1RegistrationsCreateRegistrationResponseData) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    0,
			MaxLengthSet: false,
			Email:        true,
			Hostname:     false,
			Regex:        nil,
		}).Validate(string(s.Email)); err != nil {
			return errors.Wrap(err, "string")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s V1RegistrationsCreateRegistrationResponseSuccess) Validate() error {
	switch s {
	case true:
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
