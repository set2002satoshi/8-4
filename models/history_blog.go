package models

import (
	"time"

	cErr "github.com/set2002satoshi/8-4/pkg/module/customs/errors"
	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type HistoryBlog struct {
	HistoryBlogID temporary.IDENTIFICATION `gorm:"primaryKey"`
	HistoryUserID int
	ActiveBlogID  int
	ActiveUserID  int
	Name          string
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
	name string,
	title string,
	context string,
	createdAt time.Time, // このは空のtime
	updatedAt time.Time, // time.Time
	activeTime time.Time, // ここに時間を入れる
	revision temporary.REVISION,
) (*HistoryBlog, error) {
	b := &HistoryBlog{}

	var err error
	err = cErr.Combine(err, b.setID(id))
	err = cErr.Combine(err, b.setActiveBlogID(activeBlogID))
	err = cErr.Combine(err, b.setName(name))
	err = cErr.Combine(err, b.setTitle(title))
	err = cErr.Combine(err, b.setContext(context))
	err = cErr.Combine(err, b.setCreatedAt(createdAt))
	err = cErr.Combine(err, b.setUpdatedAt(updatedAt))
	err = cErr.Combine(err, b.setActiveCreateAt(activeTime))
	err = cErr.Combine(err, b.setRevision(revision))

	return b, nil
}

func (u *HistoryBlog) setID(id int) error {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	u.HistoryBlogID = i
	return nil
}

func (u *HistoryBlog) setActiveBlogID(id int) error {
	if id < 0 {
		return nil
	}
	u.ActiveBlogID = id
	return nil
}

func (u *HistoryBlog) setName(name string) error {
	u.Name = name
	return nil
}

func (u *HistoryBlog) setTitle(title string) error {
	u.Title = title
	return nil
}

func (u *HistoryBlog) setContext(Context string) error {
	u.Context = Context
	return nil
}

func (b *HistoryBlog) setCreatedAt(createdAt time.Time) error {
	b.CreatedAt = createdAt
	return nil
}

func (b *HistoryBlog) setUpdatedAt(updatedAt time.Time) error {
	b.UpdatedAt = updatedAt
	return nil
}

func (s *HistoryBlog) setActiveCreateAt(ActiveTime time.Time) error {
	s.ActiveTime = ActiveTime
	return nil
}

func (s *HistoryBlog) setRevision(revision temporary.REVISION) error {
	s.Revision = revision
	return nil
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
