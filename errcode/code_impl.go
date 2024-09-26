package errcode

import "github.com/pkg/errors"

//go:generate stringer -type Code -linecomment
type Code int

const (
	CodeGeneralErr Code = -1           // unknow error
	CodeRepeatInit Code = 10000 + iota // repeat init exception
)

func (code Code) Code() int { return int(code) }

func (code Code) Error() *Error {
	return &Error{code: code.Code(), desc: code.String()}
}

func (code Code) WithError(err error) *Error {
	return &Error{code: code.Code(), desc: code.String()}
}

func (code Code) WithMessage(msg string) *Error {
	return &Error{code: code.Code(), desc: code.String(), cause: errors.New(msg)}
}

func (code Code) WithMessagef(format string, args ...interface{}) *Error {
	return &Error{code: code.Code(), desc: code.String(), cause: errors.Errorf(format, args...)}
}
