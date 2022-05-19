package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v2 = Vertex{X:1}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)


	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v2, v3, p)
}
