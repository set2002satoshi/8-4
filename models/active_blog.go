package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"
	cErr "github.com/set2002satoshi/8-4/pkg/module/customs/errors"
)

type ActiveBlog struct {
	ActiveBlogID temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveUserID uint
	Name         string
	Title        string `gorm:"not null;size:16"`
	Context      string `gorm:"not null;size:256"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revision     temporary.REVISION
}

func NewActiveBlog(
	id int,
	ActiveUserID int,
	name,
	title,
	context string,
	createdAt time.Time,
	updatedAt time.Time,
	revision temporary.REVISION,
) (*ActiveBlog, error) {
	b := &ActiveBlog{}

	var err error
	err = cErr.Combine(err, b.setID(id))
	err = cErr.Combine(err, b.setActiveUserID(ActiveUserID))
	err = cErr.Combine(err, b.setName(name))
	err = cErr.Combine(err, b.setTitle(title))
	err = cErr.Combine(err, b.setContext(context))
	err = cErr.Combine(err, b.setCreatedAt(createdAt))
	err = cErr.Combine(err, b.setUpdatedAt(updatedAt))
	err = cErr.Combine(err, b.setRevision(revision))

	return b, err
}

func (u *ActiveBlog) setID(id int) error {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return err
	}
	u.ActiveBlogID = i
	return nil
}

func (u *ActiveBlog) setActiveUserID(activeUserID int) error {
	u.ActiveUserID = uint(activeUserID)
	return nil
}

func (u *ActiveBlog) setName(name string) error {
	switch true {
	case name == "":
		fmt.Println(1)
		return errors.New("入力値が存在しません") // <- あとからリテラルを修正
	case name == "anonymous":
		return nil
	default:
		u.Name = name
		return nil
	}
}

func (u *ActiveBlog) setTitle(title string) error {
	u.Title = title
	return nil
}

func (u *ActiveBlog) setContext(Context string) error {
	u.Context = Context
	return nil
}
func (s *ActiveBlog) setCreatedAt(createdAt time.Time) error {
	s.CreatedAt = createdAt
	return nil
}

func (s *ActiveBlog) setUpdatedAt(updatedAt time.Time) error {
	s.CreatedAt = updatedAt
	return nil
}

func (s *ActiveBlog) setRevision(revision temporary.REVISION) error {
	s.Revision = revision
	return nil
}

func (s *ActiveBlog) CountUpRevisionNumber(num temporary.REVISION) error {

	if s.GetRevision() != num {
		return errors.New("改定番号が異なるため更新はできません")
	}
	if err := s.setRevision(num + 1); nil != err {
		return errors.New("invalid setting")
	}
	return nil
}

func (u *ActiveBlog) GetID() int {
	return int(u.ActiveBlogID)
}

func (u *ActiveBlog) GetActiveUserID() int {
	return int(u.ActiveUserID)
}

func (u *ActiveBlog) GetName() string {
	return u.Name
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

func (s *ActiveBlog) AddUserName(name string) error {
	if name == "" {
		return errors.New("name名は空です")
	}
	s.Name = name
	return nil
}

func (s *ActiveBlog) AddActiveUserID(id uint) error {
	s.ActiveUserID = id
	return nil
}
