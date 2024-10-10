package chapter2

import "fmt"

func GCD(x, y int) {
	for y != 0 {
		x, y = y, x%y
	}
	fmt.Println("GCD: ", x)
}

func Fibonacci(n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	fmt.Printf("FIb for %v: %v\n", n, x)
}
