package testcase

import (
	"testing"
	"fmt"
	"runtime"
	"time"
)

func Compare(a, b []int) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func test2() {
	var t interface{}
	t = 23
	switch b := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", b) // %T 输出 t 是什么类型
	case bool:
		fmt.Printf("boolean %t\n", b) // t 是 bool 类型
	case int:
		fmt.Printf("integer %d\n", b) // t 是 int 类型
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *b) // t 是 *bool 类型
	case *int:
		fmt.Printf("pointer to integer %d\n", *b) // t 是 *int 类型
	}
}

func test3() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	println()
}

func test4() {
	println("test4 .... ")
	switch {
	case func() bool {
		println("111")
		return false
	}():
	case func() bool {
		println("222")
		return true
	}():
	case func() bool {
		println("333")
		return false
	}():
	}
}

func TestSwitch(t *testing.T) {
	a := []int{2, 3}
	b := []int{1, 2, 3}
	println(Compare(a, b))

	test2()

	test3()

	test4()


	p(time.Now().Weekday())
	p(time.Saturday)
}
