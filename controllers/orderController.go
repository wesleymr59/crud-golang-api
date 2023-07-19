package controllers

import (
	"crud-golang-api/database"
	"crud-golang-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthyCheck(c *gin.Context) {
	c.JSON(200, "To vivao")
}

func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Client").Find(&orders)
	c.JSON(200, orders)
}

func CreateNewOrder(c *gin.Context) {
	var order []models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if result := database.DB.Preload("Client").Create(&order); result.Error != nil {
		fmt.Println("caiu no erro")
		c.JSON(http.StatusNotFound, "id nao encontrado")
		return
	}
	c.JSON(http.StatusOK, order)

}

func UpdateNewOrder(c *gin.Context) {
	var order models.Order
	id := c.Params.ByName("id")
	if result := database.DB.First(&order, id); result.Error != nil {
		fmt.Println("caiu no erro")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Nao achou nada no sql"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&order).UpdateColumns(order)
	c.JSON(http.StatusOK, order)
}
