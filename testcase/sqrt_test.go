package testcase

import (
	"testing"
)

const fix = 0.000001

func abs(x float64) float64 {
	if x > 0 {
		return x
	}
	return -x
}

func Sqrt(x float64) float64 {
	z := x / 2
	count := 0
	for ; abs(x-z*z) > fix; count += 1 {
		z = z - (z*z-x)/(2*z)
	}
	println(count)
	return z
}

func TestSqrt(t *testing.T) {
	println(Sqrt(1024))
}
