package user

import (
	"github.com/set2002satoshi/8-4/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(db *gorm.DB, id int) (user models.User, err error)
	Create(db *gorm.DB, obj *models.User) (models.User, error)
}
