package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/set2002satoshi/8-4/models"
)


type UserRepository struct {}

func (repo *UserRepository) FindByID(db *gorm.DB, id int) (user models.ActiveUser, err error) {
	// DBの処理をラップしてる関数
	user = models.ActiveUser{}
	db.First(&user, id)
	if user.ID <= 0 {
		return models.ActiveUser{}, errors.New("user is not found")
	}
	return user, nil
}


func (repo *UserRepository) Create(db *gorm.DB, obj *models.ActiveUser) (user models.ActiveUser, err error) {
	if result := db.Create(obj); result.Error != nil {
		return models.ActiveUser{}, errors.New("create user failed")
	}
	var createdUser models.ActiveUser
	db.Where("ID = ?", obj.GetID()).First(&createdUser)
	return createdUser, nil
}


func (repo *UserRepository) FindAll(db *gorm.DB) ([]models.ActiveUser, error) {
	users := []models.ActiveUser{}
	db.Find(&users)
	if users == nil {
		return nil, errors.New("DBからデータを取得するに失敗")
	}
	return users, nil
}

func (repo *UserRepository) DeleteByID(tx *gorm.DB, id int) error {
	activeUser := []models.ActiveUser{}
	if result := tx.Delete(activeUser, id); result.Error != nil {
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
	findResult := tx.Where("ID = ?", data.ID).First(&History)
	if findResult.Error != nil {
		return &models.HistoryUser{}, findResult.Error
	}
	fmt.Println(History)
	return History, nil
}