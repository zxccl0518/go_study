package exercise

import (
	"reflect"
	"testing"
)

func TestReflectFunc(t *testing.T) {
	test1 := func(n1 int, n2 int) {
		t.Log(n1, n2)
	}

	test2 := func(n1 int, n2 int, s string) {
		t.Log(n1, n2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)

	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)

		inValue = make([]reflect.Value, n)

		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}

		function = reflect.ValueOf(call)

		function.Call(inValue)
	}

	bridge(test1, 1, 2)
	bridge(test2, 1, 2, "testlalal")
}
