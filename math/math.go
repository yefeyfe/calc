package math

import "fmt"

func init() {
	fmt.Println("initialized package math (local).")
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}