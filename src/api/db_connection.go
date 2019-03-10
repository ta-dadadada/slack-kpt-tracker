package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type ConnectionConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DATABASE string
	OPTIONS  string
}

// TODO コンテナ化したら環境変数から読み込むようにする
func getDBConfig() (conf ConnectionConfig) {
	conf = ConnectionConfig{
		HOST:     "db",
		PORT:     3306,
		USER:     "api",
		PASSWORD: "api",
		DATABASE: "kpt",
		OPTIONS:  "parseTime=true",
	}
	return
}

func (c ConnectionConfig) asString() (str string) {
	str = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?%v",
		c.USER, c.PASSWORD,
		c.HOST, c.PORT,
		c.DATABASE,
		c.OPTIONS,
	)
	return
}

func ConnectDB() (db *gorm.DB, err error) {
	conf := getDBConfig().asString()
	db, err = gorm.Open("mysql", conf)
	db.LogMode(true)
	return
}