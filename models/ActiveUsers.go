package models

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

// type User struct {
// 	ID       uint
// 	Name     string
// 	Email    string
// 	Password []byte
// 	option   *Options
// }

type ActiveUser struct {
	gorm.Model
	Name     string
	Email    string
	Password []byte
}

func NewActiveUser(
	id int,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (*ActiveUser, error) {
	u := &ActiveUser{}

	if u.setID(id) {
		return nil, errors.New("idセッターのエラーが出てるよ")
	}

	if u.setName(name) {
		return nil, errors.New("Nameセッターのエラーが出てるよ")
	}
	if u.setEmail(email) {
		return nil, errors.New("ScreenNameセッターのエラーが出てるよ")
	}
	if u.setPassword(password) {
		return nil, errors.New("ScreenNameセッターのエラーが出てるよ")
	}

	if u.setCreatedAt(createdAt) {
		return nil, errors.New("createdAtセッターのエラーが出てるよ")
	}

	if u.setUpdatedAt(createdAt) {
		return nil, errors.New("createdAtセッターのエラーが出てるよ")
	}

	return u, nil
}

func (u *ActiveUser) setID(id int) bool {
	if id < 0 {
		return false
	}
	u.Model.ID = uint(id)
	return false
}

func (u *ActiveUser) setEmail(email string) bool {
	u.Email = email
	fmt.Println(u.Email)
	return false
}

func (u *ActiveUser) setName(name string) bool {
	u.Name = name
	fmt.Println(u.Name)
	return false
}

func (u *ActiveUser) setPassword(password string) bool {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return true
	}
	u.Password = []byte(pass)
	fmt.Println(u.Password)
	return false
	// u.Password = []byte(password)
}

func (u *ActiveUser) setCreatedAt(createdAt time.Time) bool {
	u.Model.CreatedAt = createdAt
	return false
}

func (u *ActiveUser) setUpdatedAt(updatedAt time.Time) bool {
	u.Model.UpdatedAt = updatedAt
	return false
}


func (u *ActiveUser) GetID() int {
	return int(u.Model.ID)
}

func (u *ActiveUser) GetName() string {
	return u.Name
}

func (u *ActiveUser) GetEmail() string {
	return u.Email
}

func (u *ActiveUser) GetPassword() string {
	return string(u.Password)
}


func (u *ActiveUser) GetCreatedAt() time.Time {
	return u.Model.CreatedAt
}

func (u *ActiveUser) GetUpdatedAt() time.Time {
	return u.Model.UpdatedAt
}

