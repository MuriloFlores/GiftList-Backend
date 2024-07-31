package controller

import (
	"casamento_api/src/config"
	"casamento_api/src/controller/auth"
	"casamento_api/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GuestControllerInterface interface {
	CreateGuest(c *gin.Context)
}

func CreateGuest(c *gin.Context) {
	var guest model.Guest
	conn := config.Connection()

	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro interno do Servidor"})
		return
	}

	result := conn.Create(&guest)

	c.JSON(http.StatusOK, result.Error)
}

func findGuestByPhoneNumber(c *gin.Context) {
	var guest model.Guest
	conn := config.Connection()

	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro interno do Servidor"})
	}

	conn.Find(&guest, "phone_number = ?", guest.PhoneNumber)

	c.JSON(http.StatusOK, guest)
}

func findGuestById(c *gin.Context) {
	var guest model.Guest
	conn := config.Connection()

	if err := c.ShouldBindJSON(&guest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro interno do Servidor"})
	}

	conn.Find(&guest, "id = ?", guest.ID)
}

func LoginGuest(c *gin.Context) {
	var login model.LoginRequest
	var guest model.Guest

	conn := config.Connection()

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do Servidor"})
	}

	result := conn.First(&guest, "phone_number = ?", login.PhoneNumber)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario não encontrado"})
		return
	}

	token, err := auth.GenerateToken(guest)
	if err != nil {
		return
	}

	response := model.LoginResponse{
		PhoneNumber: guest.PhoneNumber,
		Name:        guest.Name,
		Token:       token,
	}

	c.JSON(http.StatusOK, response)
}
