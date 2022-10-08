package user

import (
	"net/http"
	"time"

	"github.com/set2002satoshi/8-4/models"
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/pkg/module/dto/request"
)

type (
	userCreateResponse struct {
		Message  string
		ErrMeg   error
		Response *models.ActiveUser
	}
)

func (uc *UsersController) Create(ctx c.Context) {

	req := request.UserCreateRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(404, c.NewH("bindErr", nil))
		return
	}
	// バリデーションはめんどいから描かない
	reqModel, err := toModel(req)
	if err != nil {
		response := userCreateResponse{
			Message: "toModelsErr",
			ErrMeg:  err,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	createdUser, err := uc.Interactor.Post(reqModel)
	if err != nil {
		response := userCreateResponse{
			Message: "createdUser err",
			ErrMeg:  err,
		}
		ctx.JSON(500, response)
		return
	}
	response := userCreateResponse{
		Message:  "ok",
		ErrMeg:   nil,
		Response: createdUser,
	}
	ctx.JSON(201, response)

}

func toModel(req request.UserCreateRequest) (*models.ActiveUser, error) {
	return models.NewActiveUser(
		int(0),
		req.Name,
		req.Email,
		req.Password,
		time.Now(),
		time.Time{},
	)
}


