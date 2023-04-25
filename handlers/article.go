package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/models"
)

func CreateArticle(h *Handler) gin.HandlerFunc {

	return func(c *gin.Context) {
		var article models.Article
		if err := c.BindJSON(&article); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "error occured in databinding",
				"error":   err.Error()})
			return
		}

		articleBytes, err := json.Marshal(article)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}

		err = os.WriteFile("articles.txt", articleBytes, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, fmt.Sprintf(`Article Created %v   `, article.Title))

	}

}
