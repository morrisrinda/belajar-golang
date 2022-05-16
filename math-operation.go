package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	var c = a + b
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -100
	fmt.Println(negative)
	negative--
	fmt.Println(negative)
}