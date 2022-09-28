package user

import (
	"github.com/set2002satoshi/8-4/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(db *gorm.DB) ([]models.ActiveUser, error)
	FindByID(db *gorm.DB, id int) (user models.ActiveUser, err error)
	Create(db *gorm.DB, obj *models.ActiveUser) (models.ActiveUser, error)
}
