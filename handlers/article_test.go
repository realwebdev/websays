package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/models"
)

func TestCreateArticle(t *testing.T) {
	// Create a test article
	article := models.Article{
		Title:   "Test Article",
		Content: "This is a test article.",
	}

	// Create a new Gin router
	r := gin.Default()

	// Define the POST /articles endpoint
	r.POST("/articles", func(c *gin.Context) {
		var savedArticle models.Article

		if err := c.BindJSON(&savedArticle); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Verify that the saved article matches the original article
		if savedArticle.Title != article.Title || savedArticle.Content != article.Content {
			t.Errorf("Saved article to match original article, got %v", savedArticle)
		}

		c.JSON(http.StatusCreated, savedArticle)
	})

	// Create a new HTTP request with the test article as JSON
	// Create a new HTTP request with the test article as JSON
	reqBody, err := json.Marshal(article)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}
	req, err := http.NewRequest("POST", "/articles", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Perform the request and verify the response
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusCreated {
		t.Errorf("Expected response code %d, got %d", http.StatusCreated, resp.Code)
	}
}
