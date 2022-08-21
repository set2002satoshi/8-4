package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	ID          int
	ScreenName  string
	DisplayName string
	Password    []byte
	Email       string
	Created     int64
	UpdatedAt   int64
}

type UsersForPost struct {
	ScreenName  string  `json:"screenName"`
	DisplayName string  `json:"displayName"`
	Email       string `json:"email"`
	Password    string  `json:"password"`
}





func (u *Users) setId(id int) error {
	if id < 0 {
		return errors.New("数字が無効")
	}
	u.ID = id
	return nil
}

func (u *Users) setScreenName(ScreenName string) error {
	if ScreenName == "" {
		return errors.New("screenNameが定義されていません")
	}
	u.ScreenName = ScreenName
	return nil
}

func (u *Users) setDisplayName(DisplayName string) error {
	if DisplayName == "" {
		return errors.New("DisplayNameが定義されていません")
	}
	u.DisplayName = DisplayName
	return nil
}

func (u *Users) setPassword(Password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		return errors.New("Passwordがえらってる")
	}
	u.Password = pass
	return nil
}

func (u *Users) setEmail(Email string) error {
	if Email == "" {
		return errors.New("Emailが定義されていません")
	}
	u.Email = Email
	return nil
}

