package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type Options struct {
	dbstone *gorm.DB
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Server() *Options {
	o := &Options{
		dbstone: DB,
	}
	return o
}

func main() {
	s := Server()
	var user User
	if err := s.dbstone.Where("name = ?", "name").First(&user).Error; err != nil {
		panic(err)
	}
}

func init() {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=30s",
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"gorm")
	DB, err := gorm.Open("mysql", dbConnection)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection succeeded")
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.SingularTable(true)
	fmt.Println(DB)
}
