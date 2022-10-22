package user

import (
	"net/http"
	"time"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type (
	CreateUserResponse struct {
		response.CreateUserResponse
	}
)

func (r *CreateUserResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}

func (uc *UsersController) Create(ctx c.Context) {

	req := request.UserCreateRequest{}
	res := &CreateUserResponse{}
	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "bindErr")
		ctx.JSON(404, res)
		return
	}
	// バリデーションはめんどいから描かない
	reqModel, err := toModel(req)
	if err != nil {
		res.SetErr(err, "toModelsErr")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	createdUser, err := uc.Interactor.Post(reqModel)
	if err != nil {
		res.SetErr(err, "createdUser err")
		ctx.JSON(500, res)
		return
	}
	res.Result = &response.ActiveUserResult{User: uc.convertActiveToDTO(createdUser)}
	ctx.JSON(201, c.NewH("ok", res))


}

func toModel(req request.UserCreateRequest) (*models.ActiveUser, error) {
	return models.NewActiveUser(
		int(0),
		req.Name,
		req.Email,
		req.Password,
		time.Time{},
		time.Time{},
		temporary.INITIAL_REVISION,
	)
}
