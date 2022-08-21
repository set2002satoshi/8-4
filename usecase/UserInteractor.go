package usecase

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

// func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
// func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, err error) {
// 	db := interactor.DB.Connect()

// 	foundUser, err := interactor.User.FindByID(db, id)
// 	if err != nil {
// 		return domain.UsersForGet{}, NewResultStatus(404, err)
// 		return domain.UsersForGet{} , err
// 	}
// 	user = foundUser.BuildForGet()
// 	return user, NewResultStatus(200, nil)
// 	return user, nil
// }

// func (interactor *UserInteractor) Post(obj domain.UsersForPost) (user domain.Users, err error) {
// 	db := interactor.DB.Connect()
// 	BuildUser := obj.BuildForUser()
// 	CreatedUser, err := interactor.User.CreateUser(db, BuildUser)
// 	if err != nil {
// 		return domain.Users{}, err
// 	}
// 	return CreatedUser, nil

// }
