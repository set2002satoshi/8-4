package domain

import (
	"errors"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// type Users struct {
// 	gorm.Model
// 	ScreenName  string
// 	DisplayName string
// 	Password    []byte
// 	Email       string
// }

// type User struct {}
type User struct {
	gorm.Model
	Name     string
	Email    *string
	Password []byte
}

// type UserForGet struct {}
type UsersForGet struct {
	ID          int     `json:"id"`
	ScreenName  string  `json:"screenName"`
	DisplayName string  `json:"displayName"`
	Email       *string `json:"email"`
}

// type UserForPost struct {}

// type UsersForPost struct {
// 	Name     string `json:"displayName"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

type UsersForPost struct {
	Name     string `json:"displayName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func NewUser
func NewUser(
	Email string,
	Name string,
	Password string,
) (*User, error) {
	u := &User{}
	if u.setName(Name) {
		return nil, errors.New("ScreenNameセッターのエラーが出てるよ")
	}
	if u.setEmail(Email) {
		return nil, errors.New("Emailセッターのエラーが出てるよ")
	}
	if u.setPassword(Password) {
		return nil, errors.New("Passwordセッターのエラーが出てるよ")
	}

	return u, nil

}

// func (u *Users) setId(id int) error {
// 	if id < 0 {
// 		return errors.New("数字が無効")
// 	}
// 	u.ID = id
// 	return nil
// }

func (u *User) setName(Name string) bool {
	if Name != "" {
		return false
	}
	u.Name = Name
	return true
}

func (u *User) setPassword(Password string) bool {
	// pass, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	// if err != nil {
	// 	return false
	// }
	// u.Password = []byte(pass)
	// return true
	u.Password = []byte(Password)
	return true
}

func (u *User) setEmail(Email string) bool {
	if Email != "" {
		return false
	}
	u.Email = &Email
	return true
}

func (u *User) GetID() int {
	return int(u.ID)
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return *u.Email
}

func (u *User) GetPassword() string {
	return string(u.Password)
}
