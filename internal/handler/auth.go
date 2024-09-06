package handler

import (
	"fitness/internal/auth"
	"fitness/internal/model"
	"fitness/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserService service.UserService
}

// Login - Foydalanuvchini tizimga kirishi
func (h *AuthHandler) Login(c *gin.Context) {
	var user model.User
	// JSON so'rovini parse qilish
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Foydalanuvchini email bo'yicha olish
	storedUser, err := h.UserService.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Parolni tekshirish
	if !service.CheckPasswordHash(user.Password, storedUser.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// JWT token yaratish
	token, err := auth.CreateToken(storedUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// Tokenni qaytarish
	c.JSON(http.StatusOK, gin.H{"token": token})
}
