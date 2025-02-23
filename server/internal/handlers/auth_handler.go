package handlers

import (
	"net/http"
	"yacc/internal/db"
	"yacc/internal/models"
	"yacc/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PublicKey string `json:"public_key"`
}

func RegisterUser(c *gin.Context) {
	// TODO: rethink data flow. i should read up on auth and practice too.
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, _ := utils.HashPassword(req.Password)

	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		PublicKey: req.PublicKey,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user create failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func LoginUser(c *gin.Context) {}
