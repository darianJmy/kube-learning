package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(a)
	fmt.Println(a[:0])
	fmt.Println(a[:2])

}
