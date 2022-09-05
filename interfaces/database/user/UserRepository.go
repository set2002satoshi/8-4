package user

import (
	"errors"

	"gorm.io/gorm"

	"github.com/set2002satoshi/8-4/domain"
)


type UserRepository struct {}

func (repo *UserRepository) Find(db *gorm.DB, id int) (user domain.User, err error) {
	// DBの処理をラップしてる関数
	user = domain.User{}
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.User{}, errors.New("user is not found")
	}
	return user, nil
}


// func (repo *UserRepository) CreateUser(db *gorm.DB, obj *domain.Users) (user domain.Users, err error) {
// 	if result := db.Create(&obj); result.Error != nil {
// 		return domain.Users{}, errors.New("create user failed")
// 	}
// 	createdUser := domain.Users{}
// 	db.Where(createdUser, obj.GetID()).First(createdUser)
// 	return createdUser, nil
// }


