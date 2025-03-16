package routers

import (
	"github.com/gin-gonic/gin"
	"mall/controllers/admin"
	"mall/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		//初始页面逻辑,登录页面
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)

		//登陆页面逻辑
		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.GET("/logOut", admin.LoginController{}.LogOut)
		adminRouters.POST("/dologin", admin.LoginController{}.Dologin)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		//管理人员界面
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)

		//
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
	}

}
