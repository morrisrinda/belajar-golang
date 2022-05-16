/**
 ** Konversi Data
 ** int 8 max 127, jika nilai32 128, maka akan jatuh ke nilai minimumnya -128
 */

package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // Max 127

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Morris"
	var e = name[2]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}