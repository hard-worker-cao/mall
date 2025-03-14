package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ManagerController struct {
	BaseController
}

func (con *ManagerController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/manager/index.html", gin.H{})
}

func (con *ManagerController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/manager/add.html", gin.H{})
}

func (con *ManagerController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/manager/edit.html", gin.H{})
}

func (con *ManagerController) Delete(c *gin.Context) {
	//c.HTML(http.StatusOK, "templates/manager/index.html", gin.H{})
	c.String(200, "delete")
}
