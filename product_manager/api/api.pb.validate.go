// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: product_manager/api/api.proto

package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ShelfQuantity with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ShelfQuantity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ShelfQuantity with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ShelfQuantityMultiError, or
// nil if none found.
func (m *ShelfQuantity) ValidateAll() error {
	return m.validate(true)
}

func (m *ShelfQuantity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetShelfName()) < 1 {
		err := ShelfQuantityValidationError{
			field:  "ShelfName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Quantity

	if len(errors) > 0 {
		return ShelfQuantityMultiError(errors)
	}

	return nil
}

// ShelfQuantityMultiError is an error wrapping multiple validation errors
// returned by ShelfQuantity.ValidateAll() if the designated constraints
// aren't met.
type ShelfQuantityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ShelfQuantityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ShelfQuantityMultiError) AllErrors() []error { return m }

// ShelfQuantityValidationError is the validation error returned by
// ShelfQuantity.Validate if the designated constraints aren't met.
type ShelfQuantityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShelfQuantityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShelfQuantityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShelfQuantityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShelfQuantityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShelfQuantityValidationError) ErrorName() string { return "ShelfQuantityValidationError" }

// Error satisfies the builtin error interface
func (e ShelfQuantityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShelfQuantity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShelfQuantityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShelfQuantityValidationError{}

// Validate checks the field values on ImportRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ImportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ImportRequestMultiError, or
// nil if none found.
func (m *ImportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSku()) < 1 {
		err := ImportRequestValidationError{
			field:  "Sku",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := ImportRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetExpiredDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImportRequestValidationError{
					field:  "ExpiredDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImportRequestValidationError{
					field:  "ExpiredDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiredDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImportRequestValidationError{
				field:  "ExpiredDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Category

	if len(m.GetShelfQuantities()) < 1 {
		err := ImportRequestValidationError{
			field:  "ShelfQuantities",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetShelfQuantities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ImportRequestValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ImportRequestValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImportRequestValidationError{
					field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ImportRequestMultiError(errors)
	}

	return nil
}

// ImportRequestMultiError is an error wrapping multiple validation errors
// returned by ImportRequest.ValidateAll() if the designated constraints
// aren't met.
type ImportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportRequestMultiError) AllErrors() []error { return m }

// ImportRequestValidationError is the validation error returned by
// ImportRequest.Validate if the designated constraints aren't met.
type ImportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportRequestValidationError) ErrorName() string { return "ImportRequestValidationError" }

// Error satisfies the builtin error interface
func (e ImportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportRequestValidationError{}

// Validate checks the field values on ExportRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExportRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExportRequestMultiError, or
// nil if none found.
func (m *ExportRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSku()) < 1 {
		err := ExportRequestValidationError{
			field:  "Sku",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetShelfQuantities()) < 1 {
		err := ExportRequestValidationError{
			field:  "ShelfQuantities",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetShelfQuantities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExportRequestValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExportRequestValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExportRequestValidationError{
					field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ExportRequestMultiError(errors)
	}

	return nil
}

// ExportRequestMultiError is an error wrapping multiple validation errors
// returned by ExportRequest.ValidateAll() if the designated constraints
// aren't met.
type ExportRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportRequestMultiError) AllErrors() []error { return m }

// ExportRequestValidationError is the validation error returned by
// ExportRequest.Validate if the designated constraints aren't met.
type ExportRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportRequestValidationError) ErrorName() string { return "ExportRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExportRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportRequestValidationError{}

// Validate checks the field values on ExportResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExportResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExportResponseMultiError,
// or nil if none found.
func (m *ExportResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetShelfQuantities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExportResponseValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExportResponseValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExportResponseValidationError{
					field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ExportResponseMultiError(errors)
	}

	return nil
}

// ExportResponseMultiError is an error wrapping multiple validation errors
// returned by ExportResponse.ValidateAll() if the designated constraints
// aren't met.
type ExportResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportResponseMultiError) AllErrors() []error { return m }

// ExportResponseValidationError is the validation error returned by
// ExportResponse.Validate if the designated constraints aren't met.
type ExportResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportResponseValidationError) ErrorName() string { return "ExportResponseValidationError" }

// Error satisfies the builtin error interface
func (e ExportResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportResponseValidationError{}

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductRequestMultiError, or nil if none found.
func (m *GetProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetSku()) < 1 {
		err := GetProductRequestValidationError{
			field:  "Sku",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetProductRequestMultiError(errors)
	}

	return nil
}

// GetProductRequestMultiError is an error wrapping multiple validation errors
// returned by GetProductRequest.ValidateAll() if the designated constraints
// aren't met.
type GetProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductRequestMultiError) AllErrors() []error { return m }

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on GetProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductResponseMultiError, or nil if none found.
func (m *GetProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	// no validation rules for Name

	for idx, item := range m.GetShelfQuantities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetProductResponseValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetProductResponseValidationError{
						field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetProductResponseValidationError{
					field:  fmt.Sprintf("ShelfQuantities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetExpiredDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProductResponseValidationError{
					field:  "ExpiredDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProductResponseValidationError{
					field:  "ExpiredDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiredDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProductResponseValidationError{
				field:  "ExpiredDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Category

	if len(errors) > 0 {
		return GetProductResponseMultiError(errors)
	}

	return nil
}

// GetProductResponseMultiError is an error wrapping multiple validation errors
// returned by GetProductResponse.ValidateAll() if the designated constraints
// aren't met.
type GetProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductResponseMultiError) AllErrors() []error { return m }

// GetProductResponseValidationError is the validation error returned by
// GetProductResponse.Validate if the designated constraints aren't met.
type GetProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductResponseValidationError) ErrorName() string {
	return "GetProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductResponseValidationError{}
