package db

import (
	"fmt"
	"go_jwt/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (dbIns *gorm.DB){
	dsn := config.Conf.Mysql.Address
	var err error
	dbIns, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}

	
	fmt.Println("connect ok")

	return dbIns
}