package main

import (
	"fmt"
	"github.com/hxx258456/pyramidel-chain-baas/dao/mysql"
	"github.com/hxx258456/pyramidel-chain-baas/dao/redis"
	"github.com/hxx258456/pyramidel-chain-baas/internal/localconfig"
	"github.com/hxx258456/pyramidel-chain-baas/pkg/utils/logger"
	"github.com/hxx258456/pyramidel-chain-baas/route"
	"os"
)

func main() {
	//1.加载配置
	// if err := os.Setenv("PYCBAAS_CFG_PATH", "E:\\Ljx\\Test\\pyramidel-chain-baas\\configs"); err != nil {
	// 	logger.Error(err)
	// }
	if err := os.Setenv("PYCBAAS_CFG_PATH", "E:\\github.com\\hxx258456\\pyramidel-chain-baas\\configs"); err != nil {
		logger.Error(err)
	}
	localconfig.Init()
	//2.初始化日志
	logger.CfgConsoleLogger(true, true)
	//3.初始化Mysql
	if err := mysql.Init(&localconfig.Defaultconfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis
	if err := redis.Init(&localconfig.Defaultconfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	//5.注册路由
	route.SetUpRouter(&localconfig.Defaultconfig)
	//6.启动服务

}
