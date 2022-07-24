package main

import "fmt"

// +1 encoding, -1 decoding

func main() {
	var text string
	var choice int
	var shift int 

	fmt.Print("Text: ")
	fmt.Scanf("%s", &text)

	fmt.Print("+1 encoding, -1 decoding: ")
	fmt.Scanf("%d", &choice)

	fmt.Print("Shift: ")
	fmt.Scanf("%d", &shift)

	cipher(text, choice, shift)
}

func cipher(text string, choice int, shift int) {
	chars := []rune(text)
	for i := 0; i < len(chars); i++ {
		if chars[i] >= 'a' && chars[i] <= 'z' || chars[i] >= 'A' && chars[i] <= 'Z' {
			dchar := chars[i] + rune(shift*choice)
			if dchar >= 'a' && dchar <= 'z' || dchar >= 'A' && dchar <= 'Z' {
				fmt.Print(string(dchar))
			} else {
				fmt.Print(string(dchar + rune(-26 * choice)))
			}
		} else {
			fmt.Print(string(chars[i]))	
		}
	}
	fmt.Println()
}
