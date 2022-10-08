package models

import (
	"errors"
	"time"
)

type ActiveBlog struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Context   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewActiveBlog(
	id int,
	title string,
	context string,
	createdAt time.Time,
	updatedAt time.Time,
) (*ActiveBlog, error) {
	b := &ActiveBlog{}

	if b.setID(id) {
		return nil, errors.New("idにエラー")
	}

	if b.setTitle(title) {
		return nil, errors.New("titleにエラー")
	}

	if b.setContext(context) {
		return nil, errors.New("contextにエラー")
	}

	if b.setCreatedAt(createdAt) {
		return nil, errors.New("createdにエラー")
	}

	if b.setUpdatedAt(updatedAt) {
		return nil, errors.New("updatedにエラー")
	}

	return b, nil
}

func (u *ActiveBlog) setID(id int) bool {
	if id < 0 {
		return false
	}
	u.ID = uint(id)
	return false
}

func (u *ActiveBlog) setTitle(title string) bool {
	u.Title = title
	return false
}

func (u *ActiveBlog) setContext(Context string) bool {
	u.Context = Context
	return false
}

func (u *ActiveBlog) setCreatedAt(createdAt time.Time) bool {
	u.CreatedAt = createdAt
	return false
}

func (u *ActiveBlog) setUpdatedAt(updatedAt time.Time) bool {
	u.UpdatedAt = updatedAt
	return false
}

func (u *ActiveBlog) GetID() int {
	return int(u.ID)
}

func (u *ActiveBlog) GetTitle() string {
	return u.Title
}

func (u *ActiveBlog) GetContext() string {
	return u.Context
}

func (u *ActiveBlog) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *ActiveBlog) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
