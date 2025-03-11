package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct{}

func (con *LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", nil)
}

func (con *LoginController) Dologin(c *gin.Context) {
	c.String(http.StatusOK, "login")
}
