package blog

import (
	// "errors"

	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/usecase"
)

type BlogInteractor struct {
	DB   usecase.DBRepository
	Blog BlogRepository
}

func (i *BlogInteractor) FindByID(id int) (models.ActiveBlog, error) {
	db := i.DB.Connect()
	foundBlog, err := i.Blog.FindByID(db, id)
	if err != nil {
		return models.ActiveBlog{}, err
	}
	return foundBlog, nil
}


// func (i *BlogInteractor) Update(blog *models.ActiveBlog) (models.ActiveBlog, error) {
// 	db := i.DB.Begin()
// 	old, err := i.Blog.FindByID(db, int(blog.ID))
// 	if err != nil {
// 		return models.ActiveBlog{}, errors.New("Failed to retrieve the original data of change")
// 	}
// }