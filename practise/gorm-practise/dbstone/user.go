package dbstone

import (
	"context"
	"github.com/jinzhu/gorm"
	"kube-learning/practise/gorm-practise/models"
)

type UserDB struct {
	dbstone *gorm.DB
}

func NewUserDB() *UserDB {
	return &UserDB{
		dbstone: DB,
	}
}

type UserInterface interface {
	Get(ctx context.Context, name string) (*models.User, error)
	List(ctx context.Context, name string) (*[]models.User, error)
	Update(ctx context.Context, name string, age int) error
	Delete(ctx context.Context, name string, age int) error
}

func (u *UserDB) Get(ctx context.Context, name string) (*models.User, error) {
	var user models.User
	if err := u.dbstone.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDB) List(ctx context.Context, name string) (*[]models.User, error) {
	var users []models.User
	if err := u.dbstone.Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (u *UserDB) Update(ctx context.Context, name string, age int) error {
	var user models.User
	return u.dbstone.Model(&user).Where("name = ?", name).Update("age", age).Error
}

func (u *UserDB) Delete(ctx context.Context, name string, age int) error {
	var user models.User
	return u.dbstone.Where("name = ? AND age = ?", name, age).Delete(&user).Error
}
