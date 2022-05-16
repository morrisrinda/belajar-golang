package main

import "fmt"

func main() {
	type noKTP string
	type married bool

	var noKTPnya noKTP = "414343241234124121"
	fmt.Println(noKTPnya)

	var marriedStatus married = true
	fmt.Println(marriedStatus)
}