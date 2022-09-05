package user

import (

	// "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/domain"
	"github.com/set2002satoshi/8-4/usecase"
)

type UserInteractor struct {
	DB   usecase.DBRepository
	User UserRepository
}

// func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
func (interactor *UserInteractor) Find(id int) (user domain.User, err error) {
	db := interactor.DB.Connect()
	foundUser, err := interactor.User.Find(db, id)
	if err != nil {
		// return domain.UsersForGet{}, NewResultStatus(404, err)
		return domain.User{}, err
	}
	return foundUser, nil
}

// func (interactor *UserInteractor) Post(obj *domain.Users) (*domain.Users, error) {
// 	db := interactor.DB.Connect()
// 	// UserModel, err := toModel(obj)
// 	// if err != nil {
// 	// 	return &domain.Users{}, nil
// 	// }
// 	CreatedUser, err := interactor.User.CreateUser(db, obj)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CreatedUser, nil
// }
