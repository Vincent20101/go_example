package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	ServiceName string           `mapstructure:"name"`
	Port        int              `mapstructure:"port"`
	Mysql       map[string]Infos `mapstructure:"mysql"`
}

type Infos struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func main() {
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("json/mashal2/config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)
	spew.Dump(serverConfig)
	fmt.Printf("%v", v.Get("name"))

	serverConfig2 := ServerConfig{}
	if err := v.UnmarshalKey("mysql", &serverConfig2.Mysql); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig2)
	spew.Dump(serverConfig2)
}
