package todo

import (
	"errors"
	"fmt"
	"log"
)

var TodoError = errors.New("This is TODO error")

// Todoer represents a type for managing TODO items with a prefix and optional logger.
type Todoer struct {
	prefix string      // prefix for TODO messages
	logger *log.Logger // optional logger for output
}

// ChangePrefix updates the prefix used for TODO messages.
func (t *Todoer) ChangePrefix(newPrefix string) {
	t.prefix = newPrefix
}

// AttachLogger attaches a logger to the Todoer for output.
func (t *Todoer) AttachLogger(logger *log.Logger) {
	t.logger = logger
}

// FancyPrint prints a formatted message with the Todoer's prefix.
// If a logger is attached, it logs the message; otherwise, it prints to stdout.
func (t *Todoer) FancyPrint(message string) {
	msg := fmt.Sprintf("%s\"%s\"", t.prefix, message)

	if t.logger != nil {
		t.logger.Println(msg)
	} else {
		fmt.Println(msg)
	}
}

// Opts represents default options for the Todoer.
var Opts = Todoer{prefix: "[TODO]: "}

// Nil prints a TODO message and returns nil.
func Nil(message ...string) any {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}
	return nil
}

// NilF returns a function that prints a TODO message and returns nil.
func NilF(message ...string) func(...any) any {
	return func(_ ...any) any {
		return Nil(message...)
	}
}

// Error prints a TODO message, optionally followed by an error, and returns a TodoError.
func Error(msgAndMock ...string) error {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return errors.New(msgAndMock[1])
	}

	return TodoError
}

// ErrorF returns a function that prints a TODO message, optionally followed by an error, and returns a TodoError.
func ErrorF(msgAndMock ...string) func(...any) error {
	return func(_ ...any) error {
		return Error(msgAndMock...)
	}
}

// String prints a TODO message, optionally followed by a string, and returns the string.
func String(msgAndMock ...string) string {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0])
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1]
	}

	return ""
}

// StringF returns a function that prints a TODO message, optionally followed by a string, and returns the string.
func StringF(msgAndMock ...string) func(...any) string {
	return func(_ ...any) string {
		return String(msgAndMock...)
	}
}

// Ptr returns a nil pointer, optionally printing a TODO message.
func Ptr[V any](message ...string) *V {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}

	return nil
}

// PtrF returns a function that returns a nil pointer, optionally printing a TODO message.
func PtrF[V any](message ...string) func(...any) *V {
	return func(_ ...any) *V {
		return Ptr[V](message...)
	}
}

// T prints a TODO message.
func T(message ...string) {
	if len(message) > 0 {
		Opts.FancyPrint(message[0])
	}
}

// Int prints a TODO message, optionally followed by an integer, and returns 0.
func Int(msgAndMock ...interface{}) int {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int)
	}

	return 0
}

// IntF returns a function that prints a TODO message, optionally followed by an integer, and returns 0.
func IntF(msgAndMock ...interface{}) func(...any) int {
	return func(_ ...any) int {
		return Int(msgAndMock...)
	}
}

// Int8 prints a TODO message, optionally followed by an int8, and returns 0.
func Int8(msgAndMock ...interface{}) int8 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int8)
	}

	return 0
}

// Int8F returns a function that prints a TODO message, optionally followed by an int8, and returns 0.
func Int8F(msgAndMock ...interface{}) func(...any) int8 {
	return func(_ ...any) int8 {
		return Int8(msgAndMock...)
	}
}

// Int32 prints a TODO message, optionally followed by an int32, and returns 0.
func Int32(msgAndMock ...interface{}) int32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int32)
	}

	return 0
}

// Int32F returns a function that prints a TODO message, optionally followed by an int32, and returns 0.
func Int32F(msgAndMock ...interface{}) func(...any) int32 {
	return func(_ ...any) int32 {
		return Int32(msgAndMock...)
	}
}

// UInt32 prints a TODO message, optionally followed by a uint32, and returns 0.
func UInt32(msgAndMock ...interface{}) uint32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(uint32)
	}

	return 0
}

// UInt32F returns a function that prints a TODO message, optionally followed by a uint32, and returns 0.
func UInt32F(msgAndMock ...interface{}) func(...any) uint32 {
	return func(_ ...any) uint32 {
		return UInt32(msgAndMock...)
	}
}

// Int64 prints a TODO message, optionally followed by an int64, and returns 0.
func Int64(msgAndMock ...interface{}) int64 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(int64)
	}

	return 0
}

// Int64F returns a function that prints a TODO message, optionally followed by an int64, and returns 0.
func Int64F(msgAndMock ...interface{}) func(...any) int64 {
	return func(_ ...any) int64 {
		return Int64(msgAndMock...)
	}
}

// Bool prints a TODO message, optionally followed by a bool, and returns false.
func Bool(msgAndMock ...interface{}) bool {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(bool)
	}

	return false
}

// BoolF returns a function that prints a TODO message, optionally followed by a bool, and returns false.
func BoolF(msgAndMock ...interface{}) func(...any) bool {
	return func(_ ...any) bool {
		return Bool(msgAndMock...)
	}
}

// Float32 prints a TODO message, optionally followed by a float32, and returns 0.0.
func Float32(msgAndMock ...interface{}) float32 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(float32)
	}

	return 0.0
}

// Float32F returns a function that prints a TODO message, optionally followed by a float32, and returns 0.0.
func Float32F(msgAndMock ...interface{}) func(...any) float32 {
	return func(_ ...any) float32 {
		return Float32(msgAndMock...)
	}
}

// Float64 prints a TODO message, optionally followed by a float64, and returns 0.0.
func Float64(msgAndMock ...interface{}) float64 {
	if len(msgAndMock) > 0 {
		Opts.FancyPrint(msgAndMock[0].(string))
	}
	if len(msgAndMock) > 1 {
		return msgAndMock[1].(float64)
	}

	return 0.0
}

// Float64F returns a function that prints a TODO message, optionally followed by a float64, and returns 0.0.
func Float64F(msgAndMock ...interface{}) func(...any) float64 {
	return func(_ ...any) float64 {
		return Float64(msgAndMock...)
	}
}
