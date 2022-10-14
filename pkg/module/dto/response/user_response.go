package response

import (
	"time"
)




// user response data
type (
	UserEntity struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Option   Options
	}

	Options struct {
		Revision  int       `json:"revision"`
		CratedAt  time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)


type (
	FindByIDResponse struct {
		Result *UserResult `json:"results"`
	
		CodeErr string `json:"code"`
		MsgErr string `json:"msg"`
	}

	CreateUserResponse struct {
		Result *UserResult `json:"results"`
	
		CodeErr string `json:"code"`
		MsgErr string `json:"msg"`
	}
	
	DeleteUserResponse struct {
		Result *UserResult `json:"results"`
	
		CodeErr string `json:"code"`
		MsgErr string `json:"msg"`
	}

	FindAllUserResponse struct {
		Results *UserResults `json:"results"`
	
		CodeErr string `json:"code"`
		MsgErr string `json:"msg"`
	}

	UpdateUserResponse struct {
		Result *UserResult `json:"results"`
	
		CodeErr string `json:"code"`
		MsgErr string `json:"msg"`
	}
)






type (
	UserResult struct {
		User UserEntity `json:"user"`
	}
	UserResults struct {
		Users []UserEntity `json:"user"`
	}
)

