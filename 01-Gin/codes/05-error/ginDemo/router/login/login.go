package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Name string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func UserLogin(c *gin.Context)  {

	//获取表单里面的单个值
	name := c.PostForm("name")
	fmt.Printf("name=%v",name)
	var loginForm LoginForm
	if c.ShouldBind(&loginForm) ==nil{
		if loginForm.Name == "user" && loginForm.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}else {
		c.JSON(401, gin.H{"status": "missing params"})
	}
}
