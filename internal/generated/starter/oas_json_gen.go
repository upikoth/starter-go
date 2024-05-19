// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *DefaultErrorResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultErrorResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("success")
		s.Success.Encode(e)
	}
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
	{
		e.FieldStart("error")
		s.Error.Encode(e)
	}
}

var jsonFieldsNameOfDefaultErrorResponse = [3]string{
	0: "success",
	1: "data",
	2: "error",
}

// Decode decodes DefaultErrorResponse from json.
func (s *DefaultErrorResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultErrorResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "success":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Success.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"success\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Data.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "error":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.Error.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"error\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DefaultErrorResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDefaultErrorResponse) {
					name = jsonFieldsNameOfDefaultErrorResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultErrorResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultErrorResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DefaultErrorResponseData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultErrorResponseData) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfDefaultErrorResponseData = [0]string{}

// Decode decodes DefaultErrorResponseData from json.
func (s *DefaultErrorResponseData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultErrorResponseData to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode DefaultErrorResponseData")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultErrorResponseData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultErrorResponseData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DefaultErrorResponseError) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultErrorResponseError) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("code")
		e.Str(s.Code)
	}
	{
		e.FieldStart("description")
		e.Str(s.Description)
	}
}

var jsonFieldsNameOfDefaultErrorResponseError = [2]string{
	0: "code",
	1: "description",
}

// Decode decodes DefaultErrorResponseError from json.
func (s *DefaultErrorResponseError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultErrorResponseError to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "code":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Code = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		case "description":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Description = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"description\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DefaultErrorResponseError")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDefaultErrorResponseError) {
					name = jsonFieldsNameOfDefaultErrorResponseError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultErrorResponseError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultErrorResponseError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DefaultErrorResponseSuccess as json.
func (s DefaultErrorResponseSuccess) Encode(e *jx.Encoder) {
	e.Bool(bool(s))
}

// Decode decodes DefaultErrorResponseSuccess from json.
func (s *DefaultErrorResponseSuccess) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultErrorResponseSuccess to nil")
	}
	v, err := d.Bool()
	if err != nil {
		return err
	}
	*s = DefaultErrorResponseSuccess(v)

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s DefaultErrorResponseSuccess) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultErrorResponseSuccess) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DefaultSuccessResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultSuccessResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("success")
		s.Success.Encode(e)
	}
	{
		e.FieldStart("data")
		s.Data.Encode(e)
	}
}

var jsonFieldsNameOfDefaultSuccessResponse = [2]string{
	0: "success",
	1: "data",
}

// Decode decodes DefaultSuccessResponse from json.
func (s *DefaultSuccessResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultSuccessResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "success":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Success.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"success\"")
			}
		case "data":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Data.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode DefaultSuccessResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfDefaultSuccessResponse) {
					name = jsonFieldsNameOfDefaultSuccessResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultSuccessResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultSuccessResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *DefaultSuccessResponseData) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *DefaultSuccessResponseData) encodeFields(e *jx.Encoder) {
}

var jsonFieldsNameOfDefaultSuccessResponseData = [0]string{}

// Decode decodes DefaultSuccessResponseData from json.
func (s *DefaultSuccessResponseData) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultSuccessResponseData to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		default:
			return d.Skip()
		}
	}); err != nil {
		return errors.Wrap(err, "decode DefaultSuccessResponseData")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *DefaultSuccessResponseData) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultSuccessResponseData) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes DefaultSuccessResponseSuccess as json.
func (s DefaultSuccessResponseSuccess) Encode(e *jx.Encoder) {
	e.Bool(bool(s))
}

// Decode decodes DefaultSuccessResponseSuccess from json.
func (s *DefaultSuccessResponseSuccess) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode DefaultSuccessResponseSuccess to nil")
	}
	v, err := d.Bool()
	if err != nil {
		return err
	}
	*s = DefaultSuccessResponseSuccess(v)

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s DefaultSuccessResponseSuccess) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *DefaultSuccessResponseSuccess) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
