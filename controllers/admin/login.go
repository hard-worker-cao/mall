package admin

import (
	"github.com/gin-gonic/gin"
	"mall/models"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con *LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con *LoginController) Dologin(c *gin.Context) {
	c.String(http.StatusOK, "login")
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		c.String(http.StatusOK, "验证码验证成功！SUCCESS")
		return
	} else {
		c.String(http.StatusOK, "验证码验证失败！FAILED")
	}

}

func (con *LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	models.ErrHandler(err)

	c.JSON(http.StatusOK, gin.H{
		"CaptchaId":    id,
		"CaptchaImage": b64s,
	})
}
