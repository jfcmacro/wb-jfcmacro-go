package main

import "math/rand"

func main() {
	// This generate a build error var seed = 123445677
	var seed int64 = 1234456789
	// # jfcmacro/wb-jfcmacro-go/books/gp-fbtp/ch01/example01-03
        // ./main.go:10:17: cannot use seed (variable of type int) as int64 value in argument to rand.NewSource
	rand.NewSource(seed)
}
