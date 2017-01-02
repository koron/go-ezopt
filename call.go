package ezopt

import "reflect"

func call(fn reflect.Value, args []reflect.Value) ([]reflect.Value, error) {
	r := fn.Call(args)
	if len(r) == 0 {
		return nil, nil
	}
	last := r[len(r)-1]
	if !last.CanInterface() {
		return r, nil
	}
	err, ok := last.Interface().(error)
	if !ok {
		return r, nil
	}
	return r[0 : len(r)-1], err
}
