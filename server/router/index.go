package router

import (
	"crypto/tls"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/global"
	middleware "github.com/ppoonk/AirGo/router/middleware"
	"github.com/ppoonk/AirGo/web"
	"io"
	"os"
	"strconv"
	"sync"
)

type GinRouter struct {
	Router *gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{Router: nil}
}

var Server = &GinRouter{
	Router: nil,
}

func (g *GinRouter) InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	var writer io.Writer
	if global.Config.SystemParams.Mode == "dev" {
		writer = os.Stdout
	} else {
		writer = io.Discard //关闭控制台输出
	}
	gin.DefaultWriter = writer
	g.Router = gin.Default()
	// targetPtah=web 是embed和web文件夹的相对路径
	g.Router.Use(middleware.Serve("/", middleware.EmbedFolder(web.Static, "web")))
	g.Router.Use(middleware.Cors(), middleware.Recovery())

	//api路由
	apiRouter := g.Router.Group("/api")

	//swagger 路由
	//docs.SwaggerInfo.BasePath = ""
	//swaggerRouter := RouterGroup.Group("/swagger").Use(middleware.ParseJwt(), middleware.Casbin())
	//{
	//	swaggerRouter.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//}
	//注册路由
	g.InitAdminRouter(apiRouter)
	g.InitUserRouter(apiRouter)
	g.InitPublicRouter(apiRouter)
}

func (g *GinRouter) Start() {
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		err := endless.ListenAndServe(":"+strconv.Itoa(global.Config.SystemParams.HTTPPort), g.Router)
		if err != nil {
			global.Logrus.Error("listen: %s", err)
		}
		w.Done()
	}()
	go func() {
		_, err := tls.LoadX509KeyPair("./air.cer", "./air.key") //先验证证书，否则endless fork进程时会空指针panic
		if err == nil {
			err = endless.ListenAndServeTLS(":"+strconv.Itoa(global.Config.SystemParams.HTTPSPort), "./air.cer", "./air.key", g.Router)
			if err != nil {
				global.Logrus.Error("listen: %s", err)
			}
		}
		w.Done()
	}()
	w.Wait()
	//syscall.SIGHUP 将触发重启; syscall.SIGINT, syscall.SIGTERM 并将触发服务器关闭（它将完成运行请求)。https://github.com/fvbock/endless
	// TODO windows下使用endless报错：undefined: syscall.SIGUSR1
	global.Logrus.Info("Server stop")
	os.Exit(0)
}
