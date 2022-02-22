package myint

import (
	"errors"
	"math"
)

type MyInt int32

func (i MyInt) Add(in int) (MyInt, error) {
	if in >= math.MaxInt32 || int64(i)+int64(in) > math.MaxInt32 {
		return 0, errors.New("out of range")
	}

	return i + MyInt(in), nil
}

func (i MyInt) Sub(in int) MyInt {
	return i - MyInt(in)
}

func (i MyInt) Divide(in int) MyInt {
	return i / MyInt(in)
}

func (i MyInt) Multiply(in int) MyInt {
	return i * MyInt(in)
}
