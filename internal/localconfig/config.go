// Copyright (c) 2022 s1ren
// hxx258456/pyramidel-chain-baas is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
// 			http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package localconfig

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	coreconfig "github.com/hxx258456/pyramidel-chain-baas/pkg/utils/config"
	"github.com/spf13/viper"
)

var Defaultconfig TopLevel
var config *viper.Viper

type (
	TopLevel struct {
		Port        string      `json:"port" yaml:"port"`
		Logger      Logger      `json:"logger" yaml:"logger"`
		MySqlConfig MySQLConfig `mapstructure:"MySql"`
		RedisConfig RedisConfig `mapstructure:"Redis"`
	}
)

func Init() {
	//加载配置
	loadConfig()
	watchConfig()
}

// 监听配置改变
func watchConfig() {
	go func() {
		config.WatchConfig()
		config.OnConfigChange(func(e fsnotify.Event) {

			//改变重新加载
			loadConfig()
		})
	}()
}

// 加载配置
func loadConfig() {
	config = viper.New()

	if err := coreconfig.InitViper(config, "config_test"); err != nil {
		panic(err)
	}

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration1: %s", err)
		return
	}
	//fmt.Println(config.AllSettings())

	err := config.Unmarshal(&Defaultconfig)
	if err != nil {
		fmt.Printf("Error reading configuration2: %s", err)
		return
	}
	//fmt.Println(Defaultconfig)
}
