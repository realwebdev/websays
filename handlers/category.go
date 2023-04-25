package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/models"
)

func CreateCategory(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.Category // hold object
		var category models.Category
		if err := c.BindJSON(&categories); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in databinding",
				"error":   err.Error()})
			return
		}

		category.ID = uint(len(categories) + 1) // assign a new ID
		categories = append(categories, category)

		c.JSON(http.StatusOK, fmt.Sprintf(`Category Created %v   `, category.Name))
	}

}
