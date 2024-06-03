package todo_test

import (
	"errors"
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
