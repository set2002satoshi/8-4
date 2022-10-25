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


func (i *BlogInteractor) FindAll() ([]*models.ActiveBlog, error) {
	db := i.DB.Connect()
	foundBlog, err := i.Blog.FindAll(db)
	if err != nil {
		return nil, err
	}
	return foundBlog, nil
}

func (i *BlogInteractor) FindByID(id int) (*models.ActiveBlog, error) {
	db := i.DB.Connect()
	foundBlog, err := i.Blog.FindByID(db, int(id))
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

func (i *BlogInteractor) Update(data *models.ActiveBlog) (*models.ActiveBlog, error) {
	tx := i.DB.Begin()

	oldActiveBlog, err := i.Blog.FindByID(tx, int(data.ActiveBlogID))
	if err != nil {
		return nil, errors.New("can not find active user")
	}
	if oldActiveBlog.GetRevision() != data.GetRevision() {
		return &models.ActiveBlog{}, errors.New("invalid revision number")
	}

	HistoryBlogModel, err := i.toHistory(oldActiveBlog)
	if err != nil {
		return &models.ActiveBlog{}, errors.New("unable to convert blog history model")
	}
	_, err = i.Blog.InsertHistory(tx, HistoryBlogModel)
	if err != nil {
		tx.Rollback()
		return &models.ActiveBlog{}, errors.New("couldn't write a history blog")
	}
	if err := data.CountUpRevisionNumber(oldActiveBlog.GetRevision()); err != nil {
		return &models.ActiveBlog{}, errors.New("unable up number of revision ")
	}
	activeBlog, err := i.Blog.Update(tx, data)
	if err != nil {
		tx.Rollback()
		return &models.ActiveBlog{}, errors.New("couldn't update")
	}
	commitResult := tx.Commit()
	if commitResult.Error != nil {
		tx.Rollback()
		return &models.ActiveBlog{}, errors.New("couldn't commit")
	}

	return activeBlog, nil
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