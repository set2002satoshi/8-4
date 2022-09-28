package user

import (

	// "github.com/set2002satoshi/8-4/interfaces/controllers"
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