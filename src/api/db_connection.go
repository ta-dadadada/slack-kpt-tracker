package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type ConnectionConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DATABASE string
	OPTIONS  string
}

func getDBConfig() (conf ConnectionConfig) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	conf = ConnectionConfig{
		HOST:     host,
		PORT:     port,
		USER:     user,
		PASSWORD: passwd,
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
