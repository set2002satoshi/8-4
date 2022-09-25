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
func (interactor *UserInteractor) Find(id int) (user models.User, err error) {
	db := interactor.DB.Connect()
	foundUser, err := interactor.User.Find(db, id)
	if err != nil {
		// return models.UsersForGet{}, NewResultStatus(404, err)
		return models.User{}, err
	}
	return foundUser, nil
}

func (interactor *UserInteractor) Post(obj *models.User) (*models.User, error) {
	db := interactor.DB.Connect()
	// UserModel, err := toModel(obj)
	// if err != nil {
	// 	return &models.Users{}, nil
	// }
	CreatedUser, err := interactor.User.Create(db, obj)
	if err != nil {
		return nil, err
	}
	return &CreatedUser, nil
}
