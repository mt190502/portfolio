package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
	"github.com/mt190502/portfolio/src/api/controller"
	"github.com/mt190502/portfolio/src/utils/config"
	"github.com/mt190502/portfolio/src/utils/logger"
)

func InitAPI() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(ginsession.New())
	router.LoadHTMLGlob("src/files/templates/*")
	router.Static("/static", "src/files/static")

	router.GET("/", controller.Handle("root"))
	router.GET("/2", func(ctx *gin.Context) {
		ctx.HTML(200, "index2.html", map[string]any{
			"profile": map[string]string{
				"Sitetitle": "Portfolio",
			},
		})
	})
	router.GET("/admin", controller.Handle("admin"))
	router.GET("/login", controller.Handle("login"))
	router.GET("/register", controller.Handle("register"))
	router.GET("/favicon.ico", func(ctx *gin.Context) {
		ctx.File("src/files/favicon.ico")
	})

	router.POST("/admin/user/email", controller.HandleAdmin)
	router.POST("/admin/user/password", controller.HandleAdmin)
	router.POST("/admin/settings", controller.HandleAdmin)
	router.POST("/auth/login", controller.HandleLoginPost)
	router.POST("/auth/register", controller.HandleRegisterPost)
	router.POST("/logout", controller.HandleLogout)

	logger.Log.Infof("API service started on address: " + config.Config.WebServer.Address + ":" + strconv.Itoa(config.Config.WebServer.Port))
	router.Run(config.Config.WebServer.Address + ":" + strconv.Itoa(config.Config.WebServer.Port))
}
