package main

import "fmt"

func main() {
	const zero float32 = 2.1
	const has bool = true 
	//fmt.Println("Hello world.", zero, " : ",has)
	var count int = (7 / 3)
	var count2 int = 7 % 3
	fmt.Println("Hello world.", count, count2)

	var i int32
	var j int64
	i, j = 1, 2

	// if i == j {
	// 	fmt.Println("i j equals.")
	// }

	if i == 1 || j == 2 {
		fmt.Println("equals.")
	}

}