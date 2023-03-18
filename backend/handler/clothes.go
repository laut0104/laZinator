package handler

import (
	_ "github.com/lib/pq"

	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Cloth struct {
	Id          int    `json:"id" param:"id"`
	Userid      int    `json:"userid" param:"uid"`
	Cloth       string `json:"cloth" param:"cloth"`
	Details     string `json:"details" param:"details"`
	Weather     string `json:"weather" param:"weather"`
	Temperature string `json:"temperature" param:"temperature"`
	Events      string `json:"events" param:"events"`
}

type Clothes struct {
	Clothes []Cloth `json:"clothes"`
}

func GetClothes(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	uid := c.QueryParam("id")

	rows, err := db.Query(`SELECT * FROM clothes where userid=$1`, uid)
	if err != nil {
		fmt.Println("db_err")
		fmt.Println(err)
	}
	defer rows.Close()

	clothes := Clothes{}
	cloth := new(Cloth)

	for rows.Next() {
		err := rows.Scan(&cloth.Id, &cloth.Userid, &cloth.Cloth, &cloth.Details, &cloth.Weather, &cloth.Temperature, &cloth.Events)
		if err != nil {
			fmt.Println(err)
			return err
		}
		clothes.Clothes = append(clothes.Clothes, *cloth)
	}
	fmt.Println(clothes)
	return c.JSON(http.StatusOK, clothes)
}

func AddCloth(c echo.Context) error {
	cloth := new(Cloth)
	if err := c.Bind(cloth); err != nil {
		return err
	}
	db := OpenDB()
	defer db.Close()

	_, err := db.Exec(`INSERT INTO clothes (userid, cloth, details, weather, temperature, events) VALUES($1, $2, $3, $4, $5, $6)`, cloth.Userid, cloth.Cloth, &cloth.Details, &cloth.Weather, &cloth.Temperature, &cloth.Events)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, cloth)
}
