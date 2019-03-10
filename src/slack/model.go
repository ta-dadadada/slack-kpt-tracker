package main

import (
	"time"
)

type Users struct {
	UserID int `json:"UserID"`
}

type Keeps struct {
	KeepID    int       `json:"KeepID"`
	UserID    int       `json:"UserID"`
	Body      string    `json:"Body"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type Problems struct {
	ProblemID int       `json:"ProblemID"`
	UserID    int       `json:"UserID"`
	Body      string    `json:"Body"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type Trys struct {
	TryID     int       `json:"TryID"`
	UserID    int       `json:"UserID"`
	Body      string    `json:"Body"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}
