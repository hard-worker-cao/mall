package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mall/models"
	"net/http"
)

type MainController struct {
	BaseController
}

func (mCon MainController) Index(c *gin.Context) {
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	userinfoStr, ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		c.JSON(200, gin.H{
			"userinfo": userinfoStruct[0].Username,
		})
	} else {
		c.JSON(200, gin.H{
			"userinfo": "session不存在",
		})
	}

	//c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
	//	"status_code": 200,
	//})
}

func (mCon MainController) Welcome(context *gin.Context) {
	context.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
