package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Users struct {
	UserID    int    `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(100)"`
	SlackID   string `gorm:"type:varchar(100);unique_index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Keeps     []Keeps `gorm:"foreignkey:UserID"`
	Problems  []Keeps `gorm:"foreignkey:UserID"`
	Trys      []Keeps `gorm:"foreignkey:UserID"`
}

type Keeps struct {
	KeepID    int `gorm:"primary_key"`
	UserID    int
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Problems struct {
	ProblemID int `gorm:"primary_key"`
	UserID    int
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Trys struct {
	TryID     int `gorm:"primary_key"`
	UserID    int
	Body      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Migrate() {
	db, _ := ConnectDB()
	defer db.Close()
	db.AutoMigrate(&Users{}, &Keeps{}, &Problems{}, &Trys{})
	db.Model(&Keeps{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&Problems{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Model(&Trys{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
}
