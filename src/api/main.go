// TODO DB関連処理をAPI化したらSlack関連処理は切り出す
package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"regexp"
	"strings"
)

var keepRex = regexp.MustCompile(`<.+>\skeep[\s　](.+)`)
var problemRex = regexp.MustCompile(`<.+>\sproblem[\s　](.+)`)

// FIXME リファクタ。ネストやばすぎ
func generateKeepMessage(user *Users, inText string) (outText string) {
	matches := keepRex.FindAllStringSubmatch(inText, -1)
	if len(matches) > 0 && matches[0][1] != "" {
		body := matches[0][1]
		if body == "list" {
			keeps, err := GetKeepList(user.UserID)
			if err != nil {
				log.Fatal(err)
				outText = "keepの取得に失敗しました"
				return
			}
			if len(keeps) == 0 {
				outText = "keepが登録されていません"
				return
			}
			// FIXME 読みやすさ重視。GOの文字列結合コスト的にどうなのか不明
			sb := strings.Builder{}
			sb.WriteString("KEEP\n============\n")
			for i, keep := range keeps {
				sb.WriteString(
					fmt.Sprintf("%v. %v\n", i+1, keep.Body))
			}
			outText = sb.String()
			return
		}
		keep, err := CreateKeep(user.UserID, body)
		if err != nil {
			outText = "Keepの記録に失敗しました"
			return
		}
		outText = fmt.Sprintf("Keepを記録しました: `%v`", keep.Body)
		return
	}
	outText = ""
	return
}

// FIXME リファクタ。ネストやばすぎ
func generateProblemMessage(user *Users, inText string) (outText string) {
	matches := problemRex.FindAllStringSubmatch(inText, -1)
	if len(matches) > 0 && matches[0][1] != "" {
		body := matches[0][1]
		if body == "list" {
			keeps, err := GetProblemList(user.UserID)
			if err != nil {
				log.Fatal(err)
				outText = "Problemの取得に失敗しました"
				return
			}
			if len(keeps) == 0 {
				outText = "Problemが登録されていません"
				return
			}
			// FIXME 読みやすさ重視。GOの文字列結合コスト的にどうなのか不明
			sb := strings.Builder{}
			sb.WriteString("PROBLEM\n============\n")
			for i, problem := range keeps {
				sb.WriteString(
					fmt.Sprintf("%v. %v\n", i+1, problem.Body))
			}
			outText = sb.String()
			return
		}
		keep, err := CreateProblem(user.UserID, body)
		if err != nil {
			outText = "Problemの記録に失敗しました"
			return
		}
		outText = fmt.Sprintf("Problemを記録しました: `%v`", keep.Body)
		return
	}
	outText = ""
	return
}

func getReplyMessage(user *Users, data slack.MessageEvent) (replyMessage string) {
	text := data.Text
	if replyMessage = generateKeepMessage(user, text); replyMessage != "" {
		return
	}
	if replyMessage = generateProblemMessage(user, text); replyMessage != "" {
		return
	}
	replyMessage = "使い方\n" + "====================\n" +
		"`@this_bot [keep|problem] hoge`: hogeを登録します\n" +
		"`@this_bot [keep|problem] list`: リストを取得します\n"
	return
}

func run(cli *slack.Client) int {
	// botIDは不変なんだろうか・・・？
	botID := "UGSV7CUGG"
	rtm := cli.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				data := *(msg.Data.(*slack.MessageEvent))
				// bot宛のメンションか判定
				if strings.Contains(data.Text, fmt.Sprintf("<@%v>", botID)) {
					userInfo, err := cli.GetUserInfo(data.User)
					if err != nil {
						log.Fatal(fmt.Sprintf("データの取得できないSlackユーザ: %v", data.User))
					} else {
						user := GetOrCreateUser(
							userInfo.Profile.DisplayName, userInfo.ID)
						text := getReplyMessage(user, data)
						rtm.SendMessage(rtm.NewOutgoingMessage(text, ev.Channel))
					}
				}
			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}

func main() {
	Migrate()
	token := os.Getenv("SLACK_TOKEN")
	cli := slack.New(token)
	run(cli)
}
