package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/set2002satoshi/8-4/models"
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
)

type (
	userCreateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	userCreateResponse struct {
		Message  string
		ErrMeg   error
		Response *models.User
	}
)

func (uc *UsersController) Create(ctx c.Context) {

	var req userCreateRequest
	if err := ctx.BindJSON(&req); err != nil {
		response := userCreateResponse{
			Message: "bindErr",
			ErrMeg:  err,
		}
		ctx.JSON(http.StatusBadRequest, response)
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
	fmt.Println(reqModel)
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

func toModel(req userCreateRequest) (*models.User, error) {
	return models.NewUser(
		int(0),
		req.Name,
		req.Email,
		req.Password,
		time.Time{},
		time.Time{},
	)

}
