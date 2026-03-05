package main

import "fmt"

func main() {
	var x float32 = 123.78
	var y float64 = 3.4e+38 //this is testing the float for 64 and the difference with float32

	fmt.Printf("Type : %T, value : %v\n", x, x)
	fmt.Printf("Type : %T, value : %v\n", y, y)

}
