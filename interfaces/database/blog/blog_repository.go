package blog

import (
	"errors"


	"gorm.io/gorm"
	"github.com/set2002satoshi/8-4/models"
)


type BlogRepository struct {}

func (repo *BlogRepository) FindByID(db *gorm.DB, id int) (models.ActiveBlog, error) {
	blog:= models.ActiveBlog{}
	db.First(&blog, id)
	if blog.ActiveBlogID <= 0 {
		return models.ActiveBlog{}, errors.New("user is not found")
	}
	return blog, nil
}

