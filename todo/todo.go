package todo

import (
	"errors"
	"fmt"
)

var TodoError = errors.New("This is TODO error")

type Todoer struct {
	prefix string
}

func (t *Todoer) ChangePrefix(newPrefix string) {
	t.prefix = newPrefix
}

var Opts = Todoer{prefix: "[TODO]"}

func fancyPrint(message string) {
	fmt.Println(fmt.Sprintf("%s: \"%s\"", Opts.prefix, message))
}

func Nil[T any](message ...string) *T {
	if len(message) > 0 {
		fancyPrint(message[0])
	}
	return nil
}

func NilF[T any](message string) func(...interface{}) *T {
	return func(_ ...interface{}) *T {
		return Nil[T](message)
	}
}

func Error(msgAndMock ...string) error {
	if len(msgAndMock) > 0 {
		fancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return errors.New(msgAndMock[1])
	}

	return TodoError
}

func ErrorF(msgAndMock ...string) func(...interface{}) error {
	return func(_ ...interface{}) error {
		return Error(msgAndMock...)
	}
}

func String(msgAndMock ...string) string {
	if len(msgAndMock) > 0 {
		fancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1]
	}

	return ""
}

func StringF(msgAndMock ...string) func(...interface{}) string {
	return func(_ ...interface{}) string {
		return String(msgAndMock...)
	}
}

func Ptr[V any](message ...string) *V {
	if len(message) > 0 {
		fancyPrint(message[0])
	}

	return nil
}

func PtrF[V any](message ...string) func(...interface{}) *V {
	return func(_ ...interface{}) *V {
		return Ptr[V](message...)
	}
}
