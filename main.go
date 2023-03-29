package main

import (
	AuthController "krittii/golang-api-jwt/controller/auth"
	UserController "krittii/golang-api-jwt/controller/user"
	"krittii/golang-api-jwt/middleware"
	"krittii/golang-api-jwt/orm"

	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	orm.InitDB() // initialize Database
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/api/register", AuthController.Register) // register API
	r.POST("/api/login", AuthController.Login)       // Login API
	authorized := r.Group("/api/users", middleware.JWTAuthen())
	authorized.GET("/readall", UserController.ReadAll)
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
