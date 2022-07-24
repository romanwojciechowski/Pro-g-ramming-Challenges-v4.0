package main

import "fmt"

// -1 encoding, +1 decoding

func main() {
	cipher("TEXT", +1, 3)
}

func cipher(text string, choice int, shift rune) {
	chars := []rune(text)
	switch choice {
	case +1:
		for i := 0; i < len(chars); i++ {
			if chars[i] >= 'a' && chars[i] <= 'z' || chars[i] >= 'A' && chars[i] <= 'Z' {
				dchar := chars[i] + shift
				if dchar >= 'a' && dchar <= 'z' || dchar >= 'A' && dchar <= 'Z' {
					fmt.Println(string(dchar))
				} else {
					fmt.Println(string(dchar - 26))
				}
			} else {
				fmt.Println(string(chars[i]))	
			}
		}
	case -1:
		fmt.Println("encoding")
	}
}
