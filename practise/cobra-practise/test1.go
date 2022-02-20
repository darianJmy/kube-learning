package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
const (
	USER   = "root"
	PWD    = "123456"
	DBIP   = "192.168.245.145"
	DBPORT = "3306"
	DBNAME = "mysql"
)

type User struct {
	Host     string
	User     string
	Password string
}

func main() {
	info := USER + ":" + PWD + "@tcp(" + DBIP + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local&timeout=10ms"
	fmt.Println(info)
	db, err := gorm.Open("mysql", info)
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	} else {
		fmt.Println("mysql connect success")
	}
	var user User
	fmt.Println(db.Select([]string{"Host", "User", "Password"}).Where("Host = ?", "localhost").First(&user))
	db.Last(&user)
	fmt.Println(user)
}