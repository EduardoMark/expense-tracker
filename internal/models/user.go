package models

import (
	"fmt"
	"time"

	"github.com/EduardoMark/expense-tracker/internal/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string     `gorm:"type:varchar(100); uniqueIndex; not null" json:"email"`
	Password string     `gorm:"type:varchar(100); not null" json:"password"`
	Expenses []*Expense `gorm:"foreignKey:UserID" json:"expenses"`
}

func (m *User) Save() error {
	db := db.Conn()

	if m.ID == 0 {
		m.CreatedAt = time.Now()
		err := db.Create(m).Error
		if err != nil {
			return fmt.Errorf("error on create new user: %w", err)
		}
		return nil
	}

	err := db.Save(m).Error
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

func FindOneUserByEmail(email string) (*User, error) {
	user := User{}
	db := db.Conn()

	result := db.Where("email = ?", email).First(&user)

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

func DeleteUser(id uint) error {
	db := db.Conn()
	user := User{}

	result := db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return fmt.Errorf("error on delete user: %w", result.Error)
	}

	return nil
}
