package user

import (
	
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)

type (
	FindAllUserResponse struct {
		response.FindAllUserResponse
	}
)

func (r FindAllUserResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}


func (uc *UsersController) FindAll(ctx c.Context) {

	res := &FindAllUserResponse{}

	user, err := uc.Interactor.FindAll()
	if err != nil {
		res.SetErr(err, "")
		ctx.JSON(500, res)
		return
	}
	res.Results = &response.ActiveUserResults{Users: uc.convertActivesToDTOs(user)}
	ctx.JSON(200, c.NewH("ok", res))

}