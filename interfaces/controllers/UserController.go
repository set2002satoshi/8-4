package controllers

import (
	"strconv"

	"github.com/set2002satoshi/8-4/interfaces/database"
	"github.com/set2002satoshi/8-4/usecase"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor {
			// databaseからDB->(Begin, Connect)を配下に持つを代入
			DB: &database.DBRepository{ DB:db },
			User: &database.UserRepository{},
		},
	}
}

func (controller *UsersController) Get(c Context) {

	id, _ := strconv.Atoi(c.Param("id"));

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		// c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		c.JSON(404, NewH(res.Error(), nil))
		return
	}
	c.JSON(200, NewH("success", user))
}

