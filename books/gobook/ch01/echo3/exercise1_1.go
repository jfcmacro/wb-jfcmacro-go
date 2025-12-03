package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	fmt.Printf("Program: %s Arguments: %s", os.Args[0], s.Join(os.Args[1:], " "))
}
