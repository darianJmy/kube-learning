package main

import "fmt"

type delta struct {
	name string
}

type queue interface {
	abc(name string)
}

var name1 = queue(&delta{name: "jixingxing"})

func (d *delta) abc(name string) {

}
func main() {
	fmt.Println(name1)

}
