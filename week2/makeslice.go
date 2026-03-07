package main

import "fmt"

func main() {
	//using this to declare the make slice
	myslice := make([]int, 5, 10)

	//now let us print out the output
	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("lenght = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

}
