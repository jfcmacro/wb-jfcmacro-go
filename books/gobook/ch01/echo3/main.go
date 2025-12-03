package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	fmt.Println(s.Join(os.Args[1:], " "))
}
