package errcode

import (
	"errors"
	"fmt"
	"sync"
)

type Error struct {
	sync.Mutex
	errs []error
	code int
}

func New() *Error {
	return &Error{}
}

func (e *Error) WithCode(c int) *Error {
	z.WithLocker(e, func(){e.code=c})
	return e
}

func (e *Error) WithError(errs ...error) *Error {
	e.errs = append(e.errs, errs...)
	return e
}

func (e *Error) WithMsg(msgs ...string) *Error {
	for _, msg := range msgs {
	e.errs = append(e.errs, errors.New(msg))
	}
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error: %s code: %d", e.errmsg, e.code)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) String() string {
	return e.errmsg
}

func (e *Error) HasError() string {
	return e.errmsg
}

// type eFace struct {
// 	rType unsafe.Pointer
// 	data  unsafe.Pointer
// }

// func (e *Error) Append(msgs ...interface{}) *Error {
// 	err := new(Error)
// 	err.errmsg = fmt.Sprint(e.errmsg, " ", msgs)
// 	err.code = e.code
// 	return err
// }
// func IsNil(obj interface{}) bool {
// 	return isNil(obj)
// }
// func isNil(obj interface{}) bool {
// 	if obj == nil {
// 		return true
// 	}
// 	return (*eFace)(unsafe.Pointer(&obj)).data == nil
// }
// func CheckErr(err *Error) {
// 	if !isNil(err) {
// 		panic(err)
// 	}
// }
