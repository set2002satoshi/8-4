package blog

import (
	"errors"


	"gorm.io/gorm"
	"github.com/set2002satoshi/8-4/models"
)


type BlogRepository struct {}

func (repo *BlogRepository) FindByID(db *gorm.DB, id int) (*models.ActiveBlog, error) {
	blog:= models.ActiveBlog{}
	db.First(&blog, id)
	if blog.ActiveBlogID <= 0 {
		return &models.ActiveBlog{}, errors.New("user is not found")
	}
	return &blog, nil
}


func (repo *BlogRepository) Create(db *gorm.DB,data *models.ActiveBlog) (*models.ActiveBlog, error) {
	result := db.Create(data)
	if result.Error != nil {
		return nil, errors.New("データ作成に失敗")
	}

	return data, nil
}

func (repo *BlogRepository) DeleteByID(tx *gorm.DB,id int)  error {
	activeBlog := []models.ActiveBlog{}
	if result := tx.Unscoped().Delete(activeBlog, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *BlogRepository) InsertHistory(tx *gorm.DB, data *models.HistoryBlog) (*models.HistoryBlog, error) {
	createResult := tx.Create(data)
	if createResult.Error != nil {
		return &models.HistoryBlog{}, createResult.Error
	}
	var History *models.HistoryBlog
	findResult := tx.Where("history_blog_id = ?", data.HistoryBlogID).First(&History)
	if findResult.Error != nil {
		return &models.HistoryBlog{}, findResult.Error
	}
	return History, nil

}