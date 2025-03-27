package models

import (
	"fmt"
	"time"

	"github.com/EduardoMark/expense-tracker/internal/db"
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Title    string    `gorm:"varchar(100); not null" json:"title"`
	Amount   int64     `gorm:"not null" json:"amount"`
	Category string    `gorm:"varchar(100)" json:"category"`
	Date     time.Time `gorm:"not null" json:"date"`
	UserID   uint      `json:"user_id"`
}

func (m *Expense) Save() error {
	db := db.Conn()

	if m.ID == 0 {
		if err := db.Create(m).Error; err != nil {
			return fmt.Errorf("error on create new expense: %w", err)
		}
		return nil
	}

	err := db.Save(m).Error
	if err != nil {
		return fmt.Errorf("error on save expense: %w", err)
	}

	return nil
}

func FindExpenseByIDAndUserID(id, userId uint) (*Expense, error) {
	db := db.Conn()
	var expense Expense

	if err := db.Where("id = ? AND user_id = ?", id, userId).
		First(&expense).Error; err != nil {
		return nil, fmt.Errorf("expense not found: %w", err)
	}

	return &expense, nil
}

func FindAllUserExpenses(userId uint) ([]*Expense, error) {
	db := db.Conn()
	var expenses []*Expense

	if err := db.Where("user_id = ?", userId).
		Find(&expenses).Error; err != nil {
		return nil, fmt.Errorf("expenses not found: %w", err)
	}

	return expenses, nil
}

func FindExpensesByUserIDAndMonth(userId uint, year, month int) ([]*Expense, error) {
	db := db.Conn()
	var expenses []*Expense

	err := db.Where("user_id = ? AND YEAR(date) = ? AND MONTH(date) = ?", userId, year, month).
		Find(&expenses).Error
	if err != nil {
		return nil, fmt.Errorf("expenses per month %d not found: %w", month, err)
	}

	return expenses, nil
}

func FindExpensesByUserIDAndCategory(userId uint, category string) ([]*Expense, error) {
	db := db.Conn()
	var expenses []*Expense

	if err := db.Where("user_id = ? AND category = ?", userId, category).
		Find(&expenses).Error; err != nil {
		return nil, fmt.Errorf("expenses by category not found: %w", err)
	}

	return expenses, nil
}

func DeleteExpense(expenseID, userID uint) error {
	db := db.Conn()

	var expense Expense
	err := db.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense).Error
	if err != nil {
		return fmt.Errorf("expense not found or you do not have permission to delete this expense: %w", err)
	}

	err = db.Delete(&expense).Error
	if err != nil {
		return fmt.Errorf("error deleting expense: %w", err)
	}

	return nil
}
