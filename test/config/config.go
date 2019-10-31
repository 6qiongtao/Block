package config

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/config"

	log "github.com/micro/go-micro/util/log"
)

var EmailList = []string{"", }

var (
	AppName     string //项目名称
	HttpAddr    string //服务器ip地址
	HttpPort    string //服务器端口
	RunMode     string //运行模式
	RedisDbnum  string //redis db 编号
	MysqlDbname string //mysql db name
	MysqlUrl    string //mysql 地址
	RedisUrl    string //redis 地址

	ConsulCfg ConSulConfig
)


func InitConfig() {
	log.Info("Config Init")

	//cfg, err := config.NewConfig("ini", "./route/conf/app.conf")
	//if err != nil {
	//
	//}
	//从配置文件读取配置信息
	cfg := beego.AppConfig
	AppName     = cfg.String("appname")
	HttpAddr    = cfg.String("httpaddr")
	HttpPort    = cfg.String("httpport")
	RedisDbnum  = cfg.String("redisdbnum")
	MysqlDbname = cfg.String("mysqldbname")
	MysqlUrl    = cfg.String("mysqlurl")
	RedisUrl    = cfg.String("redisurl")
	RunMode     = cfg.String("runmode")
	return
}

func init() {
	InitConfig()


}