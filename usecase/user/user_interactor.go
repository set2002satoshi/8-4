package user

import (
	"errors"
	"time"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
	"github.com/set2002satoshi/8-4/usecase"
)

type UserInteractor struct {
	DB   usecase.DBRepository
	User UserRepository
}


func (i *UserInteractor) FindByID(id int) (user *models.ActiveUser, err error) {
	db := i.DB.Connect()
	foundUser, err := i.User.FindByID(db, id)
	if err != nil {
		// return models.UsersForGet{}, NewResultStatus(404, err)
		return &models.ActiveUser{}, err
	}
	return foundUser, nil
}

func (i *UserInteractor) Post(obj *models.ActiveUser) (*models.ActiveUser, error) {
	db := i.DB.Connect()
	CreatedUser, err := i.User.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return CreatedUser, nil
}

func (i *UserInteractor) FindAll() ([]*models.ActiveUser, error) {
	db := i.DB.Connect()
	users, err := i.User.FindAll(db)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (i *UserInteractor) DeleteByID(id int) (*models.HistoryUser, error) {
	tx := i.DB.Begin()
	ActiveData, err := i.User.FindByID(tx, id)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("find err")
	}
	convertedHistoryUser, err := i.toHistory(ActiveData)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("convert to active table")
	}

	ResultHistory, err := i.User.InsertHistory(tx, convertedHistoryUser)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("insert history err")
	}
	
	if err := i.User.DeleteByID(tx, id); err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("can not delete active data")
	}

	commitResult := tx.Commit()
	if commitResult.Error != nil {
		return &models.HistoryUser{}, errors.New("can not commit")
	}

	return ResultHistory, nil
}

func (i *UserInteractor) Update(data *models.ActiveUser) (*models.ActiveUser, error) {
	tx := i.DB.Begin()
	// 元データを取得
	oldActiveUser, err := i.User.FindByID(tx, int(data.ActiveUserID))

	if oldActiveUser.GetRevision() != data.GetRevision() {
		return &models.ActiveUser{}, errors.New("Invalid revision number")
	}

	if err != nil {
		return &models.ActiveUser{}, err
	}
	// 元データをactiveUserをhistory形式に変換
	HistoryUserModel, err := i.toHistory(oldActiveUser)
	if err != nil {
		return &models.ActiveUser{}, err
	}
	// historyに変換した。データをhistoryデーブルに書き込む
	_, err = i.User.InsertHistory(tx, HistoryUserModel)
	if err != nil {
		tx.Rollback()
		return &models.ActiveUser{}, err
	}
	if err := data.CountUpRevisionNumber(oldActiveUser.GetRevision()); err != nil {
		tx.Rollback()
		return &models.ActiveUser{}, err
	}
	activeUser, err := i.User.Update(tx, data)
	if err != nil {
		tx.Rollback()
		return &models.ActiveUser{}, err
	}
	commitResult := tx.Commit().Error
	if commitResult != nil {
		tx.Rollback()
		return &models.ActiveUser{}, errors.New("can not commit")
	}
	return activeUser, nil
}


func (i *UserInteractor) toHistory(data *models.ActiveUser) (*models.HistoryUser, error) {
	return models.NewHistoryUser(
		temporary.INITIAL_ID,
		int(data.GetID()),
		data.GetName(),
		data.GetEmail(),
		data.GetPassword(),
		time.Time{},
		data.GetUpdatedAt(),
		data.GetCreatedAt(),
		data.GetRevision(),
	)
}
