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

func Int(msgAndMock ...interface{}) int {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int)
	}

	return 0
}

func IntF(msgAndMock ...interface{}) func(...any) int {
	return func(_ ...any) int {
		return Int(msgAndMock...)
	}
}

func Int8(msgAndMock ...interface{}) int8 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int8)
	}

	return 0
}

func Int8F(msgAndMock ...interface{}) func(...any) int8 {
	return func(_ ...any) int8 {
		return Int8(msgAndMock...)
	}
}

func Int32(msgAndMock ...interface{}) int32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int32)
	}

	return 0
}

func Int32F(msgAndMock ...interface{}) func(...any) int32 {
	return func(_ ...any) int32 {
		return Int32(msgAndMock...)
	}
}

func UInt32(msgAndMock ...interface{}) uint32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(uint32)
	}

	return 0
}

func UInt32F(msgAndMock ...interface{}) func(...any) uint32 {
	return func(_ ...any) uint32 {
		return UInt32(msgAndMock...)
	}
}

func Int64(msgAndMock ...interface{}) int64 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int64)
	}

	return 0
}

func Int64F(msgAndMock ...interface{}) func(...any) int64 {
	return func(_ ...any) int64 {
		return Int64(msgAndMock...)
	}
}

func Bool(msgAndMock ...interface{}) bool {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(bool)
	}

	return false
}

func BoolF(msgAndMock ...interface{}) func(...any) bool {
	return func(_ ...any) bool {
		return Bool(msgAndMock...)
	}
}

func Float32(msgAndMock ...interface{}) float32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(float32)
	}

	return 0.0
}

func Float32F(msgAndMock ...interface{}) func(...any) float32 {
	return func(_ ...any) float32 {
		return Float32(msgAndMock...)
	}
}

func Float64(msgAndMock ...interface{}) float64 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(float64)
	}

	return 0.0
}

func Float64F(msgAndMock ...interface{}) func(...any) float64 {
	return func(_ ...any) float64 {
		return Float64(msgAndMock...)
	}
}
