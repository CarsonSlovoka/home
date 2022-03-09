package funcs

import (
	_math "carson.io/pkg/math"
	"carson.io/pkg/spf13/cast"
	"errors"
	"math"
)

// Add adds two numbers.
func Add(a, b interface{}) (interface{}, error) {
    return _math.Compute(a, b, '+')
}

// Sub subtracts two numbers.
func Sub(a, b interface{}) (interface{}, error) {
    return _math.Compute(a, b, '-')
}

// Mul multiplies two numbers.
func Mul(a, b interface{}) (interface{}, error) {
    return _math.Compute(a, b, '*')
}

// Div divides two numbers.
func Div(a, b interface{}) (interface{}, error) {
    return _math.Compute(a, b, '/')
}

// Ceil returns the least integer value greater than or equal to x.
func Ceil(x interface{}) (float64, error) {
    xf, err := cast.ToFloat64E(x)
    if err != nil {
        return 0, errors.New("ceil operator can't be used with non-float value")
    }

    return math.Ceil(xf), nil
}

// Floor returns the greatest integer value less than or equal to x.
func Floor(x interface{}) (float64, error) {
    xf, err := cast.ToFloat64E(x)
    if err != nil {
        return 0, errors.New("floor operator can't be used with non-float value")
    }

    return math.Floor(xf), nil
}

// Log (Nature Log: base e) returns the natural logarithm of a number.
func Log(a interface{}) (float64, error) {
    af, err := cast.ToFloat64E(a)
    if err != nil {
        return 0, errors.New("log operator can't be used with non integer or float value")
    }

    return math.Log(af), nil
}

func Sqrt(a interface{}) (float64, error) {
    af, err := cast.ToFloat64E(a)
    if err != nil {
        return 0, errors.New("sqrt operator can't be used with non integer or float value")
    }

    return math.Sqrt(af), nil
}

// Mod returns a % b.
func Mod(a, b interface{}) (int64, error) {
    ai, errA := cast.ToInt64E(a)
    bi, errB := cast.ToInt64E(b)

    if errA != nil || errB != nil {
        return 0, errors.New("modulo operator can't be used with non integer value")
    }

    if bi == 0 {
        return 0, errors.New("the number can't be divided by zero at modulo operation")
    }

    return ai % bi, nil
}

// ModBool returns the boolean of a % b.  If a % b == 0, return true.
func ModBool(a, b interface{}) (bool, error) {
    result, err := Mod(a, b)
    if err != nil {
        return false, err
    }

    return result == int64(0), nil
}

// Pow returns a raised to the power of b.
func Pow(a, b interface{}) (float64, error) {
    af, errA := cast.ToFloat64E(a)
    bf, errB := cast.ToFloat64E(b)

    if errA != nil || errB != nil {
        return 0, errors.New("pow operator can't be used with non-float value")
    }

    return math.Pow(af, bf), nil
}

// Round returns the nearest integer, rounding half away from zero.
func Round(x interface{}) (float64, error) {
    xf, err := cast.ToFloat64E(x)
    if err != nil {
        return 0, errors.New("round operator can't be used with non-float value")
    }

    return math.Round(xf), nil
}
