package handler

import (
	"strconv"
	"strings"

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
		"4cf0db89d3cf7ee2c74b2246e2fc1a86",
		"zvQy56BQ1RAidm8tPbwNHHduaWrW7w7TWKdn891VMnkEmwtLorPfUUVLQlYskeOnCZJhBTP8gfvWarKstni+rkMoXcIAP3gxrbVQBvLWy0nBB/rw1hIRwnWwd6vzCrK2L7UHOuCdGT+eGhw6HD2m9AdB04t89/1O/w1cDnyilFU=",
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

func PushMessage(cloth Cloth, lineId string) {
	fmt.Println(cloth)
	cloth.Events = cloth.Events[1 : len(cloth.Events)-1]
	cloth.Weather = cloth.Weather[1 : len(cloth.Weather)-1]
	cloth.Temperature = cloth.Temperature[1 : len(cloth.Temperature)-1]
	events := strings.Split(cloth.Events, ",")
	weathers := strings.Split(cloth.Weather, ",")
	temperatures := strings.Split(cloth.Temperature, ",")
	event := ""
	temperature := ""
	weather := ""
	for i := 0; i < len(events); i = i + 1 {
		event = event + events[i]
	}
	for i := 0; i < len(temperatures); i = i + 1 {
		temperature = temperature + temperatures[i]
	}
	for i := 0; i < len(weathers); i = i + 1 {
		weather = weather + weathers[i]
	}

	fmt.Println(event, temperature, weather)
	bot, err := linebot.New(
		// 	// os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		// 	// os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		"0eeb5550fe74acad428d9dcde42f2684",
		"2OeWynHvMPKxzpumMLjBEqjP/LIx0wUFH9K4CodnbFpJcp+bZdY/xHbIPuRSavlvMt6etI9WzhVynW38pwjCNl4qMp8EVdAnWIMJ/T1nIunYKuJUONcoyrUXMnNHXoP5Cjr7TTSf+VWea/d9XbektAdB04t89/1O/w1cDnyilFU=",
	)
	// err = godotenv.Load("env/dev.env")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// message := new(Message)
	// if err := c.Bind(message); err != nil {
	// 	return err
	// }

	// fmt.Println(os.Getenv("LIFF_URL"))
	liffurl := "https://liff.line.me/1660758429-GMo1JYaa"

	link := liffurl + "/clothes-detail/" + strconv.Itoa(cloth.Id)
	json := `{
		"type": "bubble",
		"header": {
			"type": "box",
			"layout": "vertical",
		"contents": [
			{
				"type": "text",
				"text": "laZinator",
				"size": "xl",
				"margin": "none",
				"style": "italic",
				"align": "start"
			}
		],
		"spacing": "none",
		"margin": "none",
		"height": "60px"
		},
		"hero": {
			"type": "image",
			"url": "` + cloth.Cloth + `",
			"size": "full",
			"aspectRatio": "20:13",
			"aspectMode": "fit",
			"action": {
				"type": "uri",
				"uri": "http://linecorp.com/"
			}
		},
		"body": {
			"type": "box",
			"layout": "vertical",
			"contents": [
			{
				"type": "text",
				"weight": "bold",
				"size": "xl",
				"text": "コーデ情報"
			},
			{
				"type": "box",
				"layout": "vertical",
				"margin": "lg",
				"spacing": "sm",
				"contents": [
				{
					"type": "box",
					"layout": "baseline",
					"contents": [
					{
						"type": "text",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 3,
						"text": "イベント",
						"margin": "sm",
						"wrap": true
					},
					{
						"type": "text",
						"text": "` + event + `",
						"color": "#333333",
						"size": "sm",
						"flex": 7,
						"margin": "xxl",
						"wrap": true
					}
					],
					"spacing": "sm"
				},
				{
					"type": "box",
					"layout": "baseline",
					"spacing": "sm",
					"contents": [
					{
						"type": "text",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 3,
						"text": "天気",
						"margin": "sm",
						"wrap": true
					},
					{
						"type": "text",
						"text": "` + weather + `",
						"color": "#333333",
						"size": "sm",
						"flex": 7,
						"margin": "xxl",
						"wrap": true
					}
				]
				},
				{
					"type": "box",
					"layout": "baseline",
					"contents": [
					{
						"type": "text",
						"color": "#aaaaaa",
						"size": "sm",
						"flex": 3,
						"text": "気温",
						"margin": "sm",
						"wrap": true
					},
					{
						"type": "text",
						"text": "` + temperature + `",
						"color": "#333333",
						"size": "sm",
						"flex": 4,
						"margin": "xxl",
						"wrap": true
					}
					],
					"spacing": "sm"
				}
				]
			}
			]
		},
		"footer": {
			"type": "box",
			"layout": "horizontal",
			"contents": [
			{
				"type": "button",
				"action": {
					"type": "postback",
					"label": "他のは？",
					"data": "hello"
				  },
				"flex": 1,
				"style": "secondary",
				"margin": "none"
			},
			{
				"type": "button",
				"action": {
					"type": "uri",
					"label": "詳細",
					"uri": "` + link + `"
				},
				"flex": 1,
				"style": "secondary",
				"margin": "md"
			}
			],
			"height": "60px",
			"paddingBottom": "30px"
		}
		}`
	fmt.Println(json)
	jsonData := []byte(json)
	container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
	if err != nil {
		fmt.Println("unmershallerr:")
		fmt.Println(err)
	}
	pushMessage := linebot.NewFlexMessage("alt text", container)
	if _, err := bot.PushMessage(
		lineId,
		pushMessage,
	).Do(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(pushMessage)

	fmt.Println(container)
	// if _, err = bot.ReplyMessage(, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	// 	log.Print(err)
	// }
	// events, err := bot.ParseRequest(c.Request())
	// if err != nil {
	// 	if err == linebot.ErrInvalidSignature {
	// 		c.Response().WriteHeader(400)
	// 		return c.String(400, "Hello, World!")
	// 	} else {
	// 		c.Response().WriteHeader(500)
	// 		return c.String(500, "Hello, World!")
	// 	}
	// }
	// for _, event := range events {
	// 	switch event.Type {
	// 	case linebot.EventTypeFollow:
	// 		message := "友達登録ありがとう！\nあなたの名前を教えてね！\n※10文字以内"
	// 		fmt.Printf("%v", event)
	// 		if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
	// 			log.Print(err)
	// 			errmsg := "正常にユーザー登録できませんでした\nブロックし、もう一度友達登録をお願いします"
	// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(errmsg)).Do(); err != nil {
	// 				log.Print(err)
	// 			}
	// 		}

	// 	case linebot.EventTypeMessage:
	// 		switch message := event.Message.(type) {
	// 		case *linebot.TextMessage:
	// 			// データベースの接続
	// 			connStr := "user=root dbname=randomcooking password=password host=postgres sslmode=disable"
	// 			db, err := sql.Open("postgres", connStr)
	// 			if err != nil {
	// 				log.Println(err)
	// 				return nil
	// 			}

	// 			rows, err := db.Query(`SELECT * FROM users where lineuserId=$1`, event.Source.UserID)
	// 			if err != nil {
	// 				fmt.Println(err)
	// 			}
	// 			defer rows.Close()

	// 			/*データベースに登録されていない場合*/
	// 			if !rows.Next() {
	// 				_, err := db.Exec(`INSERT INTO users (lineuserid, username) VALUES($1, $2)`, event.Source.UserID, message.Text)
	// 				if err != nil {
	// 					log.Println(err)
	// 					return nil
	// 				}
	// 			} else {
	// 				var id int
	// 				var lineuserid string
	// 				var username string
	// 				rows.Scan(&id, &lineuserid, &username)
	// 				/*メニューの登録*/
	// 				_, err := db.Exec(`INSERT INTO menus (userid, menuname, recipes) VALUES($1, $2, '{"テスト/", "メニューです/"}')`, id, message.Text)
	// 				if err != nil {
	// 					log.Println(err)
	// 					return nil
	// 				}
	// 			}

	// 			defer db.Close()

	// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
	// 				log.Print(err)
	// 			}
	// 		case *linebot.StickerMessage:
	// 			replyMessage := fmt.Sprintf(
	// 				"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
	// 			if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
	// 				log.Print(err)
	// 			}
	// 		}
	// 	}
	// }

}
