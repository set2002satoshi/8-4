package controllers

type UsersForPost struct {
	ScreenName  string `json:"screen_name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}
