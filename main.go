package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"time"
)

func main() {
	e := echo.New()
	e.GET("/", homePage)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.GET("/channel", channeling)
	e.Logger.Fatal(e.Start(":1323"))
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello this is my first project with go and echo framework")
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	go func() {
		time.Sleep(time.Second*3)
		fmt.Println("go routinggg")
	}()

	team := c.QueryParam("team")
	member := c.QueryParam("member")
	defer fmt.Print("eshak manam")

	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}


func channeling(c echo.Context) error {
	ch := make(chan int, 4)
	alaki(ch)

	for item := range ch {
		time.Sleep(time.Second)
		fmt.Print(item)
	}

	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}


func alaki (channling chan int) {
	time.Sleep(time.Second*5)
	go func() {
		for i:=0; i<10 ;i++  {
			channling <- i
		}
		close(channling)
	}()
}