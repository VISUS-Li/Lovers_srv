package handler

import (
	"Lovers_srv/api/handler/HomeHandler"
	"Lovers_srv/api/handler/JWTHandler"
	"Lovers_srv/api/handler/NotelistHandler"
	"Lovers_srv/api/handler/UserHandler"
	"github.com/gin-gonic/gin"
)

//gin的路由逻辑
func ClientEngine() *gin.Engine{
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	//用户相关接口
	userGroup := api.Group("/user")
	userGroup.POST("/login",UserHandler.Login)
	userGroup.POST("/register",UserHandler.Register)

	//首页接口
	homeGroup := api.Group("home")
	homeGroup.GET("/GetMainCard",HomeHandler.GetMainCard) //获取首页按照日期排列的主卡片
	homeGroup.GET("/GetCardInfoByCount",HomeHandler.GetCardInfoByCount) //获取指定数量的随机首页卡片
	homeGroup.GET("/GetCardInfoByIndex",HomeHandler.GetCardInfoByIndx)  //通过下标范围获取首页卡片
	homeGroup.GET("/GetCardInfoByType",HomeHandler.GetCardInfoByType)   //通过卡片类型和数量获取首页卡片
	homeGroup.POST("/PostCardInfo",HomeHandler.PostCardInfo)
	//需要验证的接口
	AuthGroup := api.Group("/Auth")
	AuthGroup.Use(JWTHandler.JWTMidWare())

	//HomeAuthGroup := AuthGroup.Group("/HomeGroup")
	//HomeAuthGroup.POST("/PostCardInfo",HomeHandler.PostCardInfo)

	NoteListGroup := AuthGroup.Group("/notelist")
	NoteListGroup.POST("/NoteListUp", NotelistHandler.NoteListUp)
	NoteListGroup.POST("/NoteListDown", NotelistHandler.NoteListDown)
	NoteListGroup.POST("/NoteListDel", NotelistHandler.NoteListDel)

	return router
}
