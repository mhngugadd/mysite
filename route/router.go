package route

import (
	"github.com/gin-gonic/gin"
	"github.com/mhngugadd/mysite/handlers"
	"github.com/appleboy/gin-jwt"
	"net/http"
)

func Router()  {
	r := gin.New()
	// 开启日志
	r.Use(gin.Logger())
	// 开启自动恢复
	r.Use(gin.Recovery())

	// 使用路由组
	auto := r.Group("/auth")
	// 初始化jwt中间件
	jwtMiddleware := handlers.JTWMiddleware()
	// 添加jwt中间件，使该组路由必须通过jwt认证才能访问
	// 所有需要认证才能访问的资源路由都在此注册
	auto.Use(jwtMiddleware.MiddlewareFunc())
	{
		auto.GET("/hello",helloHandler)
	}

	// 登录
	r.POST("/login",jwtMiddleware.LoginHandler)
	//r.POST("/register")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200,"Test server")
	})

	http.ListenAndServe(":8090",r)
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}