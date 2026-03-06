package main

import (
	"fmt"
	"slices"
)

func main() {
	num := []string{"hello", "howareyou", "did you uderstand", "i dontknowifitworks bither"}
	bign := slices.Max(num)
	fmt.Printf("the biggest word is: %s \n", bign)

}
