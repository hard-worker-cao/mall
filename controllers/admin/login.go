package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
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
	username := c.PostForm("username")
	password := c.PostForm("password")
	verifyValue := c.PostForm("verifyValue")

	fmt.Println(username, password)
	//1.验证码是否正确
	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		//c.String(http.StatusOK, "验证码验证成功！SUCCESS")
		con.success(c, "验证码验证成功", "/admin")
		//2.查询数据库中匹配者
		userinfo := []models.Manager{}
		password = models.MD5(password)
		models.DB.Where("username = ? and password=?", username, password).Find(&userinfo)

		if len(userinfo) > 0 {
			//3.执行登录信息，保存在session
			session := sessions.Default(c)
			userinfoSlice, _ := json.Marshal(userinfo)
			session.Set("userinfo", userinfoSlice)
			session.Save()
			con.success(c, "用户登录成功", "/admin")
		} else {
			con.fail(c, "用户名或密码错误,请重试！", "/admin/login")
		}
	} else {
		//c.String(http.StatusOK, "验证码验证失败！FAILED")
		con.fail(c, "验证码验证失败", "/admin/login")
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

func (con *LoginController) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.success(c, "退出登录成功", "/admin/login")
}
