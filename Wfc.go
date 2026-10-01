package main

import "fmt"

func main() {
	var N, hari int

	fmt.Scan(&N)
	hari = (4+N-1)%7 + 1

	fmt.Println(hari)
}
