package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //不能忘记导入
	"github.com/go-xorm/xorm"
	"github.com/micro/go-micro/util/log"
	"vtoken_digiccy_go/route/model"
)

var Mysql *xorm.Engine

/**
 * 实例化数据库引擎方法：mysql的数据引擎
 */
func NewMysqlEngine() *xorm.Engine {
	log.Info("Mysql Engine Init")
	//数据库引擎 数据库必须存在，
	//mysqlurl := "root:123456789@tcp(127.0.0.1:3306)/vtoken_go?charset=utf8"
	mysqlurl := "root:123456789@tcp(192.168.0.129:3306)/vtoken_go?charset=utf8"
	fmt.Println("====>", mysqlurl)
	engine, err := xorm.NewEngine("mysql", mysqlurl)
	err = engine.CreateTables(new(model.Admin))
	if err != nil {
		log.Info("mysql 连接错误")
	}

	//同步数据库结构：主要负责对数据结构实体同步更新到数据库表
	//Sync2是Sync的基础上优化的方法 表没有，会自己创建
	err = engine.Sync2(new(model.Admin),)
	if err != nil {
		log.Info("mysql NewMysql Engine Init err:", err.Error())
	}

	//设置显示sql语句
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(10)

	return engine
}

func main() {
	Mysql = NewMysqlEngine()
}