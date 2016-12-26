package ezopt

func Run(funcOrMap interface{}, args ...string) error {
	switch v := funcOrMap.(type) {
	case Map:
		return v.run(args...)
	default:
		return runAsFunc(v, args...)
	}
}

func runAsFunc(fn interface{}, args ...string) error {
	d, err := newDesc(fn)
	if err != nil {
		return err
	}
	if err := d.parse(args...); err != nil {
		return err
	}
	return d.call()
}
