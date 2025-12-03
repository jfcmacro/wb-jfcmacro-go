package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("Args[%d]=%s\n", i, arg)
	}
}
