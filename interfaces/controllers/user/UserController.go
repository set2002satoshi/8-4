package user

import (
	"github.com/set2002satoshi/8-4/interfaces/database"
	Buser "github.com/set2002satoshi/8-4/interfaces/database/user"
	Uuser "github.com/set2002satoshi/8-4/usecase/user"
)

type UsersController struct {
	Interactor Uuser.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: Uuser.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &Buser.UserRepository{},
		},
	}
}


// func (uc *UsersController) Create(ctx c.Context) {
// 	var userForm UsersForPost
// 	if err := ctx.BindJSON(&userForm); err != nil {
// 		ctx.JSON(400, c.NewH("400", "bindErr"))
// 		return
// 	}
// 	userModel, err := toModel(userForm)
// 	if err != nil {
// 		ctx.JSON(400, c.NewH("400", err.Error()))
// 		return
// 	}
// 	createdUser, err := uc.Interactor.Post(userModel)
// 	if err != nil {
// 		ctx.JSON(400, c.NewH("400", "create err"))
// 		return
// 	}
// 	ctx.JSON(200, c.NewH("201", createdUser))
// }

// func toModel(obj domain.UsersForPost) (*domain.Users, error) {
// 	return domain.NewUsers(
// 		obj.Email,
// 		obj.Password,
// 	)
// }
