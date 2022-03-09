package math_test

import (
	"carson.io/pkg/math"
	"errors"
	"fmt"
	"testing"
)

func TestCompute(t *testing.T) {
	t.Parallel()
	fmt.Printf("%v\n", map[string]rune{"+": '+', "-": '-', "*": '*', "/": '/'})
	for _, test := range []struct {
		a      interface{}
		b      interface{}
		op     rune
		expect interface{}
	}{
		{3.0, 2, '+', float64(5)},
		{3.0, 2, '-', float64(1)},
		{3.0, 2, '*', float64(6)},
		{3.0, 2, '/', 1.5}, // default: float64
		{uint(3), uint(2), '+', float64(5)},
		{uint(3), uint(2), '-', float64(1)},
		{uint(3), uint(2), '*', float64(6)},
		{uint(3), uint(2), '/', 1.5},
		{uint(3), -2, '+', float64(1)},
		{uint(3), -2, '-', float64(5)},
		{uint(3), -2, '*', float64(-6)},
		{uint(3), -2, '/', -1.5},
		{uint(3), 2.0, '+', float64(5)},
		{uint(3), 2.0, '-', float64(1)},
		{uint(3), 2.0, '*', float64(6)},
		{uint(3), 2.0, '/', 1.5},
		{0, 0, '+', float64(0)},
		{0, 0, '-', float64(0)},
		{0, 0, '*', float64(0)},
		{0, 0, '/', errors.New("can't divide the value by 0")},
		{3, 0, '/', errors.New("can't divide the value by 0")},
		{"3.0", "2", '+', float64(5)},
		{"3.0", "2", '-', float64(1)},
		{"3.0", "2", '*', float64(6)},
		{"3.0", "2", '/', 1.5},
		{"3.0", 2, '+', float64(5)},
		{"3.0", 2, '-', float64(1)},
		{"3.0", 2, '*', float64(6)},
		{"3.0", 2, '/', 1.5},
		{true, true, '+', errors.New("type error")},
		{true, false, '-', errors.New("type error")},
		{false, true, '*', errors.New("type error")},
		{false, false, '/', errors.New("type error")},
	} {
		actual, err := math.Compute(test.a, test.b, test.op)
		if err != nil {
			if test.expect.(error).Error() != err.Error() {
				t.Fatalf("%v\n%s\n%s", test, test.expect.(error).Error(), err.Error())
			}
		} else {
			if test.expect != actual {
				t.Fatalf("%v", test)
			}
		}
	}
}
