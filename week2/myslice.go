package main

import (
	"fmt"
	"slices"
)

func main() {
	myslice := []int{20, 30, 40, 50, 60, 70}
	//now we have to now give the slice a range  because that  is what we are look for
	maxvalue := slices.Max(myslice)
	fmt.Printf("the biggest slice is: %d\n", maxvalue)
}
