package response

import (
	"time"
)

// user response data
type (
	ActiveUserEntity struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Option   Options
	}

	HistoryUserEntity struct {
		ID           int    `json:"id"`
		ActiveUserID int    `json:"active_user_id"`
		Name         string `json:"name"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		Option       Options
	}

	Options struct {
		Revision  int       `json:"revision"`
		CratedAt  time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

type (
	FindByIDUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr error  `json:"error"`
		MsgErr  string `json:"msg"`
	}

	CreateUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr error `json:"code"`
		MsgErr  string `json:"msg"`
	}

	DeleteUserResponse struct {
		Result *HistoryUserResult `json:"results"`

		CodeErr error `json:"code"`
		MsgErr  string `json:"msg"`
	}

	FindAllUserResponse struct {
		Results *ActiveUserResults `json:"results"`

		CodeErr error `json:"code"`
		MsgErr  string `json:"msg"`
	}

	UpdateUserResponse struct {
		Result *ActiveUserResult `json:"results"`

		CodeErr error `json:"code"`
		MsgErr  string `json:"msg"`
	}
)

type (
	ActiveUserResult struct {
		User *ActiveUserEntity `json:"user"`
	}
	ActiveUserResults struct {
		Users []*ActiveUserEntity `json:"user"`
	}

	HistoryUserResult struct {
		User *HistoryUserEntity `json:"user"`
	}
	HistoryUserResults struct {
		Users []*HistoryUserEntity `json:"user"`
	}
)
