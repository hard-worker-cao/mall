package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleWare(c *gin.Context) {
	fmt.Println(time.Now())
	fmt.Println(c.Request.URL)
	c.Set("username", "admin")

	//定义统计日志
	cCopy := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("完成! Path:", cCopy.Request.URL.Path)
	}()
}
