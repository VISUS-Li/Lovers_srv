package handler

import (
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

	api := router.Group("/api/client")
	userGroup := api.Group("/user")
	userGroup.POST("/login",UserHandler.Login)
	userGroup.POST("/register",UserHandler.Register)
	/*/api/client*/


	NoteListGroup := api.Group("/notelist")
	NoteListGroup.Use(JWTHandler.JWTMidWare())
	NoteListGroup.POST("/NoteListUp", NotelistHandler.NoteListUp)
	NoteListGroup.POST("/NoteListDown", NotelistHandler.NoteListDown)
	NoteListGroup.POST("/NoteListDel", NotelistHandler.NoteListDel)
	return router
}
