package main

import "fmt"

func main() {
	var b *int
	fmt.Println(b) //<nil>

	var e string
	fmt.Printf("[%s]\n", e) // []

	var c bool
	fmt.Println(c) //false

}
