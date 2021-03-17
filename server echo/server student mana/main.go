package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	IsMale bool   `json:"IsMale"`
}

var STUDENTS []Student

func main() {
	STUDENTS = append(STUDENTS, Student{
		ID:     "2162398",
		Name:   "Trần Huỳnh Đức",
		IsMale: true,
	})
	STUDENTS = append(STUDENTS, Student{
		ID:     "2171289",
		Name:   "Nguyễn Hoàng Phong Phú",
		IsMale: true,
	})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running....")
	})
	//e.POST("/users", saveUser)
	e.GET("/user", getUser)
	e.GET("/users", getUserAll)
	//e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))

}

// e.GET("/users", getUserAll)
func getUserAll(c echo.Context) error {
	jData, _ := json.Marshal(STUDENTS)
	return c.String(http.StatusOK, string(jData))
}

// e.GET("/user/id", getUser)
func getUser(c echo.Context) error {
	id := c.QueryParam("id")

	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
