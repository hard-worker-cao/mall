package routers

import (
	"github.com/gin-gonic/gin"
	"mall/controllers/admin"
	"mall/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitMiddleWare)
	{
		//初始页面逻辑
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		//登陆页面逻辑
		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.POST("/dologin", admin.LoginController{}.Dologin)

		//

	}

}
