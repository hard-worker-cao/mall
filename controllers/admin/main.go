package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type MainController struct{}

func (mCon MainController) Index(context *gin.Context) {
	context.HTML(http.StatusOK, "admin/main/index.html", gin.H{
		"status_code": 200,
	})
}

func (mCon MainController) Welcome(context *gin.Context) {
	context.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}
