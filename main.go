package main

import (
	"fmt"

	"github.com/liehuonainai3000/goboot/frame/config"
	"github.com/liehuonainai3000/goboot/frame/db"
	"github.com/liehuonainai3000/goboot/frame/gen"
	"github.com/liehuonainai3000/goboot/frame/logger"
	"github.com/liehuonainai3000/goboot/frame/rbac"
	"github.com/liehuonainai3000/goboot/frame/web"
	"github.com/liehuonainai3000/goboot/global"
	"github.com/liehuonainai3000/goboot/internal/router"
	"github.com/liehuonainai3000/goboot/internal/service"
)

/*
*
GIN服务启动
Session设置
日志
*/

func main() {

	//初始化配置
	config.InitConfig(global.Conf, "app")

	//初始化日志
	err := logger.InitLogger()
	if err != nil {
		fmt.Printf("log init Err:%v\n", err)
	}

	//初始化数据库
	err = db.InitDB()
	if err != nil {
		fmt.Printf("db init Err:%v\n", err)
	}

	gen.InitGen()
	//初始化casbin
	err = rbac.InitCasbin()
	if err != nil {
		fmt.Printf("casbin init Err:%v\n", err)
	}
	//初始化service
	service.InitService(db.GetDefaultDB())
	//初始化rbac数据
	service.InitRbacData()

	//启动server
	web.StartGinServer(router.InitRouter, &web.GinConfig{
		Port:  global.Conf.Server.Port,
		Debug: global.Conf.Debug,
	})
}
