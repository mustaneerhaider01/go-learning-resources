package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)

	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	router.POST("/signup", signup)
	router.POST("/login", login)
}
