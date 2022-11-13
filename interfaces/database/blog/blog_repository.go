package blog

import (
	"errors"

	"github.com/set2002satoshi/8-4/models"
	"gorm.io/gorm"
)

type BlogRepository struct{}

func (repo *BlogRepository) FindAll(db *gorm.DB) ([]*models.ActiveBlog, error) {
	blog := []*models.ActiveBlog{}
	result := db.Find(&blog)
	if result.Error != nil {
		return []*models.ActiveBlog{}, result.Error
	}
	return blog, nil
}

func (repo *BlogRepository) FindByID(db *gorm.DB, id int) (*models.ActiveBlog, error) {
	var blog *models.ActiveBlog
	db.First(&blog, id)
	if blog.ActiveBlogID <= 0 {
		return &models.ActiveBlog{}, errors.New("blog is not found")
	}
	return blog, nil
}

func (repo *BlogRepository) UserFindByID(db *gorm.DB, id int) (*models.ActiveUser, error) {
	var user *models.ActiveUser
	db.Where("active_user_id", id).First(&user)
	if user.ActiveUserID <= 0 {
		return &models.ActiveUser{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *BlogRepository) Create(db *gorm.DB, data *models.ActiveBlog) (*models.ActiveBlog, error) {
	result := db.Create(data)
	if result.Error != nil {
		return nil, errors.New("データ作成に失敗")
	}

	return data, nil
}

func (repo *BlogRepository) Update(tx *gorm.DB, data *models.ActiveBlog) (*models.ActiveBlog, error) {
	if result := tx.Updates(data); result.Error != nil {
		return &models.ActiveBlog{}, errors.New("couldn't update blog failed")
	}
	var blog *models.ActiveBlog
	findResult := tx.Where("active_blog_id = ?", data.GetID()).First(&blog)
	if findResult.Error != nil {
		return &models.ActiveBlog{}, errors.New("couldn't find updated blog")
	}
	return blog, nil
}

func (repo *BlogRepository) DeleteByID(tx *gorm.DB, id int) error {
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
		return &models.HistoryBlog{}, errors.New("確認取得ができなかった")
	}
	return History, nil

}

