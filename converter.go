package ezopt

import (
	"errors"
	"log"
	"reflect"
	"strconv"
)

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
	convert(s string) (reflect.Value, error)
}

func findConverter(t reflect.Type) (converter, error) {
	k := t.Kind()
	c, ok := basicTypes[k]
	if ok {
		return c, nil
	}
	if k == reflect.Ptr {
		c, err := findConverter(t.Elem())
		if err != nil {
			return nil, err
		}
		return &ptrConverter{c: c}, nil
	}
	return nil, errors.New("not supported type")
}

type boolConverter struct{}

func (c *boolConverter) convert(s string) (reflect.Value, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return reflect.Value{}, err
	}
	return reflect.ValueOf(b), nil
}

type stringConverter struct{}

func (c *stringConverter) convert(s string) (reflect.Value, error) {
	return reflect.ValueOf(s), nil
}

type intConverter struct {
	size int
}

func (c *intConverter) convert(s string) (reflect.Value, error) {
	sz := c.size
	if sz == 0 {
		sz = 32
	}
	n, err := strconv.ParseInt(s, 0, sz)
	if err != nil {
		return reflect.Value{}, err
	}
	switch c.size {
	case 0:
		return reflect.ValueOf(int(n)), nil
	case 8:
		return reflect.ValueOf(int8(n)), nil
	case 16:
		return reflect.ValueOf(int16(n)), nil
	case 32:
		return reflect.ValueOf(int32(n)), nil
	case 64:
		return reflect.ValueOf(int64(n)), nil
	}
	panic("unknown size")
}

type uintConverter struct {
	size int
}

func (c *uintConverter) convert(s string) (reflect.Value, error) {
	sz := c.size
	if sz == 0 {
		sz = 32
	}
	n, err := strconv.ParseUint(s, 0, sz)
	if err != nil {
		return reflect.Value{}, err
	}
	switch c.size {
	case 0:
		return reflect.ValueOf(uint(n)), nil
	case 8:
		return reflect.ValueOf(uint8(n)), nil
	case 16:
		return reflect.ValueOf(uint16(n)), nil
	case 32:
		return reflect.ValueOf(uint32(n)), nil
	case 64:
		return reflect.ValueOf(uint64(n)), nil
	}
	panic("unknown size")
}

type floatConverter struct {
	size int
}

func (c *floatConverter) convert(s string) (reflect.Value, error) {
	n, err := strconv.ParseFloat(s, c.size)
	if err != nil {
		return reflect.Value{}, err
	}
	switch c.size {
	case 32:
		return reflect.ValueOf(float32(n)), nil
	case 64:
		return reflect.ValueOf(float64(n)), nil
	}
	panic("unknown size")
}

type ptrConverter struct {
	c converter
}

func (c *ptrConverter) convert(s string) (reflect.Value, error) {
	v, err := c.c.convert(s)
	if err != nil {
		return reflect.Value{}, err
	}
	log.Printf("%#v", v)
	return v.Addr(), nil
}
