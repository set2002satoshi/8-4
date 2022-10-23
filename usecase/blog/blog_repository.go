package blog

import (
	"github.com/set2002satoshi/8-4/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	FindByID(db *gorm.DB, id int) (*models.ActiveBlog, error)
	Create(db *gorm.DB, data *models.ActiveBlog) (*models.ActiveBlog, error)
	DeleteByID(tx *gorm.DB, id int) error
	InsertHistory(tx *gorm.DB, data *models.HistoryBlog) (*models.HistoryBlog, error)

	
}
