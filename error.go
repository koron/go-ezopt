package ezopt

import (
	"errors"
	"reflect"
)

// Error represents ezopt's error.
type Error interface {
	error

	// Help returns help message.
	Help() string
}

var (
	ErrNoSubCommand = errors.New("no sub command")
)

func newErrNotFunc(v reflect.Value) error {
	return errors.New("not function")
}
