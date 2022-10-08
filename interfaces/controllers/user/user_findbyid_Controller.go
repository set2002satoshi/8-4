package user

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	// "github.com/set2002satoshi/8-4/interfaces/database"
)

func (uc *UsersController) FindByID(ctx c.Context) {
	
	req := request.UserFindByIDRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(404, c.NewH("bindErr", nil))
		return
	}
	user, res := uc.Interactor.FindByID(req.ID)
	if res != nil {
		ctx.JSON(404, c.NewH(res.Error(), nil))
		return
	}
	ctx.JSON(200, c.NewH("success", uc.toDTO(user)))
}
