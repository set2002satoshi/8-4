package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// type Users struct {
// 	gorm.Model
// 	ScreenName  string
// 	DisplayName string
// 	Password    []byte
// 	Email       string
// }

type Users struct {
	ID          int
	ScreenName  string
	DisplayName string
	Password    []byte
	Email       *string
	CreatedAt   int64
	UpdatedAt   int64
}

type UsersForGet struct {
	ID          int     `json:"id"`
	ScreenName  string  `json:"screenName"`
	DisplayName string  `json:"displayName"`
	Email       *string `json:"email"`
}

func (u *Users) BuildForGet() UsersForGet {
	user := UsersForGet{}
	user.ID = u.ID
	user.ScreenName = u.ScreenName
	user.DisplayName = u.DisplayName
	if u.Email != nil {
		user.Email = u.Email
	} else {
		empty := ""
		user.Email = &empty
	}
	return user
}

type UsersForPost struct {
	ScreenName  string `json:"screenName"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func NewUsers(
	ScreenName string,
	DisplayName string,
	Password string,
	Email string,
) (*Users, error) {
	u := &Users{}
	if u.setScreenName(ScreenName) {
		return nil, errors.New("セッターのエラーが出てるよ")
	}
	if u.setDisplayName(DisplayName) {
		return nil, errors.New("セッターのエラーが出てるよ")
	}
	if u.setPassword(Password) {
		return nil, errors.New("セッターのエラーが出てるよ")
	}
	if u.setEmail(Email) {
		return nil, errors.New("セッターのエラーが出てるよ")
	}
	return u, nil

}

func (u *Users) setId(id int) error {
	if id < 0 {
		return errors.New("数字が無効")
	}
	u.ID = id
	return nil
}

func (u *Users) setScreenName(ScreenName string) bool {
	if ScreenName == "" {
		return false
	}
	u.ScreenName = ScreenName
	return true
}

func (u *Users) setDisplayName(DisplayName string) bool {
	if DisplayName == "" {
		return false
	}
	u.DisplayName = DisplayName
	return true
}

func (u *Users) setPassword(Password string) bool {
	pass, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		return false
	}
	u.Password = pass
	return true
}

func (u *Users) setEmail(Email string) bool {
	if Email == "" {
		return false
	}
	u.Email = &Email
	return true
}
