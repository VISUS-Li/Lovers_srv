package JWTHandler

import (
	"Lovers_srv/config"
	"Lovers_srv/helper/Utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

/******
通过用户名密码生成token
******/
func GenerateToken(username string, password string)(string, error){
	//token超时时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	//准备claims
	claims := Claims{
		Username:       username,
		Password:       password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expireTime.Unix(),
			Issuer:"liningtao",
		},
	}

	//通过claims新建token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	secret := []byte(config.GlobalConfig.JwtSecret)
	token,err := tokenClaims.SignedString(secret)
	return token,err
}

/******
解析token到Claims
******/
func ParseToken(token string) (*Claims, error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return config.GlobalConfig.JwtSecret, nil
	})
	if tokenClaims != nil{
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid{
			return claims, nil
		}
	}
	return nil,err
}

/******
JWT验证的中间件
******/
func JWTMidWare() gin.HandlerFunc{
	return func(c *gin.Context){
		//从请求中拿到token
		token ,err := Utils.GetTokenFromHeader(c)
		if err != nil {
			token = c.Query("token")
			if (token == "") {
				token = c.PostForm("token")
			}
		}

		var code int
		var data interface{}
		var msg  string
		code = config.CODE_ERR_SUCCESS
		if (token == ""){
			code = config.CODE_ERR_AUTH_TOKEN_EMPTY
			msg = config.MSG_AUTH_TOKEN_EMPTY
		}else{
			_,err := ParseToken(token)
			if err != nil{
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = config.CODE_ERR_AUTH_CHECK_TOKEN_TIMEOUT
					msg = config.MSG_AUTH_TOKEN_EXPIRE
				default:
					code = config.CODE_ERR_AUTH_CHECK_TOKEN_FAIL
					msg = config.MSG_AUTH_TOKEN_ERROR
				}
			}
		}
		if code != config.CODE_ERR_SUCCESS{
			c.JSON(http.StatusUnauthorized,gin.H{
				"code":code,
				"data":data,
				"msg":msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
