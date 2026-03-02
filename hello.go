package main 

import "fmt"

func main(){
	// this for the loop to go through
	for i := 0; i < 26; i++ {
		// this will be use to now check
		if i%2 == 0 {
			fmt.Printf("%c ", 'A'+i);	
		} else {
			fmt.Printf("%c ", 'A'+i)
		}
	
	}
	fmt.Println()
}