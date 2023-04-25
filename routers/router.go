package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realwebdev/clockify/handlers"
	"github.com/realwebdev/clockify/middleware"
)

func SetupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	router.POST("/signup", handlers.CreateUser(handler))
	router.PUT("/signin", handlers.AuthenticateUser(handler))

	//middleware
	// authorized := router.Group("/")
	// authorized.Use(middleware.TokenAuthMiddleware())
	// {
	router.GET("/getusers", middleware.TokenAuthMiddleware(), handlers.GetUsers(handler))
	router.DELETE("/deleteuser", middleware.TokenAuthMiddleware(), handlers.DeleteUser(handler))

	router.GET("/getproject", middleware.TokenAuthMiddleware(), handlers.GetProjects(handler))
	router.POST("/createproject", middleware.TokenAuthMiddleware(), handlers.CreateProject(handler))
	router.POST("/updateproject", middleware.TokenAuthMiddleware(), handlers.UpdateProject(handler))
	router.DELETE("/deleteproject", middleware.TokenAuthMiddleware(), handlers.DeleteProject(handler))

	router.POST("/createactivity", middleware.TokenAuthMiddleware(), handlers.StartActivity(handler))
	router.POST("/endactivity", middleware.TokenAuthMiddleware(), handlers.EndActivity(handler))
	router.POST("/updateactivity", middleware.TokenAuthMiddleware(), handlers.UpdateActivity(handler))
	router.DELETE("/deleteactivity", middleware.TokenAuthMiddleware(), handlers.DeleteActivity(handler))

	router.POST("/createactivity", middleware.TokenAuthMiddleware(), handlers.StartActivity(handler))
	router.POST("/endactivity", middleware.TokenAuthMiddleware(), handlers.EndActivity(handler))
	router.POST("/updateactivity", middleware.TokenAuthMiddleware(), handlers.UpdateActivity(handler))
	router.DELETE("/deleteactivity", middleware.TokenAuthMiddleware(), handlers.DeleteActivity(handler))

	router.POST("/Createcategory", middleware.TokenAuthMiddleware(), handlers.CreateCategory(handler))
	router.POST("/CreateArticle", middleware.TokenAuthMiddleware(), handlers.CreateArticle(handler))

	// }

	return router
}
