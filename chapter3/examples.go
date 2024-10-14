package chapter3

import (
	"fmt"
	"math"
)

func Examples() {

	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("1 << 1: %08b\n", 1<<1) // shifting 1 by x will leave x zeroes on RHS
	fmt.Printf("1 << 5: %08b\n", 1<<5)
	fmt.Printf("x: %08b\n", x)

	fmt.Printf("1 << 1: %08b\n", 1<<1)
	fmt.Printf("1 << 2: %08b\n", 1<<2)
	fmt.Printf("y: %08b\n", y)

	fmt.Printf("x&y: %08b\n", x&y)
	fmt.Printf("x|y: %08b\n", x|y)
	// ^ is bitwise exclusive OR operator (XOR)
	// If AND result is true i.e. both x and y bits are 1 then resultant
	// bit is 0
	fmt.Printf("x^y: %08b\n", x^y)
	// if y's bit is 1 then resultant bit is 0
	// if y's bit is 0 then resultant bit is same as x's corresponding bit
	fmt.Printf("x&^y: %08b\n", x&^y)

	var p int8 = 1<<5 | 1<<6
	var n int8 = -p
	var a int8 = -64

	fmt.Printf("p: %08b n: %08b\n", p, n)

	fmt.Printf("shift p to left: %08b\n", p<<2)
	fmt.Printf("shift n to left: %08b\n", n<<2)

	fmt.Printf("shift p to rigth: %08b\n", p>>1)
	fmt.Printf("shift n to right: %08b\n", n>>1)

	fmt.Printf("a: %08b\n", a)
	fmt.Printf("right shfit a: %08b\n", a>>5)

	// float to int conversion of out of bound numbers
	f := 1e101
	// int(f) is dependent on the implementation. In this case, max int value is returned
	fmt.Println(int(f))

	// octal numbers
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	// hexadecimal numbers
	h := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", h)

	// rune literals
	ascii := 'a'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]q\n", newline)

	// Floating Numbers
	fmt.Printf("Max Float64: %8.3g\n", math.MaxFloat64) // 1.8e308
	fmt.Printf("Max Float32: %8.3g\n", math.MaxFloat32) // 3.4e38

	fmt.Printf("Min Float64: %8.3g\n", math.SmallestNonzeroFloat64) // 4.9e-324
	fmt.Printf("Min Float32: %8.3g\n", math.SmallestNonzeroFloat32) // 1.4e-45
}
