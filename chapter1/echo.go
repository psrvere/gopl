package chapter1

import (
	"os"
	"strings"
)

func QuadraticConcatenation() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func LinearConcatenation() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(s)
}
