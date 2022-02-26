package main

import (
	"fmt"
	"kube-learning/practise/gorm-practise/dbstone"
)

var db = dbstone.NewUserDB()

func main() {
	fmt.Println(db)
}

