package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/models"
)

func StartActivity(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		activity := models.Activity{}
		if err := c.BindJSON(&activity); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in databinding",
				"error":   err.Error()})
			return
		}

		if err := h.DB.StartActivity(activity); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured while creating activity",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("activity %v has been created!", activity.Activity_name))
	}
}

func EndActivity(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 64)
		id := uint(uintt)

		duration, err := h.DB.EndActivity(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf(`total activity time =  '%v'  `, duration))
	}
}

func DeleteActivity(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)

		if err := h.DB.DeleteActivity(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, "activity has been deleted!")
	}
}

func UpdateActivity(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		new_name := c.PostForm("newname")

		updates := make(map[string]interface{})
		updates["activity_name"] = new_name
		if err := h.DB.UpdateActivity(id, updates); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "activity not found",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("activity name has been updated to %v  !", new_name))
	}
}
