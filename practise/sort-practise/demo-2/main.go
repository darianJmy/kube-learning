package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	a := User{
		Name: "jixingxing",
		Age:  18,
	}
	b := User{
		Name: "jimingyu",
		Age:  17,
	}
	c := User{
		Name: "yangxiaoqian",
		Age:  19,
	}
	d := Users{a, b, c}
	fmt.Println(d)
	sort.Sort(d)
	fmt.Println(d)

}
