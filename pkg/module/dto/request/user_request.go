package request

type (
	UserFindByIDRequest struct {
		ID int `json:"id"`
	}
	UserDeleteRequest struct {
		ID int `json:"id"`
	}
	UserCreateRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	UserUpdateRequest struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
		Revision uint   `json:"revision"`
	}
)

type (
	UserLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
