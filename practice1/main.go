package main

import (
	"fmt"
)

func main() {
	str_array := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println(str_array)

	// func 1
	// for index, value := range str_array {
	// 	fmt.Println(index)
	// 	fmt.Println(value)
	// 	str_array[2] = "smart"
	// 	str_array[4] = "strong"
	// }

	// func 2
	// for i, v := range str_array {
	// 	if v == "stupid" {
	// 		str_array[i] = "smart"
	// 	}
	// 	if v == "weak" {
	// 		str_array[i] = "strong"
	// 	}
	// }

	// func 3
	for i, v := range str_array {
		switch v {
		case "stupid":
			str_array[i] = "smart"
		case "weak":
			str_array[i] = "strong"
		default:
		}
	}
	fmt.Println(str_array)
}
