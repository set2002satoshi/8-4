package user

import (
	// "time"

	"github.com/set2002satoshi/8-4/interfaces/controllers/dto"
	"github.com/set2002satoshi/8-4/interfaces/database"
	DBuser "github.com/set2002satoshi/8-4/interfaces/database/user"
	"github.com/set2002satoshi/8-4/models"
	usecase "github.com/set2002satoshi/8-4/usecase/user"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &DBuser.UserRepository{},
		},
	}
}

func (uc *UsersController) toDTO(u models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       int(u.GetID()),
		Name:     u.GetName(),
		Password: string(u.GetPassword()),
	}
}




// func (uc *UsersController) Create(ctx c.Context) {
// 	var userForm UsersForPost
// 	if err := ctx.BindJSON(&userForm); err != nil {
// 		ctx.JSON(400, c.NewH("400", "bindErr"))
// 		return
// 	}
// 	userModels, err := toModels(userForm)
// 	if err != nil {
// 		ctx.JSON(400, c.NewH("400", err.Error()))
// 		return
// 	}
// 	createdUser, err := uc.Interactor.Post(userModels)
// 	if err != nil {
// 		ctx.JSON(400, c.NewH("400", "create err"))
// 		return
// 	}
// 	ctx.JSON(200, c.NewH("201", createdUser))
// }

// func toModels(obj models.UsersForPost) (*models.Users, error) {
// 	return models.NewUsers(
// 		obj.Email,
// 		obj.Password,
// 	)
// }
