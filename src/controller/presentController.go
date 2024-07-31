package controller

import (
	"casamento_api/src/config"
	"casamento_api/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePresent(c *gin.Context) {
	var present model.Present
	conn := config.Connection()

	if err := c.ShouldBindJSON(&present); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro interno do Servidor"})
		return
	}

	result := conn.Create(&present)

	c.JSON(http.StatusOK, result.Error)
}

func SelectPresent(c *gin.Context) {
	var present model.Present
	var guest model.Guest
	var request model.SelectPresentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro interno do Servidor"})
	}

	conn := config.Connection()

	conn.First(&guest, "phone_number = ?", request.PhoneNumber)
	conn.First(&present, "name = ?", request.PresentName)

	present.GuestID = guest.ID
	present.Selected = true

	conn.Save(&present)
}

func GetSelectedPresents(c *gin.Context) {
	var presents []model.Present
	conn := config.Connection()

	conn.Joins("JOIN guests ON presents.guest_id = guests.id").Find(&presents)
	c.JSON(http.StatusOK, presents)
}

func GetUnlectedPresents(c *gin.Context) {
	var presents []model.Present
	conn := config.Connection()

	conn.Find(&presents, "selected = false")
	c.JSON(http.StatusOK, presents)
}

func GetAllPresents(c *gin.Context) {
	var presents []model.Present
	conn := config.Connection()

	conn.Find(&presents)
	c.JSON(http.StatusOK, presents)
}
