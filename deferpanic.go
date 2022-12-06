package main

import "fmt"

func something() {
	defer fmt.Println("closed something!") // first

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 2 {
			panic("panic was called!")
		}
	}
}

func main() {
	defer fmt.Println("closed main")      //  2nd
	something()                           // stops here due to panic hit
	fmt.Println("something was finished") // last

}
