package main

import (
	"github.com/EduardoMark/expense-tracker/internal/db"
	"github.com/EduardoMark/expense-tracker/internal/models"
	"github.com/EduardoMark/expense-tracker/internal/routes"
)

func main() {
	// Init database
	db.Init()
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Expense{})

	// Routes
	routes.SetupServer()
}
