package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

// 公用成功与失败跳转页面
func (con *BaseController) success(c *gin.Context, massage string, redirectURL string) {
	c.String(http.StatusOK, "success")
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"massage":     massage,
		"redirectURL": redirectURL,
	})
}

func (con *BaseController) fail(c *gin.Context, massage string, redirectURL string) {
	c.String(http.StatusOK, "fail")
	c.HTML(http.StatusOK, "admin/public/fail.html", gin.H{
		"massage":     massage,
		"redirectURL": redirectURL,
	})
}
