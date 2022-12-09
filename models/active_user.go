package models

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"
	cErr "github.com/set2002satoshi/8-4/pkg/module/customs/errors"
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
	au := &ActiveUser{}

	var err error
	err = cErr.Combine(err, au.setID(id))
	err = cErr.Combine(err, au.setName(name))
	err = cErr.Combine(err, au.setEmail(email))
	err = cErr.Combine(err, au.setPassword(password))
	err = cErr.Combine(err, au.setCreatedAt(createdAt))
	err = cErr.Combine(err, au.setUpdatedAt(updatedAt))
	err = cErr.Combine(err, au.setRevision(revision))

	return au, err
}

func (u *ActiveUser) setID(id int) error {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	u.ActiveUserID = i
	return nil
}

func (u *ActiveUser) setEmail(email string) error {
	u.Email = email
	fmt.Println(u.Email)
	return nil
}

func (u *ActiveUser) setName(name string) error {
	u.Name = name
	fmt.Println(u.Name)
	return nil
}

func (u *ActiveUser) setPassword(password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = []byte(pass)
	return nil
	// u.Password = []byte(password)
}

func (s *ActiveUser) setCreatedAt(createdAt time.Time) error {
	s.CreatedAt = createdAt
	return nil
}

func (s *ActiveUser) setUpdatedAt(updatedAt time.Time) error {
	s.CreatedAt = updatedAt
	return nil
}

func (s *ActiveUser) setRevision(revision temporary.REVISION) error {
	s.Revision = revision
	return nil
}
func (s *ActiveUser) CountUpRevisionNumber(num temporary.REVISION) error {

	if s.GetRevision() != num {
		return errors.New("改定番号が異なるため更新はできません")
	}
	if err := s.setRevision(num + 1); err != nil {
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

func (u *ActiveUser) GetBlogs() []ActiveBlog {
	return u.Blogs
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



func (s *ActiveUser) ChangeUserName(name string) error {
	blogs := s.GetBlogs()
	for _, blog := range blogs {
		if err := blog.setName(name); err != nil {
			return errors.New("blogの名前変更でえらーが出てますよ")
		}
	}
	return nil
}