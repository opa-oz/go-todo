package todo_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/opa-oz/go-todo/todo"
	"github.com/stretchr/testify/assert"
)

func TestNil(t *testing.T) {
	assert.Nil(t, todo.Nil(), "Using empty Nil function")
	assert.Nil(t, todo.Nil("With message"))
	assert.Nil(t, todo.Nil("With", "many messages"))

	assert.Nil(t, todo.NilF()())
	assert.Nil(t, todo.NilF("With message")())
	assert.Nil(t, todo.NilF("With", "many messages")())
}

func TestString(t *testing.T) {
	assert.Equal(t, todo.String(), "")
	assert.Equal(t, todo.String("With message"), "")
	assert.Equal(t, todo.String("With default", "default_value"), "default_value")

	assert.Equal(t, todo.StringF()(), "")
	assert.Equal(t, todo.StringF("With message")(), "")
	assert.Equal(t, todo.StringF("With default", "default_value")(), "default_value")
}

func TestInts(t *testing.T) {
	assert.Equal(t, todo.Int(), 0)
	assert.Equal(t, todo.Int("With message"), 0)
	assert.Equal(t, todo.Int("With default", 5), 5)

	assert.Equal(t, todo.IntF()(), 0)
	assert.Equal(t, todo.IntF("With message")(), 0)
	assert.Equal(t, todo.IntF("With default", 6)(), 6)

	assert.Equal(t, todo.Int8(), int8(0))
	assert.Equal(t, todo.Int8("With message"), int8(0))
	assert.Equal(t, todo.Int8("With default", int8(5)), int8(5))

	assert.Equal(t, todo.Int8F()(), int8(0))
	assert.Equal(t, todo.Int8F("With message")(), int8(0))
	assert.Equal(t, todo.Int8F("With default", int8(6))(), int8(6))

	assert.Equal(t, todo.Int32(), int32(0))
	assert.Equal(t, todo.Int32("With message"), int32(0))
	assert.Equal(t, todo.Int32("With default", int32(5)), int32(5))

	assert.Equal(t, todo.Int32F()(), int32(0))
	assert.Equal(t, todo.Int32F("With message")(), int32(0))
	assert.Equal(t, todo.Int32F("With default", int32(6))(), int32(6))

	assert.Equal(t, todo.UInt32(), uint32(0))
	assert.Equal(t, todo.UInt32("With message"), uint32(0))
	assert.Equal(t, todo.UInt32("With default", uint32(5)), uint32(5))

	assert.Equal(t, todo.UInt32F()(), uint32(0))
	assert.Equal(t, todo.UInt32F("With message")(), uint32(0))
	assert.Equal(t, todo.UInt32F("With default", uint32(6))(), uint32(6))

	assert.Equal(t, todo.Int64(), int64(0))
	assert.Equal(t, todo.Int64("With message"), int64(0))
	assert.Equal(t, todo.Int64("With default", int64(5)), int64(5))

	assert.Equal(t, todo.Int64F()(), int64(0))
	assert.Equal(t, todo.Int64F("With message")(), int64(0))
	assert.Equal(t, todo.Int64F("With default", int64(6))(), int64(6))
}

func TestError(t *testing.T) {
	assert.Equal(t, todo.Error(), todo.TodoError)
	assert.Equal(t, todo.Error("With message"), todo.TodoError)
	assert.Equal(t, todo.Error("With custom error", "Custom error"), errors.New("Custom error"))

	assert.Equal(t, todo.ErrorF()(), todo.TodoError)
	assert.Equal(t, todo.ErrorF("With message")(), todo.TodoError)
	assert.Equal(t, todo.ErrorF("With custom error", "Custom error")(), errors.New("Custom error"))
}

func TestPtr(t *testing.T) {
	type mockS struct{}

	assert.Nil(t, todo.Ptr[mockS]())
	assert.Nil(t, todo.Ptr[mockS]("With message"))
	assert.Nil(t, todo.Ptr[mockS]("With", "many messages"))

	assert.Nil(t, todo.PtrF[mockS]()())
	assert.Nil(t, todo.PtrF[mockS]("With message")())
	assert.Nil(t, todo.PtrF[mockS]("With", "many messages")())
}

func TestStrings(t *testing.T) {
	printFancy := func(message string) {
		fmt.Println(
			fmt.Sprintf("[%s] -> %s", todo.String("Move prefix to function params", "MyServer"), message),
		)
	}

	printFancy("My very important message")
}

func TestInt(t *testing.T) {
	add := todo.IntF("Implement `add` function", 3)
	substract := todo.IntF("Implement `substract` function", 1)

	assert.Equal(t, add(1, 2), 3)
	assert.Equal(t, substract(add(1, 2), 4), 1)
}

func TestNilAndError(t *testing.T) {
	processor := func(shouldFail bool) (*string, error) {
		if shouldFail {
			return nil, todo.Error("Add custom error here")
		}

		return todo.Ptr[string]("Find string somewhere"), nil
	}

	r, err := processor(false)
	assert.Nil(t, err)
	assert.Nil(t, r)

	r, err = processor(true)
	assert.NotNil(t, err)
	assert.Nil(t, r)
}

func TestBool(t *testing.T) {
	isOdd := todo.BoolF("Implement `isOdd`", false)
	isEven := todo.BoolF("Implement `isOdd`", true)

	assert.True(t, isEven(3))
	assert.False(t, isOdd(3))
}

func TestFloat(t *testing.T) {
	// @see {@link https://en.wikipedia.org/wiki/Fast_inverse_square_root}
	Q_rsqrt := func(number float32) float32 {
		var x2, y float32

		threehalfs := todo.Float32("Un-magic threehalfs", float32(1.5))
		x2 = number * todo.Float32("Un-magic 1/2", float32(0.5))
		y = number
		// Note: The bit-level manipulation in C is not directly translatable to Go.
		// Instead, we use Go's math.Sqrt function to compute the square root.
		y = 1 / float32(math.Sqrt(float64(y)))
		y = y * (threehalfs - (x2 * y * y)) // 1st iteration
		// y = y * (threehalfs - (x2 * y * y)) // 2nd iteration, this can be removed

		return y
	}

	assert.Equal(t, Q_rsqrt(4.0), float32(0.5))
}
