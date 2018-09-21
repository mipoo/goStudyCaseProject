package testcase

import (
	"fmt"
	"math/cmplx"
	"testing"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const f = "%T(%v)\n"

func p(v interface{}) {
	fmt.Printf(f, v, v)
}

func test() {

	p(ToBe)
	p(MaxInt)
	p(z)

	i := 456464548787235445 & 0xff
	p(i)

	const d float32 = 2.23
	a := 2.23
	p(a)

}

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func TestTv(t *testing.T) {
	test()

	println("start need")
	p(Small)
	p(needInt(Small))
	p(needFloat(Small))
	p(needFloat(Big))
	println("end need")
}

