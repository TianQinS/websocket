package web

import (
	"net/http"
	"net/http/pprof"

	"github.com/TianQinS/websocket/config"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/recover"
)

func notFoundHandler(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"ok":   false,
		"data": "404",
	})
}

func forbiddenHandler(ctx iris.Context) {
	ctx.JSON(map[string]interface{}{
		"ok":   false,
		"data": "403",
	})
}

// demo homepage.
func demo(ctx iris.Context) {
	ctx.View("demo/index.html")
}

// InitApi initialize routing functions.
func InitApi(app *iris.Application) {
	if config.Conf.Debug {
		ppApi := app.Party("/debug")
		ppApi.Get("/pprof", pprofHandler(pprof.Index))
		ppApi.Get("/cmdline", pprofHandler(pprof.Cmdline))
		ppApi.Get("/profile", pprofHandler(pprof.Profile))
		ppApi.Post("/symbol", pprofHandler(pprof.Symbol))
		ppApi.Get("/symbol", pprofHandler(pprof.Symbol))
		ppApi.Get("/trace", pprofHandler(pprof.Trace))
		ppApi.Get("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		ppApi.Get("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		ppApi.Get("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		ppApi.Get("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		ppApi.Get("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		ppApi.Get("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}

	app.Get("/demo", demo)
}

func pprofHandler(f http.HandlerFunc) context.Handler {
	handler := http.HandlerFunc(f)
	return func(ctx iris.Context) {
		handler.ServeHTTP(ctx.ResponseWriter(), ctx.Request())
	}
}

// NewApp generate a new web service.
func NewApp(charset string) *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnErrorCode(iris.StatusForbidden, forbiddenHandler)

	app.Configure(iris.WithConfiguration(iris.Configuration{
		DisableAutoFireStatusCode: false,
		Charset:                   charset,
	}))
	InitApi(app)
	return app
}
