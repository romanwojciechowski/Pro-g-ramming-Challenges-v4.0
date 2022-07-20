package main

import "fmt"

func main() {
	var r int
	
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &r)
	
	for i := 1; i <= r; i++ {
		fizzbuzz(i)
	}
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(fizz, buzz)
	} else if i%5 == 0 {
		fmt.Println(fizz)
	} else if i%3 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(i)
	}
}
