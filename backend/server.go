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

type OpenWeatherMapAPIResponse struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Coord   Coord     `json:"coord"`
	Wind    Wind      `json:"wind"`
	Dt      int64     `json:"dt"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressuer int     `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type Position struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
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
	e.POST("/cloth", handler.AddCloth)
	e.GET("/clothes", handler.GetClothes)
	e.GET("/cloth/:id", handler.GetCloth)
	e.PUT("/cloth/:id", handler.UpdateCloth)
	e.DELETE("/cloth/:id", handler.DeleteCloth)
	e.GET("/temp", GetTemperature)
	e.POST("/propose", ProposeCloth)

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

func GetTemperature(c echo.Context) error {
	token := "6beffb309ba1da1d671c15746524791d"                   // APIトークン
	endPoint := "https://api.openweathermap.org/data/2.5/weather" // APIのエンドポイント
	// lat := c.QueryParam("lat")
	// lon := c.QueryParam("lon")

	// パラメータを設定
	values := url.Values{}
	values.Set("lat", "35.68")
	values.Set("lon", "139.77")
	values.Set("APPID", token)
	values.Set("units", "metric")

	target := endPoint + "?" + values.Encode()
	fmt.Println(target)

	req, err := http.NewRequest(
		"GET",
		target,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", `application/json"`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	fmt.Println(byteArray)
	res := new(OpenWeatherMapAPIResponse)
	if err != nil {
		fmt.Println("Error")
	}
	err = json.Unmarshal(byteArray, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return c.NoContent(200)
}

func ProposeCloth(c echo.Context) error {
	pos := new(Position)
	uid := c.Request().Header.Get("User_id")

	lineId := handler.GetLineUserIdById(uid)
	fmt.Println(lineId)

	if err := c.Bind(pos); err != nil {
		fmt.Println(err)
	}
	if pos.Lat == "" {
		fmt.Println("test")
	}

	res := GetTemp(pos.Lat, pos.Lon)

	var season []string

	if res.Main.Temp <= 10 {
		season = append(season, "冬")
		// season = "冬"
	} else if res.Main.Temp >= 20 {
		season = append(season, "夏")
		// season = "夏"
	} else {
		season = append(season, "春")
		// season = "春"
		season = append(season, "秋")
	}

	cloth := handler.GetProposeCloth(season, uid)
	handler.PushMessage(cloth, lineId)

	fmt.Println(cloth)
	return c.NoContent(200)
	// res := GetTemp(lat, lon)

}

func GetTemp(lat, lon string) *OpenWeatherMapAPIResponse {
	token := "6beffb309ba1da1d671c15746524791d"                   // APIトークン
	endPoint := "https://api.openweathermap.org/data/2.5/weather" // APIのエンドポイント

	// パラメータを設定
	values := url.Values{}
	values.Set("lat", lat)
	values.Set("lon", lon)
	values.Set("APPID", token)
	values.Set("units", "metric")

	target := endPoint + "?" + values.Encode()

	req, err := http.NewRequest(
		"GET",
		target,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", `application/json"`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	res := new(OpenWeatherMapAPIResponse)
	if err != nil {
		fmt.Println("Error")
	}
	err = json.Unmarshal(byteArray, &res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	return res
}
