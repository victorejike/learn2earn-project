package main

import "fmt"

func HashCode(dec string) string {
	//the size and length of input text or string
	size := len(dec)
	result := ""

	//this where i started using loop
	for _, r := range dec {
		newChar := (int(r) + size) % 127

		// now i will now use an if statment to complar both side
		if newChar < 32 {
			newChar += 33
		}

		result += string(rune(newChar))

	}
	return result

}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
