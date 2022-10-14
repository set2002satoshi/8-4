package user

import (
	"net/http"
	"time"
	"fmt"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/models"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
	"github.com/set2002satoshi/8-4/pkg/module/dto/response"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type (
	UpdateUserResponse struct {
		response.UpdateUserResponse
	}
)

func (r UpdateUserResponse) SetErr(err error, errMsg string) {
	r.CodeErr = err
	r.MsgErr = errMsg
}

func (uc *UsersController) Update(ctx c.Context) {

	req := request.UserUpdateRequest{}
	res := &UpdateUserResponse{}

	if err := ctx.BindJSON(&req); err != nil {
		res.SetErr(err, "BindErr")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	reqModel, err := uc.toModel(req)
	fmt.Println(reqModel)
	if err != nil {
		res.SetErr(err, "toModelErr")
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}
	activeUser, err := uc.Interactor.Update(reqModel)
	if err != nil {
		res.SetErr(err, "updateErr")
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Result = &response.ActiveUserResult{User: uc.convertActiveToDTO(activeUser)}
	ctx.JSON(http.StatusOK, res)

}

func (uc *UsersController) toModel(req request.UserUpdateRequest) (*models.ActiveUser, error) {
	return models.NewActiveUser(
		int(req.ID),
		req.Name,
		req.Email,
		req.Password,
		time.Now(),
		time.Time{},
		temporary.REVISION(req.Revision),
	)
}
