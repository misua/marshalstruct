package main

import "fmt"

func a(i int) {
	i = 0
}

func b(i *int) {
	*i = 5
}

func main() {

	x := 100

	a(x)
	fmt.Println("value if a(x) = ", x)

	b(&x)
	fmt.Println("X value actually pointed to =", x)
	fmt.Println("mem address of a(*x) =", &x)

}
