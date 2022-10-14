package user

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	// "github.com/set2002satoshi/8-4/interfaces/database"
)

type (
	FindByIdUserResponse struct {
		response.FindByIDUserResponse
	}
)

func (r FindByIdUserResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (uc *UsersController) FindByID(ctx c.Context) {
	
	req := request.UserFindByIDRequest{}
	res := FindByIdUserResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "bindErr")
		ctx.JSON(404, res)
		return
	}
	user, err := uc.Interactor.FindByID(req.ID)
	if err != nil {
		res.SetErr(err, "")
		ctx.JSON(404, res)
		return
	}
	res.Result = &response.ActiveUserResult{User: uc.convertActiveToDTO(user)}
	ctx.JSON(200, c.NewH("success", res))
}
