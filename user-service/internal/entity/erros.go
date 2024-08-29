package entity

import (
	"errors"
	"fmt"
)

var (
	ErrorAlreadyExists    = fmt.Errorf("entity already exists")
	ErrorNotFound         = fmt.Errorf("entity not found")
	ErrorInvalidArguments = fmt.Errorf("invalid arguments")
	ErrInternal           = fmt.Errorf("internal error")
	ErrReqIsEmpty         = fmt.Errorf("req is empty")
	ErrUserNotRegistered  = fmt.Errorf("user not registered")
	ErrUserDeleted        = fmt.Errorf("user is deleted")
	ErrUserIsEmpty        = fmt.Errorf("user is empty")
)

var (
	ErrorConflict       = NewErrConflict("object")
	ErrorInvalidOTPCode = errors.New("code is invalid")
	ErrorOTPExpired     = errors.New("one time password has expired")
)

type ErrNotFound struct {
	name string
}

func (e *ErrNotFound) Error() string {
	return e.name + " not found"
}

func NewErrNotFound(text string) *ErrNotFound {
	return &ErrNotFound{text}
}

// error conflict
type ErrConflict struct {
	name string
}

func (e *ErrConflict) Error() string {
	return e.name + " already exist"
}

func NewErrConflict(text string) *ErrConflict {
	return &ErrConflict{text}
}

type ErrValidation struct {
	Err    error
	Errors map[string]string
}

func (e ErrValidation) Error() string {
	return e.Err.Error()
}

func NewErrValidation() *ErrValidation {
	return &ErrValidation{Errors: make(map[string]string)}
}

type ErrBadRequest struct {
	Err error
}

func (e ErrBadRequest) Error() string {
	return e.Err.Error()
}

func NewErrBadRequest(err error) *ErrBadRequest {
	return &ErrBadRequest{err}
}
