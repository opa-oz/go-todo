package todo

import (
	"errors"
	"fmt"
	"log"
)

var TodoError = errors.New("This is TODO error")

type Todoer struct {
	prefix string
	logger *log.Logger
}

func (t *Todoer) ChangePrefix(newPrefix string) {
	t.prefix = newPrefix
}

func (t *Todoer) AttachLogger(logger *log.Logger) {
	t.logger = logger
}

var Opts = Todoer{prefix: "[TODO]: "}

func (t *Todoer) FancyPrint(message string) {
	msg := fmt.Sprintf("%s\"%s\"", t.prefix, message)

	if t.logger != nil {
		t.logger.Println(msg)
	} else {
		fmt.Println()
	}
}

func Nil(message ...string) any {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}
	return nil
}

func NilF(message ...string) func(...any) any {
	return func(_ ...any) any {
		return Nil(message...)
	}
}

func Error(msgAndMock ...string) error {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return errors.New(msgAndMock[1])
	}

	return TodoError
}

func ErrorF(msgAndMock ...string) func(...any) error {
	return func(_ ...any) error {
		return Error(msgAndMock...)
	}
}

func String(msgAndMock ...string) string {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1]
	}

	return ""
}

func StringF(msgAndMock ...string) func(...any) string {
	return func(_ ...any) string {
		return String(msgAndMock...)
	}
}

func Ptr[V any](message ...string) *V {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}

	return nil
}

func PtrF[V any](message ...string) func(...any) *V {
	return func(_ ...any) *V {
		return Ptr[V](message...)
	}
}

func T(message ...string) {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}
}
