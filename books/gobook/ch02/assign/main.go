package main

import "fmt"

func main() {
	var x,y = 12,3
	fmt.Printf("gcd(%d,%d)=%d\n", x, y, gcd(x,y))

	for i:=20;i<40;i++ {
		fmt.Printf("fib(%d)=%d\n", i, fib(i))
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x,y = y, x % y
	}
	return x
}

func fib(n int) int {
	x,y := 0, 1
	for i:=0;i<n;i++ {
		x,y = y,x+y
	}
	return x
}
