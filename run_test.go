package ezopt

import (
	"log"
	"testing"
)

func TestRunSimple(t *testing.T) {
	var arg0 int
	fn := func(n int) {
		arg0 = n
	}
	check := func(s string, ex int) {
		err := Run(fn, s)
		if err != nil {
			t.Errorf("error: %s", err)
			return
		}
		if arg0 != ex {
			t.Errorf("expected:%d, actual:%d", ex, arg0)
		}
	}
	check("123", 123)
	check("0", 0)
	check("-1", -1)
}

func TestRunMuch(t *testing.T) {
	var arg0 int
	fn := func(n int) {
		arg0 = n
	}
	err := Run(fn, "123", "456")
	if err == nil {
		t.Fatal("should return error")
	}
	// TODO: check err restrictly
	log.Printf("Run()=%s", err)
}

func TestRunLess(t *testing.T) {
	var arg0 int
	fn := func(n int) {
		arg0 = n
	}
	err := Run(fn)
	if err == nil {
		t.Fatal("should return error")
	}
	// TODO: check err restrictly
	log.Printf("Run()=%s", err)
}