package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type HistoryUser struct {
	gorm.Model
	activeID int
	name     string
	email    string
	Password []byte
}

func NewHistoryUser(
	id int,
	activeID int,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (*HistoryUser, error) {
	u := &HistoryUser{}

	if u.setID(id) {
		return nil, errors.New("idセッターのエラーが出てるよ")
	}

	if u.setActiveID(activeID) {
		return nil, errors.New("activeIDセッターのエラーが出てるよ")
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

func (u *HistoryUser) setID(id int) bool {
	if id < 0 {
		return false
	}
	u.Model.ID = uint(id)
	return false
}

func (u *HistoryUser) setActiveID(activeID int) bool {
	u.activeID = activeID
	return false
}

func (u *HistoryUser) setEmail(email string) bool {
	u.email = email
	fmt.Println(u.email)
	return false
}

func (u *HistoryUser) setName(name string) bool {
	u.name = name
	fmt.Println(u.name)
	return false
}

func (u *HistoryUser) setPassword(password string) bool {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return true
	}
	u.Password = []byte(pass)
	fmt.Println(u.Password)
	return false
	// u.Password = []byte(password)
}

func (u *HistoryUser) setCreatedAt(createdAt time.Time) bool {
	u.Model.CreatedAt = createdAt
	return false
}

func (u *HistoryUser) setUpdatedAt(updatedAt time.Time) bool {
	u.Model.UpdatedAt = updatedAt
	return false
}

func (u *HistoryUser) GetID() int {
	return int(u.Model.ID)
}

func (u *HistoryUser) GetName() string {
	return u.name
}

func (u *HistoryUser) GetEmail() string {
	return u.email
}

func (u *HistoryUser) GetPassword() string {
	return string(u.Password)
}

func (u *HistoryUser) GetCreatedAt() time.Time {
	return u.Model.CreatedAt
}

func (u *HistoryUser) GetUpdatedAt() time.Time {
	return u.Model.UpdatedAt
}
