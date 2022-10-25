package blog

import (
	"errors"
	"time"

	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
	"github.com/set2002satoshi/8-4/usecase"
)

type BlogInteractor struct {
	DB   usecase.DBRepository
	Blog BlogRepository
}

func (i *BlogInteractor) FindByID(id int) (*models.ActiveBlog, error) {
	db := i.DB.Connect()
	foundBlog, err := i.Blog.FindByID(db, id)
	if err != nil {
		return &models.ActiveBlog{}, err
	}
	return foundBlog, nil
}

func (i *BlogInteractor) Post(data *models.ActiveBlog) (*models.ActiveBlog, error) {
	tx := i.DB.Connect()
	createdBlog, err := i.Blog.Create(tx, data)
	if err != nil {
		return &models.ActiveBlog{}, err
	}
	return createdBlog, nil
}


func (i *BlogInteractor) DeleteByID(id int) (*models.HistoryBlog, error) {
	tx := i.DB.Begin()
	activeBlog, err  := i.Blog.FindByID(tx, id)
	if err != nil {
		return &models.HistoryBlog{}, nil
	}
	convertedHistoryBlog, err := i.toHistory(activeBlog)
	if err != nil {
		return &models.HistoryBlog{}, err
	}
	resultHistory, err := i.Blog.InsertHistory(tx, convertedHistoryBlog)
	if err != nil {
		tx.Rollback()
		return &models.HistoryBlog{}, err
	}
	if err := i.Blog.DeleteByID(tx, id); err != nil {
		tx.Rollback()
		return &models.HistoryBlog{},  err
	}
	commitResult := tx.Commit()
	if commitResult.Error != nil {
		return &models.HistoryBlog{}, errors.New("can not commit")
	}
	return resultHistory, nil

}





func (i *BlogInteractor) toHistory(data *models.ActiveBlog) (*models.HistoryBlog, error) {
	return models.NewHistoryBlog(
		temporary.INITIAL_ID,
		int(data.GetID()),
		data.GetTitle(),
		data.GetContext(),
		time.Time{},
		time.Time{},
		data.GetCreatedAt(),
		data.GetRevision(),
	)
}