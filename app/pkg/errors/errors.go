package errors

import (
	"errors"
	"fmt"
)

const (
	Canceled           = "canceled"
	Unknown            = "unknown_error"
	DeadlineExceeded   = "deadline_exceeded"
	NotFound           = "not_found"
	AlreadyExists      = "already_exists"
	PermissionDenied   = "permission_denied"
	ResourceExhausted  = "resource_exhausted"
	FailedPrecondition = "failed_precondition"
	Aborted            = "aborted"
	OutOfRange         = "out_of_range"
	Unimplemented      = "unimplemented"
	Internal           = "internal"
	Unavailable        = "unavailable"
	DataLoss           = "data_loss"
	Unauthenticated    = "unauthenticated"
	InvalidArgument    = "invalid_argument"
	InvalidID          = "invalid_id"
	InvalidName        = "invalid_name"
)

// Error represents an application-specific error. Application errors can be
// unwrapped by the caller to extract out the code & message.
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable error message.
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("error: code=%s message=%s", e.Code, e.Message)
}

// ErrorCode unwraps an application error and returns its code.
// Non-application errors always return Internal.
func ErrorCode(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}

	return Internal
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "internal".
func ErrorMessage(err error) string {
	var e *Error

	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}

	return Internal
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}
