package dbstone

import "github.com/jinzhu/gorm"

type UserDB struct {
	dbstone *gorm.DB
}

func NewUserDB() *UserDB {
	return &UserDB{
		dbstone: DB,
	}
}
