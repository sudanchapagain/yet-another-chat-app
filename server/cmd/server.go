package cmd

import (
	"log"

	"yacc/internal/db"
	"yacc/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	db.InitDB()
	router := gin.Default()

	public := router.Group("/api")
	{
		public.POST("/register", handlers.RegisterUser)
		public.POST("/login", handlers.LoginUser)
	}

	protected := router.Group("/api")
	// set a middlware
	protected.Use()
	{
		// TODO
		// protected.GET("/users", handlers.GetUsers)
		// protected.GET("/me", handlers.GetProfile)
	}

	router.GET("/ws", func(c *gin.Context) {
		// TODO: idk
	})

	port := ":8080"
	log.Println("running on port", port)
	if err := router.Run(port); err != nil {
		log.Fatal("failed to start:", err)
	}
}
