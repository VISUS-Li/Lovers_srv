package handler

import (
	"github.com/gin-gonic/gin"
)

//gin的路由逻辑
func ClientEngine() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//配置JWT中间件
	//AuthMiddleWare := &helper.GinJWTMiddleware{
	//	Realm:                 "UserLogin",
	//	SigningAlgorithm:      "",
	//	Key:                   []byte("lovers"),
	//	Timeout:               time.Hour * 2,
	//	MaxRefresh:            time.Hour,
	//	Authenticator:         func( c *gin.Context) (string,bool){
	//		return "", true
	//	},
	//	Authorizator:          func(userId string, c *gin.Context) bool{
	//		return true
	//	},
	//	PayloadFunc:           nil,
	//	Unauthorized:          func(c *gin.Context, code int, message string) {
	//		c.JSON(code, gin.H{
	//			"code":    code,
	//			"message": message,
	//		})
	//	},
	//	IdentityHandler:       nil,
	//	TokenLookup:           "header:Authorization",
	//	TokenHeadName:         "Bearer",
	//	TimeFunc:              time.Now,
	//	HTTPStatusMessageFunc: nil,
	//	PrivKey:                helper.GetRsaPriKey(),
	//	PubKey:                helper.GetRsaPublicKey(),
	//}
	api := router.Group("/api/client")
	//api.Use(AuthMiddleWare.MiddlewareParseUser)
	userGroup := api.Group("/user")
	userGroup.POST("/login",Login)
	userGroup.POST("/register",Register)
	/*/api/client*/

	NoteListGroup := api.Group("/api/notelist")
	NoteListGroup.POST(".", NoteListUp)
	NoteListGroup.POST(".", NoteListDown)
	return router
}
