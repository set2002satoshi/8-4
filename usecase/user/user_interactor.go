package user

import (
	"errors"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/usecase"
)

type UserInteractor struct {
	DB   usecase.DBRepository
	User UserRepository
}

// func (interactor *UserInteractor) Get(id int) (user models.UsersForGet, resultStatus *ResultStatus) {
func (interactor *UserInteractor) FindByID(id int) (user models.ActiveUser, err error) {
	db := interactor.DB.Connect()
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		// return models.UsersForGet{}, NewResultStatus(404, err)
		return models.ActiveUser{}, err
	}
	return foundUser, nil
}

func (interactor *UserInteractor) Post(obj *models.ActiveUser) (*models.ActiveUser, error) {
	db := interactor.DB.Connect()
	// UserModel, err := toModel(obj)
	// if err != nil {
	// 	return &models.ActiveUsers{}, nil
	// }
	CreatedUser, err := interactor.User.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return &CreatedUser, nil
}

func (interactor *UserInteractor) FindAll() (*[]models.ActiveUser, error) {
	db := interactor.DB.Connect()
	users, err := interactor.User.FindAll(db)
	if err != nil {
		return nil, err
	}
	return &users, nil

}

func (interactor *UserInteractor) DeleteByID(id int) (*models.HistoryUser, error) {
	// idからactiveテーブルを参照する
	// toHistoryを行なってヒストリーテーブルに書き込み
	// idを元に削除
	// Historyモデルを返す。
	tx := interactor.DB.Begin()
	ActiveData, err := interactor.User.FindByID(tx, id)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("find err")
	}
	convertedHistoryUser, err := toHistory(&ActiveData)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("convert to active table")
	}

	ResultHistory, err := interactor.User.InsertHistory(tx, convertedHistoryUser)
	if err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("insert history err")
	}
	
	if err := interactor.User.DeleteByID(tx, id); err != nil {
		tx.Rollback()
		return &models.HistoryUser{}, errors.New("can not delete active data")
	}

	commitResult := tx.Commit()
	if commitResult.Error != nil {
		return &models.HistoryUser{}, errors.New("can not commit")
	}

	return ResultHistory, nil
}

func toHistory(data *models.ActiveUser) (*models.HistoryUser, error) {
	return models.NewHistoryUser(
		int(0),
		int(data.GetID()),
		data.GetName(),
		data.GetEmail(),
		data.GetPassword(),
		data.GetCreatedAt(),
		data.GetUpdatedAt(),
	)
}
