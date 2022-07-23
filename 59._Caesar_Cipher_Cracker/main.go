package main

import "fmt"

// -1 encoding, +1 decoding

func main() {
	cipher("text", +1, 3)
}

func cipher(text string, choice int, shift rune) {
	chars := []rune(text)
	switch choice {
	case +1:
		for i := 0; i < len(chars); i++ {
			fmt.Println(string(chars[i] + shift))
		}
	case -1:
		fmt.Println("encoding")
	}
}
