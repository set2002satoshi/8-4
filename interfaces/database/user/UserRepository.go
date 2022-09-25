package user

import (
	"errors"
	"gorm.io/gorm"

	"github.com/set2002satoshi/8-4/models"
)


type UserRepository struct {}

func (repo *UserRepository) Find(db *gorm.DB, id int) (user models.User, err error) {
	// DBの処理をラップしてる関数
	user = models.User{}
	db.First(&user, id)
	if user.ID <= 0 {
		return models.User{}, errors.New("user is not found")
	}
	return user, nil
}


func (repo *UserRepository) Create(db *gorm.DB, obj *models.User) (user models.User, err error) {
	if result := db.Create(obj); result.Error != nil {
		return models.User{}, errors.New("create user failed")
	}
	var createdUser models.User
	db.Where("ID = ?", obj.GetID()).First(&createdUser)
	return createdUser, nil
}


