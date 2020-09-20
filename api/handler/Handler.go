package handler

import (
	"Lovers_srv/api/handler/NotelistHandler"
	"Lovers_srv/api/handler/UserHandler"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

//gin的路由逻辑
func ClientEngine() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//配置JWT中间件
	//AuthMiddleWare := &helper.GinJWTMiddleware{
	//	Realm:                 "UserAuth",
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
	userGroup.POST("/login",UserHandler.Login)
	userGroup.POST("/register",UserHandler.Register)
	/*/api/client*/

	NoteListGroup := api.Group("/notelist")
	NoteListGroup.POST("/NoteListUp", NotelistHandler.NoteListUp)
	NoteListGroup.POST("/NoteListDown", NotelistHandler.NoteListDown)
	NoteListGroup.POST("/NoteListDel", NotelistHandler.NoteListDel)
	return router
}
