package handlers

import (
	"ProjectCinema/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	var body map[string]interface{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	resp, err := utils.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", auth).
		SetBody(body).
		Post("http://order-service:8082/orders")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call order services"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetUserOrders(c *gin.Context) {
	auth := c.GetHeader("Authorization")

	resp, err := utils.Client.R().
		SetHeader("Authorization", auth).
		Get("http://order-service:8082/orders/my")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func CancelOrder(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	orderID := c.Param("id")

	resp, err := utils.Client.R().
		SetHeader("Authorization", auth).
		Delete("http://order-service:8082/orders/" + orderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel order"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}

func GetOrderByID(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	orderID := c.Param("id")

	resp, err := utils.Client.R().
		SetHeader("Authorization", auth).
		Get("http://order-service:8082/orders/" + orderID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get order"})
		return
	}

	c.Data(resp.StatusCode(), "application/json", resp.Body())
}
