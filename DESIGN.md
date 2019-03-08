# slack-kpt-tracker設計書

go言語でKPTできるslackbotをつくる

## 最小要件
 
機能要件

* ユーザー登録ができる
* ユーザーはkeepのリストを取得できる
* ユーザーはproblemのリストを取得できる
* ユーザーはtryのリストを取得できる
* ユーザーはkeepを登録できる
* ユーザーはproblemを登録できる
* ユーザーはtryを登録できる

非機能要件

* コンテナ化されていること

構成要素

* slack用 messaging API
* データ管理API
* データベース

フレームワーク

* web: 
  * [goa](https://github.com/goadesign/goa)
    * swagger対応してるらしい
  * [gin](https://github.com/gin-gonic/gin)
    * 圧倒的なスター数
    * [go-swagger](https://github.com/go-swagger/go-swagger) 使えばドキュメント作れそう
* ORM
  * [gorm](https://github.com/jinzhu/gorm) 
* slack
  * [nlopes/slack](https://github.com/nlopes/slack)
  * ここが主題ではないので python で作ってもいい

ミドルウェア

* DB
  * MySQL

実装順序

1. DBの準備
1. ユーザー取得機能
1. ユーザー登録機能
1. 上記がslackから使える
1. DBの準備
1. keep取得機能
1. problem取得機能
1. try取得機能
1. 上記がslackから使える
1. keep登録機能
1. problem登録機能
1. try登録機能
1. 上記がslackから使える

## 展望

### 非機能

こちらを優先。

* Swagger API Document を生成する
* slack-messaging-API、データ管理APIそれぞれにウェブサーバを置いて完全なマイクロサービス
* テストコードを書く
* CI
* ユーザ情報をキャッシュするために Redis を準備する

### 機能

* 日時指定してkptを取得するなどの検索機能
* Problemを解決済にする/Tryを達成済にする
* Problemに関連するTryを取得するなどレポート機能
* GUI
