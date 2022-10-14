package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/set2002satoshi/8-4/models"
)


type UserRepository struct {

}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (*models.ActiveUser, error) {
	var user *models.ActiveUser
	db.First(&user, id)
	if user.ActiveUserID <= 0 {
		return &models.ActiveUser{}, errors.New("user is not found")
	}
	return user, nil
}


func (repo *UserRepository) Create(db *gorm.DB, obj *models.ActiveUser) (user *models.ActiveUser, err error) {
	if result := db.Create(obj); result.Error != nil {
		return &models.ActiveUser{}, errors.New("create user failed")
	}
	var createdUser models.ActiveUser
	db.Where("active_user_id = ?", obj.GetID()).First(&createdUser)
	return &createdUser, nil
}


func (repo *UserRepository) FindAll(db *gorm.DB) ([]*models.ActiveUser, error) {
	users := []*models.ActiveUser{}
	db.Find(&users)
	if users == nil {
		return nil, errors.New("DBからデータを取得するに失敗")
	}
	return users, nil
}

func (repo *UserRepository) Update(tx *gorm.DB, obj *models.ActiveUser) (*models.ActiveUser, error) {
	if result := tx.Updates(obj); result.Error != nil {
		fmt.Println("found not Save")
		fmt.Println(result.Error)
		return &models.ActiveUser{}, errors.New("update user failed")
	}
	var user *models.ActiveUser
	findResult := tx.Where("active_user_id = ?", obj.ActiveUserID).First(&user)
	if findResult.Error != nil {
		return &models.ActiveUser{}, errors.New("Updated user not found")
	}
	return user, nil

}


func (repo *UserRepository) DeleteByID(tx *gorm.DB, id int) error {
	activeUser := []models.ActiveUser{}
	if result := tx.Unscoped().Delete(activeUser, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *UserRepository) InsertHistory(tx *gorm.DB, data *models.HistoryUser) (*models.HistoryUser, error) {
	createResult := tx.Create(data)
	if createResult.Error != nil {
		return &models.HistoryUser{}, createResult.Error 
	}
	var History *models.HistoryUser
	findResult := tx.Where("history_user_id = ?", data.HistoryUserID).First(&History)
	if findResult.Error != nil {
		return &models.HistoryUser{}, findResult.Error
	}
	return History, nil
}