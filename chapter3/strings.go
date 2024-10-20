package chapter3

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func ContainsEfficient(s, substr string) bool {
	return strings.Contains(s, substr) // used stings.Index underneath
}

func RuneSize(s string) {
	fmt.Println("size with len(s): ", len(s))
	fmt.Println("size with utf8.RuneCountInString(s): ", utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}
