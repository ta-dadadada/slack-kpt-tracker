package main

import (
	"log"
)

func GetUser(userID int) *Users {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	user := &Users{}
	err = db.Where(Users{UserID: userID}).First(user).Error
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return user
}

func CreateUser(userName string, slackID string) *Users {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	user := &Users{
		UserName: userName,
		SlackID:  slackID,
	}
	err = db.Create(user).Error
	if err != nil {
		log.Fatal(err)
	}
	return user
}

// useName と slackId をキーにDBに問い合わせを行い Users エンティティを返す。
// レコードが無い場合は新規レコードを作成して返す。
func GetOrCreateUser(userName string, slackID string) *Users {
	log.Println(userName, slackID)
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	user := &Users{}
	err = db.Where(
		&Users{UserName: userName, SlackID: slackID},
	).FirstOrCreate(user).Error
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return user
}

func CreateKeep(userID int, body string) (*Keeps, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	keep := &Keeps{
		UserID: userID,
		Body:   body,
	}
	err = db.Create(keep).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return keep, err
}
