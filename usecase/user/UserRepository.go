package user

import (
	"github.com/set2002satoshi/8-4/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(db *gorm.DB, id int) (user domain.User, err error)
	// CreateUser(db *gorm.DB, obj *domain.Users) (domain.Users, error)
}
