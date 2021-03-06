package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
)

func main() {
	//初始化DotServer
	app := dotweb.Classic(file.GetCurrentDirectory())

	app.SetDevelopmentMode()

	app.HttpServer.SetEnabledAutoHEAD(true)
	//app.HttpServer.SetEnabledAutoOPTIONS(true)

	app.SetMethodNotAllowedHandle(func(ctx dotweb.Context) {
		ctx.Redirect(301, "/")
	})

	//设置路由
	InitRoute(app.HttpServer)

	//启动 监控服务
	//app.SetPProfConfig(true, 8081)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

func Index(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	flag := ctx.HttpServer().Router().MatchPath(ctx, "/d/:x/y")
	return ctx.WriteString("index - " + ctx.Request().Method + " - " + ctx.RouterNode().Path() + " - " + fmt.Sprint(flag))
}

func Any(ctx dotweb.Context) error {
	ctx.Response().Header().Set("Content-Type", "text/html; charset=utf-8")
	return ctx.WriteString("any - " + ctx.Request().Method + " - " + ctx.RouterNode().Path())
}

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("go raw http func"))
}

func InitRoute(server *dotweb.HttpServer) {
	server.GET("/", Index)
	server.GET("/d/:x/y", Index)
	server.GET("/x/:y", Index)
	server.GET("/x/", Index)

	server.POST("/post", Index)

	server.Any("/any", Any)
	server.RegisterHandlerFunc("GET", "/h/func", HandlerFunc)
}
