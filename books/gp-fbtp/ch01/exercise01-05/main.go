package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := true
	LogLevel := "info"
	startUpTime := time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
