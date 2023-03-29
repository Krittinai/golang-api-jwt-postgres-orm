package user

import (
	"krittii/golang-api-jwt/orm"
	"krittii/golang-api-jwt/orm/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []model.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User Read Success", "user": users})

}
