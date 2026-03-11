package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return true, "info", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()

	fmt.Println(Debug, LogLevel, startUpTime)
}
