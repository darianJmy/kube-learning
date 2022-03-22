package main

import (
	"fmt"
	"sort"
	"time"
)

type User struct {
	Name string
	Age int
	time time.Time
}

type Users []User

func (u Users) Len() int {
	return len(u)
}

func (u Users) Less(i, j int) bool {
	return  u[i].time.Before(u[j].time)
}

func (u Users) Swap(i, j int)  {
	u[i], u[j] = u[j], u[i]
}

func main() {
	a := User{
		Name: "jixingxing",
		Age: 18,
		time: time.Now(),
	}

	time.Sleep(1 * 20)

	b := User{
		Name: "jimingyu",
		Age: 18,
		time: time.Now(),
	}

	c := User{
		Name: "yangxiaoqian",
		Age: 20,
		time: time.Now(),
	}

	d := Users{a, b, c}

	fmt.Println(d)
	sort.Sort(d)
	fmt.Println(d)

}