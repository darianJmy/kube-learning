package main

import (
	"context"
	"fmt"
	"kube-learning/practise/gorm-practise/dbstone"
)

var db = dbstone.NewUserDB()

func main() {
	user, err := db.Get(context.TODO(), "jixingxing")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)

	users, err := db.List(context.TODO(), "jixingxing")
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	err = db.Update(context.TODO(), "jixingxing", 19)
	if err != nil {
		panic(err)
	}
	fmt.Println("update ok")

	err = db.Delete(context.TODO(), "jixingxing", 20)
	if err != nil {
		panic(err)
	}
	fmt.Println("delete ok")
}
