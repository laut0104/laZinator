package handler

import (
	// "crypto/rand"
	"math/rand"
	"time"

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

	uid := c.Request().Header.Get("User_id")

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

func GetCloth(c echo.Context) error {
	db := OpenDB()
	defer db.Close()
	id := c.Param("id")
	uid := c.Request().Header.Get("User_id")
	cloth := new(Cloth)
	err := db.QueryRow(`SELECT * FROM clothes where id=$1 and userid=$2 `, id, uid).Scan(&cloth.Id, &cloth.Userid, &cloth.Cloth, &cloth.Details, &cloth.Weather, &cloth.Temperature, &cloth.Events)
	if err != nil {
		fmt.Println("db_err")
		fmt.Println(err)
	}

	fmt.Println(cloth)
	return c.JSON(http.StatusOK, cloth)
}

func GetProposeCloth(season []string, uid string) Cloth {
	db := OpenDB()
	defer db.Close()

	fmt.Println(season)
	// AND temperature = ANY($2::text[])
	clothes := Clothes{}
	cloth := new(Cloth)

	if len(season) == 1 {
		rows, err := db.Query(`SELECT * FROM clothes where userid=$1 and $2=any(temperature)`, uid, season[0])
		if err != nil {
			fmt.Println("db_err")
			fmt.Println(err)
		}
		defer rows.Close()
		if !rows.Next() {
			rows, err = db.Query(`SELECT * FROM clothes where userid=$1`, uid)
			if err != nil {
				fmt.Println("db_err")
				fmt.Println(err)
			}
		}
		for rows.Next() {
			err := rows.Scan(&cloth.Id, &cloth.Userid, &cloth.Cloth, &cloth.Details, &cloth.Weather, &cloth.Temperature, &cloth.Events)
			if err != nil {
				fmt.Println(err)
				// return err
			}
			clothes.Clothes = append(clothes.Clothes, *cloth)
		}
	} else {
		rows, err := db.Query(`SELECT * FROM clothes where userid=$1 and ($2 = any(temperature) or $3 = any(temperature))`, uid, season[0], season[1])
		if err != nil {
			fmt.Println("db_err")
			fmt.Println(err)
		}
		defer rows.Close()
		if !rows.Next() {
			rows, err = db.Query(`SELECT * FROM clothes where userid=$1`, uid)
			if err != nil {
				fmt.Println("db_err")
				fmt.Println(err)
			}
		}
		for rows.Next() {
			err := rows.Scan(&cloth.Id, &cloth.Userid, &cloth.Cloth, &cloth.Details, &cloth.Weather, &cloth.Temperature, &cloth.Events)
			if err != nil {
				fmt.Println(err)
				// return err
			}
			clothes.Clothes = append(clothes.Clothes, *cloth)
		}
	}

	rand.Seed(time.Now().UnixNano())
	length := len(clothes.Clothes)
	randNum := rand.Intn(length)

	return clothes.Clothes[randNum]
}

func AddCloth(c echo.Context) error {
	cloth := new(Cloth)
	if err := c.Bind(cloth); err != nil {
		return err
	}
	db := OpenDB()
	defer db.Close()

	_, err := db.Exec(`INSERT INTO clothes (userid, cloth, details, weather, temperature, events) VALUES($1, $2, $3, $4, $5, $6)`, cloth.Userid, cloth.Cloth, cloth.Details, cloth.Weather, cloth.Temperature, cloth.Events)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, cloth)
}

func UpdateCloth(c echo.Context) error {
	cloth := new(Cloth)
	id := c.Param("id")
	if err := c.Bind(cloth); err != nil {
		return err
	}
	db := OpenDB()
	defer db.Close()

	_, err := db.Exec(`UPDATE clothes SET (cloth, details, weather, temperature, events)=($1, $2, $3, $4, $5) WHERE id=$6 AND userid=$7`, cloth.Cloth, cloth.Details, cloth.Weather, cloth.Temperature, cloth.Events, id, cloth.Userid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(cloth)

	return c.JSON(http.StatusOK, cloth)
}

func DeleteCloth(c echo.Context) error {
	id := c.Param("id")
	uid := c.QueryParam("id")

	db := OpenDB()
	defer db.Close()
	_, err := db.Exec(`DELETE FROM clothes WHERE id=$1 AND userid=$2`, id, uid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.NoContent(http.StatusOK)
}
