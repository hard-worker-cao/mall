// 权限判断模块
package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall/models"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//1.获取URL /admin/captcha?t=0.890692596889653
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	//2.获取session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string)

	//3.判断session用户信息匹配
	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || (len(userinfoStruct) > 0 || userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/dologin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else {
		if pathname != "/admin/login" && pathname != "/admin/dologin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}

	}
}
