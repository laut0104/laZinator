package handler

import (
	_ "github.com/lib/pq"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Line(c echo.Context) error {
	bot, err := linebot.New(
		// 	// os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		// 	// os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		"0eeb5550fe74acad428d9dcde42f2684",
		"2OeWynHvMPKxzpumMLjBEqjP/LIx0wUFH9K4CodnbFpJcp+bZdY/xHbIPuRSavlvMt6etI9WzhVynW38pwjCNl4qMp8EVdAnWIMJ/T1nIunYKuJUONcoyrUXMnNHXoP5Cjr7TTSf+VWea/d9XbektAdB04t89/1O/w1cDnyilFU=",
	)
	events, err := bot.ParseRequest(c.Request())
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Response().WriteHeader(400)
			return c.String(400, "Hello, World!")
		} else {
			c.Response().WriteHeader(500)
			return c.String(500, "Hello, World!")
		}
	}
	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeFollow:
			message := "友達登録ありがとう！\nあなたの名前を教えてね！\n※10文字以内"
			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
				fmt.Println(err)
				errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errmsg)).Do(); err != nil {
					fmt.Println(err)
				}
			}

		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// データベースの接続
				db := OpenDB()
				defer db.Close()
				fmt.Println(db, message)

				rows, err := db.Query(`SELECT * FROM users where lineuserId=$1`, event.Source.UserID)
				if err != nil {
					fmt.Println("db_err")
					fmt.Println(err)
				}
				defer rows.Close()

				/*データベースに登録されていない場合*/
				if !rows.Next() {
					_, err := db.Exec(`INSERT INTO users (lineuserid, username) VALUES($1, $2)`, event.Source.UserID, message.Text)
					if err != nil {
						fmt.Println("db_error")
						fmt.Println(err)
						errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errmsg)).Do(); err != nil {
							fmt.Println(err)
						}
					}
					message := "友達登録が完了しました！"
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
						fmt.Println(err)
						errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errmsg)).Do(); err != nil {
							fmt.Println(err)
						}
					}
				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						fmt.Println(err)
					}
				}
				defer db.Close()
			}
		}

	}
	return c.String(http.StatusOK, "Hello, World!")
}
