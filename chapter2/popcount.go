package chapter2

import "fmt"

// pc is population count or Hamming Weight
// population count is the number of bits set to 1
// pc is also called lookup table - as 'master list' used as a reference
var pc [256]byte

// init function is used to populate initial values at the startup
func init() {
	for i := range pc {
		fmt.Printf("i: %11d %b\n", i, byte(i))
		fmt.Printf("pc[i/2]: %5d\n", pc[i/2])
		fmt.Printf("i&1: %9d\n", i&1)
		fmt.Printf("byte(i&1) %4d\n", byte(i&1))

		// population cout of x = population count x/2 + add if least significant bit is 1
		pc[i] = pc[i/2] + byte(i&1)

		fmt.Printf("pc[i]: %7d\n", pc[i])
		fmt.Println("------------------")
	}
}

// Example of Popcount
// Let's take a large number - 0x123456789ABCDEF0 (in hex)
// This can be written in binary as
// 00010010 00110100 01010110 01111000 10011010 10111100 11011110 11110000
// All 8 bytes are used
// population count is 32

// Popcount1 returns population count of a 64 bits number
// It used a for loop
func Popcount1(x uint64) int {
	var val byte
	for i := 0; i < 8; i++ {
		// fmt.Printf("i: %5v i*8: %v\n", i, i*8)
		// fmt.Printf("x>>(i*8): %v\n", x>>(i*8))
		// fmt.Printf("pc[byte(x>>(i*8))]: %v\n", pc[byte(x>>(i*8))])

		val += pc[byte(x>>(i*8))]
	}
	// fmt.Printf("Population count for %v is %v\n", x, int(val))
	return int(val)
}

// Popcount2 doesn't use for loop
func Popcount2(x uint64) int {
	val := pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))]
	return int(val)
}
