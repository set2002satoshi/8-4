package user

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	
)

func (uc *UsersController) DeleteByID(ctx c.Context) {

	req := request.UserDeleteRequest{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(404, c.NewH("bindErr", nil))
		return
	}

	movedByResult, res := uc.Interactor.DeleteByID(req.ID)
	if res != nil {
		ctx.JSON(404, c.NewH(res.Error(), nil))
		return
	}
	ctx.JSON(200, c.NewH("success", uc.convertHistoryToDTO(movedByResult)))
}