package main

import "fmt"

func main() {
	m := make(map[int]int)

	m = map[int]int{
		1: 1,
		2: 2,
	}

	fmt.Println(m)

	m2 := map[int]int{
		3:3,
		4:4,
	}

	fmt.Println(m2)
}
