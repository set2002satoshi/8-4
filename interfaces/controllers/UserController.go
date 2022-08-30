package controllers

import (
	"strconv"

	"github.com/set2002satoshi/8-4/domain"
	"github.com/set2002satoshi/8-4/interfaces/database"
	"github.com/set2002satoshi/8-4/usecase"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (uc *UsersController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := uc.Interactor.Get(id)
	if res != nil {
		// c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		c.JSON(404, NewH(res.Error(), nil))
		return
	}
	c.JSON(200, NewH("success", user))
}

func (uc *UsersController) Create(c Context) {
	var userForm UsersForPost
	if err := c.BindJSON(&userForm); err != nil {
		c.JSON(400, NewH("400", nil))
		return
	}
	userModel, err := toModel(userForm)
	if err != nil {
		c.JSON(400, NewH("400", nil))
		return
	}
	createdUser, res := uc.Interactor.CreateUser(userModel)
	if res != nil {
		c.JSON(400, NewH("400", res))
		return
	}
	c.JSON(200, NewH("201", createdUser))
}

func toModel(obj UsersForPost) (*domain.Users, error) {
	return domain.NewUsers(
		obj.ScreenName,
		obj.DisplayName,
		obj.Password,
		obj.Email,
	)
}
