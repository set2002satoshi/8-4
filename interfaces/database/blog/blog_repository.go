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
	resultData := models.ActiveBlog{}
	findResult := db.First(&resultData, int(data.ActiveBlogID))
	if findResult.Error != nil {
		return nil, errors.New("作成したデータを取得できなかった")
	}
	return &resultData, nil
}