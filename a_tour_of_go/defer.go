package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("couting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	fmt.Println("Done")
}
