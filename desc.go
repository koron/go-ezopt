package ezopt

import (
	"errors"
	"reflect"
)

type desc struct {
	fn   reflect.Value
	args []reflect.Value

	numIn    int
	variadic bool
	convs    []converter
}

func newDesc(fn interface{}) (*desc, error) {
	vfn := reflect.ValueOf(fn)
	if vfn.Kind() != reflect.Func {
		return nil, newErrNotFunc(vfn)
	}
	var (
		tfn      = vfn.Type()
		numIn    = tfn.NumIn()
		variadic = tfn.IsVariadic()
		convs    = make([]converter, 0, numIn)
	)
	for i := 0; i < numIn; i++ {
		c, err := findConverter(tfn.In(i))
		if err != nil {
			return nil, err
		}
		convs = append(convs, c)
	}
	// TODO: build descriptors
	return &desc{
		fn:       vfn,
		numIn:    numIn,
		variadic: variadic,
		convs:    convs,
	}, nil
}

func (d *desc) parse(args ...string) error {
	n := 0
	for i := 0; i < len(args); i++ {
		s := args[i]
		// parse normal string as an arg.
		if n >= len(d.convs) {
			if !d.variadic {
				// FIXME:
				return errors.New("too much args")
			}
			n = len(d.convs) - 1
		}
		v, err := d.convs[n].convert(s)
		if err != nil {
			return err
		}
		d.args = append(d.args, v)
		n++
	}
	if d.numIn != len(d.args) {
		return errors.New("too less args")
	}
	return nil
}

func (d *desc) call() error {
	_, err := call(d.fn, d.args)
	return err
}
