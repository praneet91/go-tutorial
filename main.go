package main

import (
	"go-tutorial/controllers"
	"go-tutorial/services"

	// "net/http"

	"github.com/gin-gonic/gin"

	internal "go-tutorial/internal/database"
)

func main() {
	r := gin.Default()

	db := internal.InitDB()

	if db != nil {

	}

	// notesService := new(services.NotesService)
	notesService := &services.NotesService{}
	notesService.Init(db)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.POST("/me", func(c *gin.Context) {
	// 	type MeRequest struct {
	// 		Email    string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	var meRequest MeRequest
	// 	if err := c.ShouldBindJSON(&meRequest); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return

	// 	}
	// 	// c.BindJSON(&meRequest)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"email":    meRequest.Email,
	// 		"password": meRequest.Password,
	// 	})
	// })

	NotesController := new(controllers.NotesController)
	NotesController.Init(notesService)
	NotesController.InitRoutes(r)

	authService := services.InitAuthService(db)
	authController := controllers.InitAuth(authService)
	authController.InitRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
