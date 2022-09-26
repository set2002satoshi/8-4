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

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password []byte
}

func NewUser(
	id int,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (*User, error) {
	u := &User{}
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

func (u *User) setID(id int) bool {
	if id < 0 {
		return false
	}
	u.Model.ID = uint(id)
	return false
}

func (u *User) setEmail(email string) bool {
	u.Email = email
	fmt.Println(u.Email)
	return false
}

func (u *User) setName(name string) bool {
	u.Name = name
	fmt.Println(u.Name)
	return false
}

func (u *User) setPassword(password string) bool {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return true
	}
	u.Password = []byte(pass)
	fmt.Println(u.Password)
	return false
	// u.Password = []byte(password)
}

func (u *User) setCreatedAt(createdAt time.Time) bool {
	u.Model.CreatedAt = createdAt
	return false
}

func (u *User) setUpdatedAt(updatedAt time.Time) bool {
	u.Model.UpdatedAt = updatedAt
	return false
}


func (u *User) GetID() int {
	return int(u.Model.ID)
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return string(u.Password)
}


func (u *User) GetCreatedAt() time.Time {
	return u.Model.CreatedAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.Model.UpdatedAt
}

