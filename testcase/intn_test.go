package testcase

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestIntn(t *testing.T) {
	for i := 0; i < 40; i += 1 {
		fmt.Print(rand.Intn(10), "\t")
	}
	fmt.Println()
}
