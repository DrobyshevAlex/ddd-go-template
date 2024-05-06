package domainErrors

import "errors"

type UnauthenticatedError struct {
	msg string
}

func (e UnauthenticatedError) Error() string {
	return e.msg
}

func Unauthenticated(msg string) error {
	return UnauthenticatedError{msg: msg}
}

type NotAllowedError struct {
	msg string
}

func (e NotAllowedError) Error() string {
	return e.msg
}

func NotAllowed(msg string) error {
	return NotAllowedError{msg: msg}
}

type NotFoundError struct {
	msg string
}

func (e NotFoundError) Error() string {
	return e.msg
}

func NotFound(msg string) error {
	return NotFoundError{msg: msg}
}

func IsNotFound(err error) bool {
	var t NotFoundError
	return errors.As(err, &t)
}

type AlreadyExistsError struct {
	msg string
}

func (e AlreadyExistsError) Error() string {
	return e.msg
}

func AlreadyExists(msg string) error {
	return AlreadyExistsError{msg: msg}
}

type ValidationError struct {
	msg string
}

func (e ValidationError) Error() string {
	return e.msg
}

func Validation(msg string) error {
	return ValidationError{msg: msg}
}

type InternalError struct {
	msg string
}

func (e InternalError) Error() string {
	return e.msg
}

func Internal(msg string) error {
	return InternalError{msg: msg}
}
