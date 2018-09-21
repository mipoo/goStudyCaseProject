package testcase

import (
	"testing"
	"math"
)

func TestNextAfter(t *testing.T) {
	for i := 0; i < 10; i += 1 {
		print(math.Nextafter(2, 5),"\t")
	}
	println()
}
