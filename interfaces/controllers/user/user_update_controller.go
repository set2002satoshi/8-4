package user

import (

	"net/http"
	"time"

	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/models"
)

type (
	userUpdateRequest struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	userUpdateResponse struct {
		Message  string
		ErrMeg   error
		Response *models.ActiveUser
	}
)

func (uc *UsersController) Update(ctx c.Context) {

	var req userUpdateRequest
	if err := ctx.BindJSON(&req); err != nil {
		response := &userUpdateResponse{
			Message: "BindErr",
			ErrMeg:  err,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	reqModel, err := uc.toModel(req)
	if err != nil {
		response := &userUpdateResponse{
			Message: "toModelErr",
			ErrMeg:  err,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	activeUser, err := uc.Interactor.Update(reqModel)
	if err != nil {
		response := &userUpdateResponse{
			Message:  "updateErr",
			ErrMeg:   err,
			Response: activeUser,
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	response := &userUpdateResponse{
		Message:  "ok",
		ErrMeg:   err,
		Response: activeUser,
	}
	ctx.JSON(http.StatusOK, response)
}

func (uc *UsersController) toModel(req userUpdateRequest) (*models.ActiveUser, error) {
	return models.NewActiveUser(
		int(req.ID),
		req.Name,
		req.Email,
		req.Password,
		time.Now(),
		time.Time{},
	)
}
