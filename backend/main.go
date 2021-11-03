package main

import (
	"github.com/Earnny/sa-64/controller"
	"github.com/Earnny/sa-64/entity"
	"github.com/Earnny/sa-64/middlewares"
	"github.com/gin-gonic/gin"
)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Professor Routes
			protected.GET("/api/professors", controller.ListProfessors)
			protected.GET("/api/professor/:id", controller.GetProfessor)

			// Course Routes
			protected.GET("/api/courses", controller.ListCourses)
			protected.GET("/api/course/:id", controller.GetCourse)
			protected.POST("/api/courses", controller.CreateCourse)
			protected.PATCH("/api/courses", controller.UpdateCourse)
			protected.DELETE("/api/courses/:id", controller.DeleteCourse)

			// TA Routes
			protected.GET("/api/tas", controller.ListTAs)
			protected.GET("/api/ta/:id", controller.GetTA)
			protected.POST("/api/tas", controller.CreateTA)
			protected.PATCH("/api/tas", controller.UpdateTA)
			protected.DELETE("/api/tas/:id", controller.DeleteTA)

			// Room Routes
			protected.GET("/api/rooms", controller.ListRooms)
			protected.GET("/api/room/:id", controller.GetRoom)
			protected.POST("/api/rooms", controller.CreateRoom)
			protected.PATCH("/api/rooms", controller.UpdateRoom)
			protected.DELETE("/api/rooms/:id", controller.DeleteRoom)

		}
	}

	// Authentication Routes
	r.POST("/api/login", controller.Login)

	// Run the server
	r.Run()
}


