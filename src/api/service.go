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

func GetKeepList(userID int) ([]Keeps, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	keeps := &[]Keeps{}
	err = db.Where(Keeps{UserID: userID}).Find(keeps).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return *keeps, err
}

func CreateProblem(userID int, body string) (*Problems, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	problem := &Problems{
		UserID: userID,
		Body:   body,
	}
	err = db.Create(problem).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return problem, err
}

func GetProblemList(userID int) ([]Problems, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	problems := &[]Problems{}
	err = db.Where(Problems{UserID: userID}).Find(problems).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return *problems, err
}

func CreateTry(userID int, body string) (*Trys, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	try := &Trys{
		UserID: userID,
		Body:   body,
	}
	err = db.Create(try).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return try, err
}

func GetTryList(userID int) ([]Trys, error) {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	tries := &[]Trys{}
	err = db.Where(Trys{UserID: userID}).Find(tries).Error
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return *tries, err
}
