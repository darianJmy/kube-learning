package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	USER   = "root"
	PWD    = "123456"
	DBIP   = "127.0.0.1"
	DBPORT = "3306"
	DBNAME = "gorm"
)

const (
	name = "jixingxing"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age string  `json:"age"`
}
func main() {
	info := USER + ":" + PWD + "@tcp(" + DBIP + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	db, err := gorm.Open("mysql", info)
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	} else {
		fmt.Println("mysql connect success")
	}
	var user User

	err = db.Where("name = ?", name).First(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}