package v2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login-regist.html", gin.H{
		"title": "hello Go",
	})
}

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags [get]

func ToIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "hello Go",
	})
}

func ToContact(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.html", gin.H{
		"title": "hello Go",
	})
}

func ToBlogPage(c *gin.Context) {
	c.HTML(http.StatusOK, "blogpage.html", gin.H{
		"title": "hello Go",
	})
}

func ToSingle(c *gin.Context) {
	c.HTML(http.StatusOK, "single-post.html", gin.H{
		"title": "hello Go",
	})
}
