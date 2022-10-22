package blog

import (
	"strconv"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
)

func (bc *BlogsController) FindByID(ctx c.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	blog, err := bc.Interactor.FindByID(id)
	if err != nil {
		ctx.JSON(404, c.NewH(err.Error(), nil))
		return
	}
	ctx.JSON(200, c.NewH("success", blog))
}
