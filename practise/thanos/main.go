package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	a := filepath.Base(os.Args[0])
	fmt.Print(a)

	target := new(bool)
	fmt.Print(*target)

	b := new(string)
	fmt.Print(*b)
}
