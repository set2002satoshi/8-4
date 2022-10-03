package blog

import (
	"github.com/set2002satoshi/8-4/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	FindByID(db *gorm.DB, id int) (models.ActiveBlog, error)
}
