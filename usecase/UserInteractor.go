package usecase

import (
	
	"github.com/set2002satoshi/8-4/domain"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

// func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, err error) {
	db := interactor.DB.Connect()

	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		// return domain.UsersForGet{}, NewResultStatus(404, err)
		return domain.UsersForGet{} , err
	}
	user = foundUser.BuildForGet()
	// return user, NewResultStatus(200, nil)
	return user, nil
}
