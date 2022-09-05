package user

import (
	"net/http"

	"github.com/set2002satoshi/8-4/domain"
	c "github.com/set2002satoshi/8-4/interfaces/controllers"
	"github.com/set2002satoshi/8-4/interfaces/controllers/dto"
)

type (
	userCreateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	userCreateResponse struct {
		Message string
		ErrMeg  error
	}
)

func (uc *UsersController) Create(ctx c.Context) {

	var resp userCreateRequest
	if err := ctx.BindJSON(&resp); err != nil {
		response := &userCreateResponse{
			Message: "bindErr",
			ErrMeg:  err,
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

}

func (uc *UsersController) toDTO(u domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:       int(u.ID),
		Name:     u.Name,
		Password: string(u.Password),
	}
}

func (uc *UsersController) toModel(u userCreateRequest) domain.User {
	return &domain.NewUser(
		u.Email,
		u.Name,
		u.Password,
	)
}
