package models

import (
	"errors"
	"time"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type ActiveBlog struct {
	ActiveBlogID      temporary.IDENTIFICATION `gorm:"primaryKey"`
	Title   string
	Context string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revision     temporary.REVISION
}

func NewActiveBlog(
	id int,
	title string,
	context string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt time.Time,
	revision temporary.REVISION,
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
		return nil, errors.New("setCreatedAtにエラーが発生しています。")
	}
	if b.setUpdatedAt(updatedAt) {
		return nil, errors.New("setUpdatedAtにエラーが発生しています。")
	}

	if b.setRevision(revision) {
		return nil, errors.New("setRevisionにエラーが発生しています。")
	}


	return b, nil
}

func (u *ActiveBlog) setID(id int) bool {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return true
	}
	u.ActiveBlogID = i
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
func (s *ActiveBlog) setCreatedAt(createdAt time.Time) bool {
	s.CreatedAt = createdAt
	return false
}

func (s *ActiveBlog) setUpdatedAt(updatedAt time.Time) bool {
	s.CreatedAt = updatedAt
	return false
}

func (s *ActiveBlog) setRevision(revision temporary.REVISION) bool {
	s.Revision = revision
	return false
}

func (s *ActiveBlog) CountUpRevisionNumber(num temporary.REVISION) error {
	
	if s.GetRevision() != num {
		return errors.New("改定番号が異なるため更新はできません")
	}
	if ok := s.setRevision(num + 1); ok {
		return errors.New("Invalid setting")
	}
	return nil
}


func (u *ActiveBlog) GetID() int {
	return int(u.ActiveBlogID)
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
func (s *ActiveBlog) GetRevision() temporary.REVISION {
	return s.Revision
}