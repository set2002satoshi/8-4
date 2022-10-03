package blog

import (

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