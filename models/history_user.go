package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/set2002satoshi/8-4/pkg/module/temporary"
)

type HistoryUser struct {
	HistoryUserID temporary.IDENTIFICATION `gorm:"primaryKey"`
	ActiveID      int
	Name          string 
	Email         string 
	Password      []byte
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
	createdAt time.Time, // このは空のtime
	updatedAt time.Time, // time.Time
	activeTime time.Time, // ここに時間を入れる
	revision temporary.REVISION,
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
		return nil, errors.New("setCreatedAtにエラーが発生しています。")
	}

	if u.setUpdatedAt(updatedAt) {
		return nil, errors.New("setUpdatedAtにエラーが発生しています。")
	}

	if u.setActiveCreateAt(activeTime) {
		return nil, errors.New("setActiveCreateAtにエラーが発生しています。")
	}

	if u.setRevision(revision) {
		return nil, errors.New("setRevisionにエラーが発生しています。")
	}

	return u, nil
}

func (u *HistoryUser) setID(id int) bool {
	i, err := temporary.NewIDENTIFICATION(id)
	if err != nil {
		return true
	}
	u.HistoryUserID = i
	return false
}

func (u *HistoryUser) setActiveID(activeID int) bool {
	u.ActiveID = activeID
	return false
}

func (u *HistoryUser) setEmail(email string) bool {
	u.Email = email
	fmt.Println(u.Email)
	return false
}

func (u *HistoryUser) setName(name string) bool {
	u.Name = name
	fmt.Println(u.Name)
	return false
}

func (u *HistoryUser) setPassword(password string) bool {
	u.Password = []byte(password)
	return false
}

func (b *HistoryUser) setCreatedAt(createdAt time.Time) bool {
	b.CreatedAt = createdAt
	return false
}

func (b *HistoryUser) setUpdatedAt(updatedAt time.Time) bool {
	b.UpdatedAt = updatedAt
	return false
}

func (s *HistoryUser) setActiveCreateAt(ActiveTime time.Time) bool {
	s.ActiveTime = ActiveTime
	return false
}

func (s *HistoryUser) setRevision(revision temporary.REVISION) bool {
	s.Revision = revision
	return false
}

func (u *HistoryUser) GetID() temporary.IDENTIFICATION {
	return temporary.IDENTIFICATION(u.HistoryUserID)
}

func (u *HistoryUser) GetActiveID() int {
	return int(u.ActiveID)
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
