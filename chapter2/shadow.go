package chapter2

import (
	"fmt"
	"log"
	"os"
)

func ShadowExample1() {
	// x is declared in the function scope
	x := "hello!"
	fmt.Printf("1. x: %v, *x: %v\n", x, &x)
	for i := 0; i < len(x); i++ {
		fmt.Printf("2. x[i]: %v, type: %T, *x: %v\n", x[i], x[i], &x)

		// x is the local variable in the scope of for loop - body lexical scope
		// x[i] is the outer variable - function lexical scope
		x := x[i]

		fmt.Printf("3. x: %v, *x: %v\n", x, &x)

		// x is the local variable same as above in the if implicit lexical scope
		if x != '!' {
			fmt.Printf("4. x: %v, *x: %v\n", x, &x)

			// x is new local variable in the if body lexical scope
			x := x + 'A' - 'a'

			fmt.Printf("5. x: %v, *x: %v\n", x, &x)
			fmt.Printf("%c\n", x)
		}
	}
}

func ShadowExample2() {
	// x is declared in the function scope
	x := "hello"
	// x is same as above in 'range x'
	// new x declared in the for implicit scope
	for _, x := range x {
		// new x declared as 'x := ' in the for body scope
		// the x on right hand side is the x from for implicity loop
		x := x + 'A' - 'a'

		// here x is the local varibale from for body scope
		fmt.Printf("%c", x)
	}
}

// cwd decalred in the package scope
var cwd string

func ShandowExample3() {
	// both cwd and err are declared in the function scope
	// NOTE ':=' statement declares variable in the current scope even if
	// there are variable with the same name in the outer scopes
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	// This will print local cwd
	// NOTE package level variable is unused. Local cwd variable is shadowing is package
	// level cwd variable
	log.Printf("Current Working Directory is %v\n", cwd)
}
