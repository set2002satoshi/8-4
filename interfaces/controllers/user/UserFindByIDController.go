package user

import (
	"strconv"

	"github.com/set2002satoshi/8-4/domain"
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/interfaces/controllers/dto"
	// "github.com/set2002satoshi/8-4/interfaces/database"
)

func (uc *UsersController) FindByID(ctx c.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	user, res := uc.Interactor.Find(id)
	if res != nil {
		// c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		ctx.JSON(404, c.NewH(res.Error(), nil))
		return
	}
	ctx.JSON(200, c.NewH("success", uc.toDTO(user)))
}

func (uc *UsersController) toDTO(u domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       int(u.ID),
		Name:     u.Name,
		Password: string(u.Password),
	}
}


