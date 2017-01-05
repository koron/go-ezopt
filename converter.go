package ezopt

import (
	"errors"
	"reflect"
	"strconv"
)

const convertTerminator = "--"

var basicTypes = map[reflect.Kind]converter{
	reflect.Bool:   &boolConverter{},
	reflect.String: &stringConverter{},

	reflect.Int:   &intConverter{size: 0},
	reflect.Int8:  &intConverter{size: 8},
	reflect.Int16: &intConverter{size: 16},
	reflect.Int32: &intConverter{size: 32},
	reflect.Int64: &intConverter{size: 64},

	reflect.Uint:   &uintConverter{size: 0},
	reflect.Uint8:  &uintConverter{size: 8},
	reflect.Uint16: &uintConverter{size: 16},
	reflect.Uint32: &uintConverter{size: 32},
	reflect.Uint64: &uintConverter{size: 64},

	reflect.Float32: &floatConverter{size: 32},
	reflect.Float64: &floatConverter{size: 64},
}

type converter interface {
	convert([]string) (reflect.Value, []string, error)
}

func findConverter(t reflect.Type) (converter, error) {
	k := t.Kind()
	c, ok := basicTypes[k]
	if ok {
		return c, nil
	}
	switch k {
	case reflect.Ptr:
		t2 := t.Elem()
		c, err := findConverter(t2)
		if err != nil {
			return nil, err
		}
		return &ptrConverter{
			t: t2,
			c: c,
		}, nil
	case reflect.Struct:
		c, err := newStructConverter(t)
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	return nil, errors.New("not supported type")
}

type boolConverter struct{}

func (c *boolConverter) convert(args []string) (reflect.Value, []string, error) {
	s, args := args[0], args[1:]
	b, err := strconv.ParseBool(s)
	if err != nil {
		return reflect.Value{}, nil, err
	}
	return reflect.ValueOf(b), args, nil
}

type stringConverter struct{}

func (c *stringConverter) convert(args []string) (reflect.Value, []string, error) {
	s, args := args[0], args[1:]
	return reflect.ValueOf(s), args, nil
}

type intConverter struct {
	size int
}

func (c *intConverter) convert(args []string) (reflect.Value, []string, error) {
	s, args := args[0], args[1:]
	n, err := strconv.ParseInt(s, 0, c.bitNum())
	if err != nil {
		return reflect.Value{}, nil, err
	}
	switch c.size {
	case 0:
		return reflect.ValueOf(int(n)), args, nil
	case 8:
		return reflect.ValueOf(int8(n)), args, nil
	case 16:
		return reflect.ValueOf(int16(n)), args, nil
	case 32:
		return reflect.ValueOf(int32(n)), args, nil
	case 64:
		return reflect.ValueOf(int64(n)), args, nil
	}
	panic("unknown size")
}

func (c *intConverter) bitNum() int {
	if c.size == 0 {
		return 32
	}
	return c.size
}

type uintConverter struct {
	size int
}

func (c *uintConverter) convert(args []string) (reflect.Value, []string, error) {
	s, args := args[0], args[1:]
	n, err := strconv.ParseUint(s, 0, c.bitNum())
	if err != nil {
		return reflect.Value{}, nil, err
	}
	switch c.size {
	case 0:
		return reflect.ValueOf(uint(n)), args, nil
	case 8:
		return reflect.ValueOf(uint8(n)), args, nil
	case 16:
		return reflect.ValueOf(uint16(n)), args, nil
	case 32:
		return reflect.ValueOf(uint32(n)), args, nil
	case 64:
		return reflect.ValueOf(uint64(n)), args, nil
	}
	panic("unknown size")
}

func (c *uintConverter) bitNum() int {
	if c.size == 0 {
		return 32
	}
	return c.size
}

type floatConverter struct {
	size int
}

func (c *floatConverter) convert(args []string) (reflect.Value, []string, error) {
	s, args := args[0], args[1:]
	n, err := strconv.ParseFloat(s, c.size)
	if err != nil {
		return reflect.Value{}, nil, err
	}
	switch c.size {
	case 32:
		return reflect.ValueOf(float32(n)), args, nil
	case 64:
		return reflect.ValueOf(float64(n)), args, nil
	}
	panic("unknown size")
}

type ptrConverter struct {
	t reflect.Type
	c converter
}

func (c *ptrConverter) convert(args []string) (reflect.Value, []string, error) {
	if args[0] == convertTerminator {
		return reflect.Zero(reflect.PtrTo(c.t)), args[1:], nil
	}
	v, args, err := c.c.convert(args)
	if err != nil {
		return reflect.Value{}, nil, err
	}
	p := reflect.New(c.t)
	p.Elem().Set(v)
	return p, args, nil
}

type structConverter struct {
	// TODO:
}

func newStructConverter(t reflect.Type) (*structConverter, error) {
	// TODO:
	return nil, nil
}

func (c *structConverter) convert(args[]string) (reflect.Value,[]string, error) {
	// TODO:
	return reflect.Value{}, nil, errors.New("not support yet")
}
