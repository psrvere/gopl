package echo1

import (
	"os"
	"strings"
)

func quadraticConcatenation() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	// fmt.Println(s)
}

func linearConcatenation() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(s)
}
