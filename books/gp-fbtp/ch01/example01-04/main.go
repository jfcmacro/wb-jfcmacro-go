package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startUpTime := true, "info", time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
