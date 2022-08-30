package controllers



type UsersForPost struct {
	ScreenName  string  `json:"screenName"`
	DisplayName string  `json:"displayName"`
	Email       string `json:"email"`
	Password    string  `json:"password"`
}
