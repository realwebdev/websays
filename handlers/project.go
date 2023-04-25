package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/models"
)

func CreateProject(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		project := models.Project{}
		if err := c.BindJSON(&project); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in data binding",
				"error":   err.Error()})
			return
		}

		if err := h.DB.CreateProject(project); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured creating project",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("project  created %v", project.Project_name))
	}
}

func GetProjects(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.DB.GetProjects()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Project list not found",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func UpdateProject(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)
		update := c.PostForm("update")

		updates := make(map[string]interface{})
		updates["project_name"] = update
		if err := h.DB.UpdateProject(id, updates); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "The project does not exist in the database",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf(`project name changed to '%v'`, update))
	}
}

func DeleteProject(h *Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		uintt, _ := strconv.ParseUint(c.PostForm("id"), 10, 32)
		id := uint(uintt)

		deleteproject, err := h.DB.DeleteProject(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "The project does not exist in the database",
				"error":   err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf("project %v  deleted !", deleteproject))
	}
}
