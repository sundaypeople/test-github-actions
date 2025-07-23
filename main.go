package main

import "fmt"

func main() {
	a := 1
	b := 2
	ans := add(a, b)
	fmt.Println("The sum is:", ans)
}

func add(x, y int) int {
	return x + y
}
