package Utils

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func GetTokenFromHeader(c *gin.Context) (string, error){
	token := c.GetHeader("Token")
	if len(token) <= 0{
		return token, errors.New("Token is nil")
	}
	return token,nil
}
