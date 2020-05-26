package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}




type User struct {
	Id       int    `json:"id",gorm:"auto-increment"`
	Name     string `json:"username"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}


// post form
func login(c *gin.Context) {
	// 获取表单数据 tel，password
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		var user User
		db.Where("tel=?", tel).First(&user)
		if user.Tel == "" {
			response := ResponseData{
				Code:    50001,
				Status:  "error",
				Message: "用户不存在",
				Data:    "",
			}
			c.JSON(200, response)
		} else {
			if user.Password == psd {
				response := ResponseData{
					Code:    200,
					Status:  "success",
					Message: "登录成功",
					Data:    "",
				}
				c.JSON(200, response)
			} else {
				response := ResponseData{
					Code:    5001,
					Status:  "error",
					Message: "密码错误",
					Data:    "",
				}
				c.JSON(200, response)
			}
		}
	}
}

// post form
func register(c *gin.Context) {
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	var user User
	db.Where("tel=?", tel).First(&user)
	if user.Tel == tel {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "手机号已注册",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		newUser := User{Tel: tel, Password: psd}
		db.Create(&newUser)
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "注册成功",
			Data:    "",
		}
		c.JSON(200, response)
	}
}

// get url
func getUserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "参数错误",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	var user User
	db.First(&user, id)
	if user.Id > 0 {
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "",
			Data:    user,
		}
		c.JSON(200, response)
	} else {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "用户不存在",
			Data:    "",
		}
		c.JSON(200, response)
	}
}
