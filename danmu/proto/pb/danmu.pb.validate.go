// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: danmu.proto

package pb

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

// Validate checks the field values on Danmu with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Danmu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Danmu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DanmuMultiError, or nil if none found.
func (m *Danmu) ValidateAll() error {
	return m.validate(true)
}

func (m *Danmu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for VideoId

	// no validation rules for Content

	// no validation rules for BeginAt

	if len(errors) > 0 {
		return DanmuMultiError(errors)
	}

	return nil
}

// DanmuMultiError is an error wrapping multiple validation errors returned by
// Danmu.ValidateAll() if the designated constraints aren't met.
type DanmuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DanmuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DanmuMultiError) AllErrors() []error { return m }

// DanmuValidationError is the validation error returned by Danmu.Validate if
// the designated constraints aren't met.
type DanmuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DanmuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DanmuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DanmuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DanmuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DanmuValidationError) ErrorName() string { return "DanmuValidationError" }

// Error satisfies the builtin error interface
func (e DanmuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDanmu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DanmuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DanmuValidationError{}

// Validate checks the field values on SendReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendReq with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SendReqMultiError, or nil if none found.
func (m *SendReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SendReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VideoId

	// no validation rules for UserId

	// no validation rules for Content

	// no validation rules for BeginAt

	if len(errors) > 0 {
		return SendReqMultiError(errors)
	}

	return nil
}

// SendReqMultiError is an error wrapping multiple validation errors returned
// by SendReq.ValidateAll() if the designated constraints aren't met.
type SendReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendReqMultiError) AllErrors() []error { return m }

// SendReqValidationError is the validation error returned by SendReq.Validate
// if the designated constraints aren't met.
type SendReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendReqValidationError) ErrorName() string { return "SendReqValidationError" }

// Error satisfies the builtin error interface
func (e SendReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendReqValidationError{}

// Validate checks the field values on SendResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SendResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SendResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SendRespMultiError, or nil
// if none found.
func (m *SendResp) ValidateAll() error {
	return m.validate(true)
}

func (m *SendResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDanmu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SendRespValidationError{
					field:  "Danmu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SendRespValidationError{
					field:  "Danmu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDanmu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SendRespValidationError{
				field:  "Danmu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SendRespMultiError(errors)
	}

	return nil
}

// SendRespMultiError is an error wrapping multiple validation errors returned
// by SendResp.ValidateAll() if the designated constraints aren't met.
type SendRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SendRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SendRespMultiError) AllErrors() []error { return m }

// SendRespValidationError is the validation error returned by
// SendResp.Validate if the designated constraints aren't met.
type SendRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SendRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SendRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SendRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SendRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SendRespValidationError) ErrorName() string { return "SendRespValidationError" }

// Error satisfies the builtin error interface
func (e SendRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSendResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SendRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SendRespValidationError{}

// Validate checks the field values on GetDanmusReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetDanmusReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDanmusReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetDanmusReqMultiError, or
// nil if none found.
func (m *GetDanmusReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDanmusReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for VideoId

	if len(errors) > 0 {
		return GetDanmusReqMultiError(errors)
	}

	return nil
}

// GetDanmusReqMultiError is an error wrapping multiple validation errors
// returned by GetDanmusReq.ValidateAll() if the designated constraints aren't met.
type GetDanmusReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDanmusReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDanmusReqMultiError) AllErrors() []error { return m }

// GetDanmusReqValidationError is the validation error returned by
// GetDanmusReq.Validate if the designated constraints aren't met.
type GetDanmusReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDanmusReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDanmusReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDanmusReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDanmusReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDanmusReqValidationError) ErrorName() string { return "GetDanmusReqValidationError" }

// Error satisfies the builtin error interface
func (e GetDanmusReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDanmusReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDanmusReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDanmusReqValidationError{}

// Validate checks the field values on GetDanmusResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetDanmusResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDanmusResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetDanmusRespMultiError, or
// nil if none found.
func (m *GetDanmusResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDanmusResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDanmus() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetDanmusRespValidationError{
						field:  fmt.Sprintf("Danmus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetDanmusRespValidationError{
						field:  fmt.Sprintf("Danmus[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetDanmusRespValidationError{
					field:  fmt.Sprintf("Danmus[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetDanmusRespMultiError(errors)
	}

	return nil
}

// GetDanmusRespMultiError is an error wrapping multiple validation errors
// returned by GetDanmusResp.ValidateAll() if the designated constraints
// aren't met.
type GetDanmusRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDanmusRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDanmusRespMultiError) AllErrors() []error { return m }

// GetDanmusRespValidationError is the validation error returned by
// GetDanmusResp.Validate if the designated constraints aren't met.
type GetDanmusRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDanmusRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDanmusRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDanmusRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDanmusRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDanmusRespValidationError) ErrorName() string { return "GetDanmusRespValidationError" }

// Error satisfies the builtin error interface
func (e GetDanmusRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDanmusResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDanmusRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDanmusRespValidationError{}
