package main

import "fmt"

func main() {
	// for
	sum_1 := 0
	for i := 0; i < 10; i++ {
		sum_1 += i
	}

	fmt.Println(sum_1)

	// for continued
	sum_2 := 1
	for sum_2 < 100 {
		sum_2 += sum_2
	}

	fmt.Println(sum_2)

	// forever
	// for {
	// }

	// while
	sum_3 := 1
	for sum_3 < 100 {
		sum_3 += sum_3
	}

	fmt.Println(sum_3)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
