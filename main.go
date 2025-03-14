package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"mall/routers"
)

func main() {
	//初始化主路由
	rMain := gin.Default()
	//加载设置模板
	rMain.LoadHTMLGlob("templates/**/**/*")
	//设置cookie
	store := cookie.NewStore([]byte("secret_mall"))
	rMain.Use(sessions.Sessions("mysession", store))

	routers.AdminRoutersInit(rMain)

	rMain.Run(":8080")
	
}
