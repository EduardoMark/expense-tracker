package handlers

import (
	"net/http"

	"github.com/EduardoMark/expense-tracker/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RequestBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResponseJSON struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func Login(ctx *gin.Context) {
	body := RequestBody{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := models.FindOneUserByEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Ivalid credentials"})
		return
	}

	// TODO: ADD JWT

	ctx.JSON(http.StatusOK, nil)
}

func Create(ctx *gin.Context) {
	body := RequestBody{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hashedPass),
	}

	if err := user.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func FindAllUsers(ctx *gin.Context) {
	users, err := models.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var result []ResponseJSON
	for _, user := range *users {
		result = append(result, ResponseJSON{
			ID:    user.ID,
			Email: user.Email,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"users": result})
}
