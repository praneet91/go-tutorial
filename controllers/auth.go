package controllers

import (
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuth(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (ac *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login", ac.Login())
	routes.POST("/register", ac.Register())
}

func (ac *AuthController) Ok() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	}
}

func (ac *AuthController) Register() gin.HandlerFunc {

	type RegisterRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {

		var registerRequest RegisterRequest

		if err := c.ShouldBindJSON(&registerRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, error := ac.authService.Register(&registerRequest.Email, &registerRequest.Password)
		if error != nil {
			c.JSON(500, gin.H{"error": error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": user,
		})
	}
}

func (ac *AuthController) Login() gin.HandlerFunc {

	type LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	return func(c *gin.Context) {

		var loginRequest LoginRequest

		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, error := ac.authService.Login(&loginRequest.Email, &loginRequest.Password)
		if error != nil {
			c.JSON(500, gin.H{"error": error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": user,
		})
	}
}
