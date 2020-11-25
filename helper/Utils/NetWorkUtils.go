package Utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTokenFromHeader(c *gin.Context) (string, error){
	token := c.GetHeader("Token")
	if len(token) <= 0{
		return token, errors.New("Token is nil")
	}
	return token,nil
}

func HttpGetTokenFromHeader(r *http.Request) (string, error){
	token := r.Header.Get("Token")
	if len(token) <= 0{
		return token, errors.New("Token is nil")
	}
	return token,nil
}
