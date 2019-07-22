package v2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]
func ToIndex(c *gin.Context) {
	//appG := app.Gin{C: c}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello Go",
	})
}

func GetTitle(c *gin.Context) {
	//appG := app.Gin{C: c}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "测试请求成功！",
	})
}

func ToBlogPage(c *gin.Context) {
	//appG := app.Gin{C: c}
	c.HTML(http.StatusOK, "blogpage.html", gin.H{
		"title": "hello Go",
	})
}

func ToContact(c *gin.Context) {
	//appG := app.Gin{C: c}
	fmt.Println("执行后台函数---------------")
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "hello Go",
	})
}

func ToSingle(c *gin.Context) {
	//appG := app.Gin{C: c}
	c.HTML(http.StatusOK, "single-post.html", gin.H{
		"title": "hello Go",
	})
}
