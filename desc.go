package ezopt

import "reflect"

type desc struct {
	fn   reflect.Value
	args []reflect.Value
}

func newDesc(fn interface{}) (*desc, error) {
	vfn := reflect.ValueOf(fn)
	if vfn.Kind() != reflect.Func {
		return nil, newErrNotFunc(vfn)
	}
	// TODO:
	return &desc{
		fn: vfn,
	}, nil
}

func (d *desc) parse(args ...string) error {
	// TODO:
	return nil
}

func (d *desc) call() error {
	r := d.fn.Call(d.args)
	// TODO: return value r
	_ = r
	return nil
}
