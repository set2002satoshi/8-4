package database

import (
	"errors"

	"gorm.io/gorm"

	"github.com/set2002satoshi/8-4/domain"
)


type UserRepository struct {}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user domain.Users, err error) {
	// DBの処理をラップしてる関数
	user = domain.Users{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}


func (repo *UserRepository) CreateUser(db *gorm.DB, obj domain.Users) (user domain.Users, err error) {
	var u domain.Users
	if result := db.Create(&obj); result.Error != nil {
		return domain.Users{}, errors.New("create user failed")
	}
	db.Where("id = ?", obj.ID).First(&u)
	return u, nil
}

