// Code generated by protoc-gen-validate
// source: communication.proto
// DO NOT EDIT!!!

package communication

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on IP with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *IP) Validate() error {
	if m == nil {
		return nil
	}

	if ip := net.ParseIP(m.GetAddr()); ip == nil {
		return IPValidationError{
			Field:  "Addr",
			Reason: "value must be a valid IP address",
		}
	}

	if err := m._validateHostname(m.GetName()); err != nil {
		return IPValidationError{
			Field:  "Name",
			Reason: "value must be a valid hostname",
			Cause:  err,
		}
	}

	if val := m.GetPort(); val < 1024 || val > 65535 {
		return IPValidationError{
			Field:  "Port",
			Reason: "value must be inside range [1024, 65535]",
		}
	}

	return nil
}

func (m *IP) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// IPValidationError is the validation error returned by IP.Validate if the
// designated constraints aren't met.
type IPValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e IPValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIP.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = IPValidationError{}

// Validate checks the field values on GreetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GreetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDoublenum().Valid {

		if m.GetDoublenum().Valid {

		}

		if m.GetDoublenum().Valid {

			if m.GetDoublenum().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Doublenum",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetDoublenum().Valid {

		}

	}

	if m.GetFloatnum().Valid {

		if m.GetFloatnum().Valid {

		}

		if m.GetFloatnum().Valid {

			if m.GetFloatnum().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Floatnum",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetFloatnum().Valid {

		}

	}

	if m.GetInt64Num().Valid {

		if m.GetInt64Num().Valid {

		}

		if m.GetInt64Num().Valid {

			if m.GetInt64Num().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Int64Num",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetInt64Num().Valid {

		}

	}

	if m.GetUint64Num().Valid {

		if m.GetUint64Num().Valid {

		}

		if m.GetUint64Num().Valid {

			if m.GetUint64Num().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Uint64Num",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetUint64Num().Valid {

		}

	}

	if m.GetInt32Num().Valid {

		if m.GetInt32Num().Valid {

		}

		if m.GetInt32Num().Valid {

			if m.GetInt32Num().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Int32Num",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetInt32Num().Valid {

		}

	}

	if m.GetUint32Num().Valid {

		if m.GetUint32Num().Valid {

		}

		if m.GetUint32Num().Valid {

			if m.GetUint32Num().Value <= 10 {
				return GreetRequestValidationError{
					Field:  "Uint32Num",
					Reason: "value must be greater than 10",
				}
			}

		}

		if m.GetUint32Num().Valid {

		}

	}

	if m.GetStr().Valid {

		if m.GetStr().Valid {

		}

		if m.GetStr().Valid {

		}

		if utf8.RuneCountInString(m.GetStr().Value) < 2 {
			return GreetRequestValidationError{
				Field:  "Str",
				Reason: "value length must be at least 2 runes",
			}
		}

	}

	if m.GetBo().Valid {

		if m.GetBo().Value != true {
			return GreetRequestValidationError{
				Field:  "Bo",
				Reason: "value must equal true",
			}
		}

	}

	if m.GetB().Valid {

		if !bytes.HasPrefix(m.GetB().Value, []uint8{0x68, 0x61, 0x68, 0x61}) {
			return GreetRequestValidationError{
				Field:  "B",
				Reason: "value does not have prefix \"\\x68\\x61\\x68\\x61\"",
			}
		}

	}

	return nil
}

// GreetRequestValidationError is the validation error returned by
// GreetRequest.Validate if the designated constraints aren't met.
type GreetRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GreetRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGreetRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GreetRequestValidationError{}

// Validate checks the field values on GreetReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GreetReply) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetMessage()); l < 0 || l > 65535 {
		return GreetReplyValidationError{
			Field:  "Message",
			Reason: "value length must be between 0 and 65535 runes, inclusive",
		}
	}

	if m.GetIp() == nil {
		return GreetReplyValidationError{
			Field:  "Ip",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GreetReplyValidationError{
				Field:  "Ip",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// GreetReplyValidationError is the validation error returned by
// GreetReply.Validate if the designated constraints aren't met.
type GreetReplyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e GreetReplyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGreetReply.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = GreetReplyValidationError{}

// Validate checks the field values on AccessRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccessRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIp() == nil {
		return AccessRequestValidationError{
			Field:  "Ip",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessRequestValidationError{
				Field:  "Ip",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for IsCN

	if m.GetEmail() == nil {
		return AccessRequestValidationError{
			Field:  "Email",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetEmail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessRequestValidationError{
				Field:  "Email",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AccessRequestValidationError is the validation error returned by
// AccessRequest.Validate if the designated constraints aren't met.
type AccessRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessRequestValidationError{}

// Validate checks the field values on AccessReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccessReply) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetIp() == nil {
		return AccessReplyValidationError{
			Field:  "Ip",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetIp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccessReplyValidationError{
				Field:  "Ip",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if l := utf8.RuneCountInString(m.GetMessage()); l < 0 || l > 65535 {
		return AccessReplyValidationError{
			Field:  "Message",
			Reason: "value length must be between 0 and 65535 runes, inclusive",
		}
	}

	return nil
}

// AccessReplyValidationError is the validation error returned by
// AccessReply.Validate if the designated constraints aren't met.
type AccessReplyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessReplyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessReply.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessReplyValidationError{}

// Validate checks the field values on AccessRequest_Email with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AccessRequest_Email) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return AccessRequest_EmailValidationError{
			Field:  "Email",
			Reason: "value must be a valid email address",
			Cause:  err,
		}
	}

	if l := utf8.RuneCountInString(m.GetPwd()); l < 0 || l > 16 {
		return AccessRequest_EmailValidationError{
			Field:  "Pwd",
			Reason: "value length must be between 0 and 16 runes, inclusive",
		}
	}

	return nil
}

func (m *AccessRequest_Email) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *AccessRequest_Email) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// AccessRequest_EmailValidationError is the validation error returned by
// AccessRequest_Email.Validate if the designated constraints aren't met.
type AccessRequest_EmailValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AccessRequest_EmailValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccessRequest_Email.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AccessRequest_EmailValidationError{}
