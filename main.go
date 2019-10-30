package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"

	"github.com/TianQinS/fastapi/post"
	"github.com/TianQinS/phaser"
	"github.com/TianQinS/websocket/config"
	"github.com/TianQinS/websocket/event"
	"github.com/TianQinS/websocket/module"
	"github.com/TianQinS/websocket/module/web"

	// for hotfix
	"github.com/TianQinS/websocket/hotfix"
)

var (
	port = flag.String("p", "23456", "socket port")
)

func init() {
	if config.Conf.Debug {
		hotfix.Update()
	}
}

func main() {
	// runtime.GOMAXPROCS(2)
	debug.SetGCPercent(300)

	var (
		exitOnce = sync.Once{}
		exitCh   = make(chan byte)
		app      = event.NewEventMgr()
	)

	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs

		exitOnce.Do(func() {
			app.Close()
			post.Close()
			fmt.Printf("catch sig: %v, exit\n", sig)
			close(exitCh)
		})
	}()

	go func() {
		game := phaser.NewPhaserModule("phaser", 1024000)
		app.Run(
			module.NewRPCModule(config.Conf.RpcTopic, 102400),
			web.NewWebModule(config.Conf.Web.Topic, 1024),
			game,
		)
		// for hotfix
		hotfix.RegisterApp(app)
		game.RegisterApp(app)
		app.Serve("tcp://0.0.0.0:" + *port)
		close(exitCh)
	}()

	<-exitCh
	fmt.Println("exit finish.")
}
