package user

import (
	"strconv"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	
)

func (uc *UsersController) DeleteByID(ctx c.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	movedByResult, res := uc.Interactor.DeleteByID(id)
	if res != nil {
		ctx.JSON(404, c.NewH(res.Error(), nil))
		return 
	}
	ctx.JSON(200, c.NewH("success", movedByResult))
}