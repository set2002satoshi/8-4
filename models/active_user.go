package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type ActiveUser struct {
	ActiveUserID temporary.IDENTIFICATION `gorm:"primaryKey"`
	Name         string                   `gorm:"not null;size:16"`
	Email        string                   `gorm:"unique;not null"`
	Password     []byte                   `gorm:"not null"`
	Blogs        []ActiveBlog             `gorm:"foreignKey:ActiveUserID; constraint:OnUpdate:CASCADE, OnDelete:CASCADE;"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revision     temporary.REVISION
}

func NewActiveUser(
	id int,
	name string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
	revision temporary.REVISION,
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
		return nil, errors.New("setCreatedAtにエラーが発生しています。")
	}
	if u.setUpdatedAt(updatedAt) {
		return nil, errors.New("setUpdatedAtにエラーが発生しています。")
	}

	if u.setRevision(revision) {
		return nil, errors.New("setRevisionにエラーが発生しています。")
	}

	return u, nil
}

func (u *ActiveUser) setID(id int) bool {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return true
	}
	u.ActiveUserID = i
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
	return false
	// u.Password = []byte(password)
}

func (s *ActiveUser) setCreatedAt(createdAt time.Time) bool {
	s.CreatedAt = createdAt
	return false
}

func (s *ActiveUser) setUpdatedAt(updatedAt time.Time) bool {
	s.CreatedAt = updatedAt
	return false
}

func (s *ActiveUser) setRevision(revision temporary.REVISION) bool {
	s.Revision = revision
	return false
}
func (s *ActiveUser) CountUpRevisionNumber(num temporary.REVISION) error {

	if s.GetRevision() != num {
		return errors.New("改定番号が異なるため更新はできません")
	}
	if ok := s.setRevision(num + 1); ok {
		return errors.New("Invalid setting")
	}
	return nil
}

func (u *ActiveUser) GetID() int {
	return int(u.ActiveUserID)
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

func (s *ActiveUser) GetCreatedAt() time.Time {
	return s.CreatedAt

}

func (s *ActiveUser) GetUpdatedAt() time.Time {
	return s.CreatedAt
}

func (s *ActiveUser) GetRevision() temporary.REVISION {
	return s.Revision

}
