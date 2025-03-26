package models

import (
	"fmt"
	"time"

	"github.com/EduardoMark/expense-tracker/internal/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100); uniqueIndex; not null" json:"email"`
	Password string `gorm:"type:varchar(100); not null" json:"password"`
}

func Save(user *User) error {
	db := db.Conn()

	if user.ID == 0 {
		user.CreatedAt = time.Now()
		err := db.Create(user).Error
		if err != nil {
			return fmt.Errorf("error on create new user: %w", err)
		}
		return nil
	}

	err := db.Save(user).Error
	if err != nil {
		return fmt.Errorf("error on save user: %w", err)
	}

	return nil
}

func FindOneUser(id uint) (*User, error) {
	user := User{}
	db := db.Conn()

	result := db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("error when find user: %w", result.Error)
	}

	return &user, nil
}

func FindAllUsers() (*[]*User, error) {
	var users []*User
	db := db.Conn()

	result := db.Find(&users)

	if result.Error != nil {
		return nil, fmt.Errorf("error on find all users: %w", result.Error)
	}

	return &users, nil
}

func UpdateUser(id uint, user User) error {
	db := db.Conn()

	result := db.Model(User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return fmt.Errorf("error on update user: %w", result.Error)
	}

	return nil
}

func DeleteUser(id uint) error {
	db := db.Conn()
	user := User{}

	result := db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("error on delete user: %w", result.Error)
	}

	return nil
}
