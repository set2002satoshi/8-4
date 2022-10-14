package user

import (
	
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
)

func (uc *UsersController) FindAll(ctx c.Context) {

	user, err := uc.Interactor.FindAll()
	if err != nil {
		ctx.JSON(500, c.NewH(err.Error(), nil))
		return
	}
	
	ctx.JSON(200, c.NewH("ok", uc.convertActivesToDTOs(user)))

	return
}