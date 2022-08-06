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
	db.First(&user, id)
	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}