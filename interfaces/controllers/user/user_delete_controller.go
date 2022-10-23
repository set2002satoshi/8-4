package user

import (

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)

type (
	DeleteUserResponse struct {
		response.DeleteUserResponse
	}
)

func (r *DeleteUserResponse) SetErr(err error, ErrMsg string) {
	r.CodeErr = err
	r.MsgErr = ErrMsg
}


func (uc *UsersController) Delete(ctx c.Context) {

	req := request.UserDeleteRequest{}
	res := &DeleteUserResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "bindErr")
		ctx.JSON(404, res)
		return
	}

	movedByResult, err := uc.Interactor.DeleteByID(req.ID)
	if err != nil {
		res.SetErr(err, "")
		ctx.JSON(404, res)
		return
	}
	res.Result = &response.HistoryUserResult{User: uc.convertHistoryToDTO(movedByResult)}
	ctx.JSON(200, c.NewH("success",res))
}