package controller

import (
	"net/http"
	"sam0307204/jwt-Authentication/database"
	"sam0307204/jwt-Authentication/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": error.Error})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"user_ID":  user.ID,
		"email":    user.Email,
		"username": user.Username})
}
