package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func main() {
	var a int
	var c, python, java = true, false, "no!"
	k := 3

	fmt.Println(i, j, a, c, python, java)
	fmt.Println(k)
}
