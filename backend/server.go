package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/laut0104/laZinator/handler"
	_ "github.com/lib/pq"

	"fmt"
)

type LineAuthResponse struct {
	Iss     string   `json:"iss"`
	Sub     string   `json:"sub"`
	Aud     string   `json:"aud"`
	Exp     int64    `json:"exp,string"`
	Iat     int64    `json:"iat,string"`
	Nonce   string   `json:"nonce"`
	Amr     []string `json:"amr"`
	Name    string   `json:"name"`
	Picture string   `json:"picture"`
}
type JwtClaims struct {
	Name   string `json:"name"`
	UserId string `json:"uid"`
	jwt.RegisteredClaims
}

func main() {
	fmt.Println("Hello")
	// // インスタンスを作成
	e := echo.New()

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.GET("/", hello)
	e.POST("/callback", handler.Line)
	e.GET("/auth/line/callback", login)
	e.GET("/user", handler.GetUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func login(c echo.Context) error {
	code := c.QueryParam("access_token")
	values := url.Values{}
	values.Set("client_id", "1660690567")
	values.Add("id_token", code)
	req, err := http.NewRequest(
		"POST",
		"https://api.line.me/oauth2/v2.1/verify",
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	fmt.Println(req)

	byteArray, err := io.ReadAll(resp.Body)
	post := new(LineAuthResponse)
	if err != nil {
		fmt.Println("Error")
	}
	err = json.Unmarshal(byteArray, &post)
	if err != nil {
		fmt.Println(err)
	}
	token := handler.Getjwt(post.Name, post.Sub)

	fmt.Println(token)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
