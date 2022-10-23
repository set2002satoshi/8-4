package models

import (
	"errors"
	"time"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"

)

type HistoryBlog struct {
	HistoryBlogID temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveBlogID  uint
	Title         string
	Context       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ActiveTime    time.Time
	Revision      temporary.REVISION
}

func NewHistoryBlog(
	id int,
	activeBlogID int,
	title string,
	context string,
	createdAt time.Time, // このは空のtime
	updatedAt time.Time, // time.Time
	activeTime time.Time, // ここに時間を入れる
	revision temporary.REVISION,
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
		return nil, errors.New("setCreatedAtにエラーが発生しています。")
	}

	if b.setUpdatedAt(updatedAt) {
		return nil, errors.New("setUpdatedAtにエラーが発生しています。")
	}

	if b.setActiveCreateAt(activeTime) {
		return nil, errors.New("setActiveCreateAtにエラーが発生しています。")
	}

	if b.setRevision(revision) {
		return nil, errors.New("setRevisionにエラーが発生しています。")
	}

	return b, nil
}

func (u *HistoryBlog) setID(id int) bool {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return true
	}
	u.HistoryBlogID = i
	return false
}

func (u *HistoryBlog) setActiveBlogID(id int) bool {
	if id < 0 {
		return false
	}
	u.ActiveBlogID = uint(id)
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

func (b *HistoryBlog) setCreatedAt(createdAt time.Time) bool {
	b.CreatedAt = createdAt
	return false
}

func (b *HistoryBlog) setUpdatedAt(updatedAt time.Time) bool {
	b.UpdatedAt = updatedAt
	return false
}

func (s *HistoryBlog) setActiveCreateAt(ActiveTime time.Time) bool {
	s.ActiveTime = ActiveTime
	return false
}

func (s *HistoryBlog) setRevision(revision temporary.REVISION) bool {
	s.Revision = revision
	return false
}

func (u *HistoryBlog) GetID() temporary.IDENTIFICATION {
	return temporary.IDENTIFICATION(u.HistoryBlogID)
}

func (u *HistoryBlog) GetActiveBlogID() temporary.IDENTIFICATION {
	return temporary.IDENTIFICATION(u.ActiveBlogID)
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

func (s *HistoryBlog) GetActiveCreateAt(ActivetTime time.Time) time.Time {
	return s.ActiveTime
}

func (s *HistoryBlog) GetRevision() int {
	return int(s.Revision)

}
