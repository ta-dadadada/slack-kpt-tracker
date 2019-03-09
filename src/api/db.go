package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type Users struct {
	UserID    uint   `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(100)"`
	SlackId   string `gorm:"type:varchar(100);unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Keeps     []Keeps `gorm:"foreignkey:UserID"`
	Problems  []Keeps `gorm:"foreignkey:UserID"`
	Trys      []Keeps `gorm:"foreignkey:UserID"`
}

type Keeps struct {
	KeepID    uint `gorm:"primary_key"`
	UserID    uint
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Problems struct {
	ProblemID uint `gorm:"primary_key"`
	UserID    uint
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Trys struct {
	TryID     uint `gorm:"primary_key"`
	UserID    uint
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DBConfig struct {
	HOST     string
	PORT     int
	USER     string
	PASSWORD string
	DATABASE string
	OPTIONS  string
}

func getDBConfig() (conf DBConfig) {
	conf = DBConfig{
		HOST:     "127.0.0.1",
		PORT:     3306,
		USER:     "api",
		PASSWORD: "api",
		DATABASE: "kpt",
		OPTIONS:  "parseTime=true",
	}
	return
}

// TODO コンテナ化したら環境変数から読み込むようにする
func (c DBConfig) asString() (str string) {
	str = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?%v",
		c.USER, c.PASSWORD,
		c.HOST, c.PORT,
		c.DATABASE,
		c.OPTIONS,
	)
	return
}

func connectDB() (db *gorm.DB, err error) {
	conf := getDBConfig().asString()
	db, err = gorm.Open("mysql", conf)
	db.LogMode(true)
	return
}

func Migrate() {
	db, _ := connectDB()
	defer db.Close()
	db.AutoMigrate(&Users{}, &Keeps{}, &Problems{}, &Trys{})
	db.Model(&Keeps{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&Problems{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&Trys{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}

func GetUser(userId int) (user Users) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Where(Users{UserID: uint(userId)}).First(&user).Error
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateUser(userName string, slackId string) {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	user := Users{
		UserName: userName,
		SlackId:  slackId,
	}
	db.Create(&user)
}

// useName と slackId をキーにDBに問い合わせを行い Users エンティティを返す。
// レコードが無い場合は新規レコードを作成して返す。
func GetOrCreateUser(userName string, slackId string) (user Users) {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Where(Users{UserName: userName, SlackId: slackId}).FirstOrCreate(&user).Error
	if err != nil {
		panic(err)
	}
	return user
}
