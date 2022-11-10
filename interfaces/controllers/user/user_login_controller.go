package user

import (
	"fmt"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"

	"github.com/set2002satoshi/8-4/pkg/module/dto/auth"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
)

type (
	LoginUserResponse struct {
		response.LoginUserResponse
	}
)

func (r *LoginUserResponse) setErr(err error, errMsg string) {ß
	r.CodeErr = err
	r.MsgErr = errMsg
}

func (uc *UsersController) Login(ctx c.Context) {
	req := request.UserLoginRequest{}
	res := &LoginUserResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		res.setErr(err, "bindErr")
		ctx.JSON(404, res)
		return
	}
	loginModel, err := toLoginModel(req)
	if err != nil {
		return
	}
	token, err := uc.Interactor.FetchToken(loginModel)
	if err != nil {
		fmt.Println(err)
		res.setErr(err, "loginErr")
		ctx.JSON(404, c.NewH("ok", res))
		return
	}
	res.Result = &response.LoginUserResult{Token: token}
	ctx.JSON(200, res)

}

func toLoginModel(req request.UserLoginRequest) (*auth.UserLoginModel, error) {
	return &auth.UserLoginModel{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}
