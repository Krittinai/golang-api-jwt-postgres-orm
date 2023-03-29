package auth

import (
	"krittii/golang-api-jwt/orm"
	"krittii/golang-api-jwt/orm/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

func Register(c *gin.Context) {
	var json RegisterRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check User Exists
	var checkUser model.User
	orm.Db.Where("username = ?", json.Username).First(&checkUser)

	if checkUser.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": "User Exists"})
		return
	}

	encrypted, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
	user := model.User{
		Username:  json.Username,
		Password:  string(encrypted),
		Firstname: json.Firstname,
		Lastname:  json.Lastname,
	}
	orm.Db.Create(&user)
	if user.ID > 0 {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "created successfully", "userId": user.ID})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "created failed"})
	}

}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check User Exists
	var checkUser model.User
	orm.Db.Where("username = ?", json.Username).First(&checkUser)

	if checkUser.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User Does Not Exists"})
		return
	}

	compareUser := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(json.Password))

	if compareUser == nil {
		mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId":  checkUser.ID,
			"expires": jwt.NewNumericDate(time.Now().Add(5 * time.Second)),
		})
		tokenString, _ := token.SignedString(mySigningKey)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "token": tokenString})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Login Failed"})
	}
}
