package main

import "fmt"

var (
	c, python, java bool
	i, j            int = 1, 2
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var a int
	c, python, java := true, false, "no!"
	k := 3

	fmt.Println(i, j, a, c, python, java)
	fmt.Println(k)

	// swap
	n, m := swap("hello", "world")
	fmt.Println(n, m)
}
