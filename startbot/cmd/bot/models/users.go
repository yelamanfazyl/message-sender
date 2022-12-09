package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	TelegramId int64  `json:"telegram_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ChatId     int64  `json:"chat_id"`
	MessageAmount int `json:"message_amount"`
}

type UserModel struct {
	Db *gorm.DB
}

func (m *UserModel) Create(user User) error {

	result := m.Db.Create(&user)

	return result.Error
}

func (m *UserModel) Delete(user User) error {
	result := m.Db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}
	
	return nil
}

func (m *UserModel) FindOne(telegramId int64) (*User, error) {
	existUser := User{}

	result := m.Db.First(&existUser, User{TelegramId: telegramId})

	if result.Error != nil {
		return nil, result.Error
	}

	return &existUser, nil
}

func (m *UserModel) FindAll() ([]User, error) {
	var users []User

    rows := m.Db.Find(&users)

    if rows.Error != nil {
        return nil, rows.Error
    }

    return users, nil
}

func (m *UserModel) Update(user User) error {
	result := m.Db.Save(&user)

	return result.Error
}