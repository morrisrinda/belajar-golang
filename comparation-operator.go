package main

import "fmt"

func main() {
	var nama1 = "Morris"
	var nama2 = "Morris"
	var ceknama = nama1 == nama2
	fmt.Println(ceknama)

	var (
		value1 = 100
		value2 = 200
	)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}