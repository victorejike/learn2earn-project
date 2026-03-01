package main

import "fmt"

func main() {
	for i := 0; i < 26; i++ {
		if i%2 == 0 {
			fmt.Printf("%c ", 'a'+i)
		} else {
			fmt.Printf("%c ", 'A'+i)
		}
	}
	fmt.Printf("")
}
