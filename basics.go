package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, a := range nums {
		total = total + a
	}
	return total

}
func main() {

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("first five numbers is", total)

}
