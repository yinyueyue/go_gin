package router

import (
	"error/ginDemo/middleware/logger"
	"error/ginDemo/middleware/recover"
	"error/ginDemo/middleware/sign"
	"error/ginDemo/router/ascii"
	"error/ginDemo/router/form"
	"error/ginDemo/router/image"
	"error/ginDemo/router/login"
	"error/ginDemo/router/v1"
	"error/ginDemo/router/v2"
	"error/ginDemo/validator/member"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func InitRouter(r *gin.Engine) {

	r.Use(logger.LoggerToFile(), recover.Recover())

	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/product/add", v1.AddProduct)
		GroupV1.Any("/member/add", v1.AddMember)
	}

	// v2 版本
	GroupV2 := r.Group("/v2").Use(sign.Sign())
	{
		GroupV2.Any("/product/add", v2.AddProduct)
		GroupV2.Any("/member/add", v2.AddMember)
	}

	// 绑定验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValid", member.NameValid)
	}

	demoGroup := r.Group("/demo")
	{
		demoGroup.GET("/ascii-json", ascii.ResponseAsciiJson)
		demoGroup.POST("/login", login.UserLogin)
		demoGroup.POST("/single-file", form.SingleFileUpload)
		demoGroup.POST("/multiple-files", form.MultipleFileUpload)
		demoGroup.GET("/image", image.GetImage)
	}

	//basic auth
	authorizedGroup := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	authorizedGroup.GET("/secrets", func(context *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := context.MustGet(gin.AuthUserKey).(string)

		if secret, ok := secrets[user]; ok {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	//绑定url path参数

	r.GET("/user/:Id", func(context *gin.Context) {
		id := context.Param("Id")

		fmt.Println(id)
		var user User
		if context.ShouldBindUri(&user) == nil {
			(&user).Name = "gin"
			context.JSON(http.StatusOK, user)
		}
	})

	//路由重定向
	r.GET("/index", func(context *gin.Context) {
		context.Request.URL.Path = "https://www.baidu.com"
		r.HandleContext(context)
	})

	//http重定向
	r.GET("/index2", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently,"https://www.baidu.com")
	})

}

type User struct {
	Id   int
	Name string
}
