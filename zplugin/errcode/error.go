package errcode

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/pkg/errors"
)

type Error struct {
	mu   sync.Mutex
	err  error
	code int
}

func New() *Error {
	return &Error{}
}

func (e *Error) WithCode(c int) *Error {
	e.code = c
	return e
}

func (e *Error) WithError(err error) *Error {
	return e.WithMsg(e.Error())
}

func (e *Error) WithMsg(msg string) *Error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.err == nil {
		e.err = errors.New(msg)
	} else {
		e.err = errors.WithMessage(e.err, msg)
	}
	return e
}

func (e *Error) WithMsgf(format string, args ...string) *Error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.err == nil {
		e.err = errors.Errorf(format, args)
	} else {
		e.err = errors.WithMessagef(e.err, format, args)
	}
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %v, Msg: %v", e.code, e.err)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) String() string {
	return e.err.Error()
}

func (e *Error) HasError() bool {
	return e.code != 0 || e.err != nil
}

type interfaceStructure struct {
	pt uintptr
	pv uintptr
}

func IsNil(obj interface{}) bool {
	if obj == nil {
		return true
	}
	return (*interfaceStructure)(unsafe.Pointer(&obj)).pv == 0
}
