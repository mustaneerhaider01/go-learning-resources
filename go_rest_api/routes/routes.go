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
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelEventRegistration)
	authenticated.GET("/events/:id/registrations", getEventRegistrations)

	router.POST("/signup", signup)
	router.POST("/login", login)
}
