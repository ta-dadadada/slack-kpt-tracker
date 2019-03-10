package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserParams struct {
	UserName string
	SlackID  string
}

type NoteBody struct {
	Body string
}

func getUserBySlackID(c *gin.Context) {
	params := UserParams{}
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&params) == nil {
		user := GetOrCreateUser(params.UserName, params.SlackID)
		if user == nil {
			// なければ作成するので、失敗していたらサーバエラーを返す
			c.String(http.StatusInternalServerError, "Internal Server Error")
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.String(http.StatusBadRequest, "Bad Request")
	}
}

func getUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	user := GetUser(userID)
	if user == nil {
		c.String(http.StatusNotFound, "Not Found")
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func createUser(c *gin.Context) {
	params := UserParams{}
	if c.ShouldBind(&params) != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	user := CreateUser(params.UserName, params.SlackID)
	if user == nil {
		c.String(http.StatusNotFound, "Not Found")
	} else {
		c.String(http.StatusCreated, "ok")
	}
}

func listKeep(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	entities, err := GetKeepList(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	c.JSON(http.StatusOK, entities)
}

func createKeep(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	body := NoteBody{}
	if c.ShouldBind(&body) != nil {
		c.String(http.StatusBadRequest, "リクエストボディの取得に失敗")
		return
	}
	_, err = CreateKeep(userID, body.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	c.JSON(http.StatusCreated, "ok")
}

func listProblem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	entities, err := GetProblemList(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	c.JSON(http.StatusOK, entities)
}

func createProblem(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	body := NoteBody{}
	if c.ShouldBind(&body) != nil {
		c.String(http.StatusBadRequest, "リクエストボディの取得に失敗")
		return
	}
	_, err = CreateProblem(userID, body.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	c.JSON(http.StatusCreated, "ok")
}

func listTry(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	entities, err := GetTryList(userID)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	c.JSON(http.StatusOK, entities)
}

func createTry(c *gin.Context) {
	userID, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	body := NoteBody{}
	if c.ShouldBind(&body) != nil {
		c.String(http.StatusBadRequest, "リクエストボディの取得に失敗")
		return
	}
	_, err = CreateTry(userID, body.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "パラメータが不正です")
		return
	}
	c.JSON(http.StatusCreated, "ok")
}

func RunAPIServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:id", getUser)
	r.GET("/user", getUserBySlackID)
	r.GET("/user/:id/keep", listKeep)
	r.POST("/user/:id/keep", createKeep)
	r.GET("/user/:id/problem", listProblem)
	r.POST("/user/:id/problem", createProblem)
	r.GET("/user/:id/try", listTry)
	r.POST("/user/:id/try", createTry)
	r.Run() // listen and serve on 0.0.0.0:8080
}
