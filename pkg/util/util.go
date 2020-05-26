package util

import "go_blog/go-gin-example/pkg/setting"

//import "github.com/go-gin-example/pkg/setting"

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}