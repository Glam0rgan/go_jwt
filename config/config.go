package config

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var realPath string
var Conf *Config

type Config struct{
	Mysql MysqlConf  `mapstructure:"mysql-base"`
}

type MysqlConf struct {
	Address string `mapstructure:"address"`
}

func getCurrentDir() string {
	_, fileName, _, _ := runtime.Caller(1)
	aPath := strings.Split(fileName, "/")
	dir := strings.Join(aPath[0:len(aPath)-1], "/")
	return dir
}

func Init() {
	once.Do(func() {
		realPath = getCurrentDir()
		fmt.Println(realPath)
		viper.SetConfigType("toml")
		viper.SetConfigName("mysql")
		viper.AddConfigPath(realPath)
		
		err := viper.ReadInConfig()
		
		if err != nil {
			panic(err)
		}
		
		Conf = new(Config)
		viper.Unmarshal(&Conf)
		fmt.Println(Conf.Mysql.Address)
	})
}