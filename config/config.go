package config

import (
	"fmt"
	"runtime/debug"

	"github.com/lytics/confl"
	"github.com/rs/xid"
)

const (
	// json格式配置文件
	CONF_FILE = "./conf.ini"
)

var (
	Conf = new(Config)
)

type Config struct {
	Debug    bool   // 是否调试模式
	Pid      string // 进程全局唯一Id
	RpcTopic string // rpc模块topic名
	NumLoop  int    // socket监听处理协程数
	Kdb      Kdb    // Kvdb
	Mdb      Mdb    // mongo
	Rdb      Rdb    // redis
	Hotfix   Hotfix // 在线更新
	Web      Web
}

type Kdb struct {
	Path             string
	ShrinkPercentage int
	ShrinkMinSize    int
}

type Mdb struct {
	Url  string
	Db   string
	User string
}

type Rdb struct {
	Type       string
	FailFast   bool
	Addrs      []string
	Password   string
	DB         int
	MasterName string
}

type Hotfix struct {
	StdOutput    string
	ModulePrefix []string
	Modules      []string
}

type Web struct {
	Topic      string
	Charset    string
	StaticPath string
	JsPath     string
	HtmlPath   string
	Port       string
}

func init() {
	if _, err := confl.DecodeFile(CONF_FILE, Conf); err != nil {
		fmt.Printf("[Config] read conf file error. info=%s trace=%s\n", err.Error(), string(debug.Stack()))
	} else {
		if Conf.Pid == "" {
			Conf.Pid = xid.New().String()
		}
		fmt.Printf("[Config] read conf ok, pid=%s.\n", Conf.Pid)
	}
}
