package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
func Multi(a int, b int) int {
	return a * b
}
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	result := Add(a, b)*Sub(a, b) - Multi(a, b)
	fmt.Println(result)
}
