package models

import (
	"errors"
	"time"
)

type HistoryBlog struct {
	ID        uint `gorm:"primaryKey"`
	ActiveBlogID uint 
	Title     string
	Context   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  time.Time `gorm:"index"`
}

func NewHistoryBlog(
	id int,
	activeBlogID int,
	title string,
	context string,
	createdAt time.Time,
	updatedAt time.Time,
) (*HistoryBlog, error) {
	b := &HistoryBlog{}

	if b.setID(id) {
		return nil, errors.New("idにエラー")
	}

	if b.setActiveBlogID(activeBlogID) {
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

func (u *HistoryBlog) setID(id int) bool {
	if id < 0 {
		return false
	}
	u.ID = uint(id)
	return false
}

func (u *HistoryBlog) setActiveBlogID(id int) bool {
	if id < 0 {
		return false
	}
	u.ID = uint(id)
	return false
}

func (u *HistoryBlog) setTitle(title string) bool {
	u.Title = title
	return false
}

func (u *HistoryBlog) setContext(Context string) bool {
	u.Context = Context
	return false
}

func (u *HistoryBlog) setCreatedAt(createdAt time.Time) bool {
	u.CreatedAt = createdAt
	return false
}

func (u *HistoryBlog) setUpdatedAt(updatedAt time.Time) bool {
	u.UpdatedAt = updatedAt
	return false
}

func (u *HistoryBlog) GetID() int {
	return int(u.ID)
}

func (u *HistoryBlog) GetTitle() string {
	return u.Title
}

func (u *HistoryBlog) GetContext() string {
	return u.Context
}

func (u *HistoryBlog) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *HistoryBlog) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}
