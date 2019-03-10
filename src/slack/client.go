// TODO Get.+Listを共通化する
// TODO Create.+Listを共通化する
package main

import (
	bytes2 "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const HOST = "http://api:8080"

func GetOrCreateUser(userName string, slackID string) (*Users, error) {
	params := url.Values{}
	params.Add("UserName", userName)
	params.Add("SlackID", slackID)
	uri := fmt.Sprintf("%v/user?%v", HOST, params.Encode())
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	user := &Users{}
	if err := json.Unmarshal(body, user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetKeepList(userID int) ([]Keeps, error) {
	var entities []Keeps
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "keep")
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	if err := json.Unmarshal(body, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

func CreateKeep(userID int, body string) (*http.Response, error) {
	entity := Keeps{
		Body: body,
	}
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "keep")
	bytes, _ := json.Marshal(&entity)
	res, err := http.Post(uri, "application/json",
		bytes2.NewReader(bytes))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Fail to create")
	}
	return res, nil
}

func GetProblemList(userID int) ([]Problems, error) {
	var entities []Problems
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "problem")
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// jsonを構造体へデコード
	if err := json.Unmarshal(body, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

func CreateProblem(userID int, body string) (*http.Response, error) {
	entity := Problems{
		Body: body,
	}
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "problem")
	bytes, _ := json.Marshal(&entity)
	res, err := http.Post(uri, "application/json",
		bytes2.NewReader(bytes))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Fail to create")
	}
	return res, nil
}

func GetTryList(userID int) ([]Trys, error) {
	var entities []Trys
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "try")
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", res.StatusCode)
	}
	defer res.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// jsonを構造体へデコード
	if err := json.Unmarshal(body, &entities); err != nil {
		return nil, err
	}
	return entities, nil
}

func CreateTry(userID int, body string) (*http.Response, error) {
	entity := Trys{
		Body: body,
	}
	uri := fmt.Sprintf("%v/user/%v/%v", HOST, userID, "try")
	bytes, _ := json.Marshal(&entity)
	res, err := http.Post(uri, "application/json",
		bytes2.NewReader(bytes))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Fail to create")
	}
	return res, nil
}
