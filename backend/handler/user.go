package handler

import (
	"fmt"

	_ "github.com/lib/pq"

	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id         int    `json:"id"`
	Lineuserid string `json:"lineuserid"`
	Username   string `json:"username"`
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	db := OpenDB()
	defer db.Close()

	u := new(User)
	err := db.QueryRow(`SELECT * FROM users where id=$1`, id).Scan(&u.Id, &u.Lineuserid, &u.Username)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)

	defer db.Close()
	return c.JSON(http.StatusOK, u)

}

func GetUser(c echo.Context) error {
	fmt.Println(c)
	lineid := c.QueryParam("lineuserid")
	fmt.Println(lineid)
	db := OpenDB()
	defer db.Close()

	u := new(User)
	err := db.QueryRow(`SELECT * FROM users where lineuserid=$1`, lineid).Scan(&u.Id, &u.Lineuserid, &u.Username)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)

	defer db.Close()
	return c.JSON(http.StatusOK, u)

}
