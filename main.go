package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"mall/models"
	"mall/routers"
)

func main() {
	//初始化主路由
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	//加载设置模板
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "./static")
	//基于cookie
	store := cookie.NewStore([]byte("secret111"))
	r.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(r)

	r.Run("127.0.0.1:8088")

}
