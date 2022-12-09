package models

import (
	"fmt"
	"time"

	cErr "github.com/set2002satoshi/8-4/pkg/module/customs/errors"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type HistoryUser struct {
	HistoryUserID temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveUserID  int
	Name          string
	Email         string
	Password      []byte
	Blogs         []HistoryBlog
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ActiveTime    time.Time
	Revision      temporary.REVISION
}

func NewHistoryUser(
	id int,
	activeID int,
	name string,
	email string,
	password string,
	blogs []ActiveBlog,
	createdAt time.Time, // このは空のtime
	updatedAt time.Time, // time.Time
	activeTime time.Time, // ここに時間を入れる
	revision temporary.REVISION,
) (*HistoryUser, error) {
	hu := &HistoryUser{}

	var err error
	err = cErr.Combine(err, hu.setID(id))
	err = cErr.Combine(err, hu.setActiveID(activeID))
	err = cErr.Combine(err, hu.setName(name))
	err = cErr.Combine(err, hu.setEmail(email))
	err = cErr.Combine(err, hu.setPassword(password))
	err = cErr.Combine(err, hu.setBlogs(blogs))
	err = cErr.Combine(err, hu.setCreatedAt(createdAt))
	err = cErr.Combine(err, hu.setUpdatedAt(updatedAt))
	err = cErr.Combine(err, hu.setActiveCreateAt(activeTime))
	err = cErr.Combine(err, hu.setRevision(revision))

	return hu, err
}

func (u *HistoryUser) setID(id int) error {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	u.HistoryUserID = i
	return nil
}

func (u *HistoryUser) setActiveID(activeID int) error {
	u.ActiveUserID = activeID
	return nil
}

func (u *HistoryUser) setEmail(email string) error {
	u.Email = email
	fmt.Println(u.Email)
	return nil
}

func (u *HistoryUser) setName(name string) error {
	u.Name = name
	fmt.Println(u.Name)
	return nil
}

func (u *HistoryUser) setPassword(password string) error {
	u.Password = []byte(password)
	return nil
}

func (u *HistoryUser) setBlogs(blogs []ActiveBlog) error {
	HBs := []HistoryBlog{}
	for _, blog := range blogs {
		HB := HistoryBlog{
			HistoryBlogID: temporary.INITIAL_ID,
			HistoryUserID: temporary.INITIAL_ID,
			ActiveBlogID:  blog.GetID(),
			ActiveUserID:  blog.GetActiveUserID(),
			Name:          blog.GetName(),
			Title:         blog.GetTitle(),
			Context:       blog.GetContext(),
			CreatedAt:     time.Time{},
			UpdatedAt:     blog.GetUpdatedAt(),
			ActiveTime:    blog.GetCreatedAt(),
			Revision:      blog.GetRevision(),
		}
		HBs = append(HBs, HB)
	}
	u.Blogs = HBs
	return nil
}

func (b *HistoryUser) setCreatedAt(createdAt time.Time) error {
	b.CreatedAt = createdAt
	return nil
}

func (b *HistoryUser) setUpdatedAt(updatedAt time.Time) error {
	b.UpdatedAt = updatedAt
	return nil
}

func (s *HistoryUser) setActiveCreateAt(ActiveTime time.Time) error {
	s.ActiveTime = ActiveTime
	return nil
}

func (s *HistoryUser) setRevision(revision temporary.REVISION) error {
	s.Revision = revision
	return nil
}

func (u *HistoryUser) GetID() temporary.IDENTIFICATION {
	return temporary.IDENTIFICATION(u.HistoryUserID)
}

func (u *HistoryUser) GetActiveID() int {
	return int(u.ActiveUserID)
}

func (u *HistoryUser) GetName() string {
	return u.Name
}

func (u *HistoryUser) GetEmail() string {
	return u.Email
}

func (u *HistoryUser) GetPassword() string {
	return string(u.Password)
}

func (u *HistoryUser) GetBlogs() []HistoryBlog {
	return u.Blogs
}

func (u *HistoryUser) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *HistoryUser) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (s *HistoryUser) GetActiveCreateAt(ActiveTime time.Time) time.Time {
	return s.ActiveTime
}

func (u *HistoryUser) GetRevision() temporary.REVISION {
	return u.Revision
}
