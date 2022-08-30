package usecase

import (
	"github.com/set2002satoshi/8-4/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(db *gorm.DB, id int) (user domain.Users, err error)
	Save(db *gorm.DB, obj *domain.Users) (*domain.Users,  error)
}

