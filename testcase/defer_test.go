package testcase

import (
	"testing"
	"fmt"
)


func a() (i int){
	i = 0
	defer fmt.Println(i)
	defer func() {
		i++
		fmt.Println(i)
	}()
	i++
	return
}

func TestDefer(t *testing.T) {
	println(a())
}