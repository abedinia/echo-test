package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"time"
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"encoding/json"
	"bytes"
	"github.com/labstack/echo/middleware"
)

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


func youtube(c echo.Context) error {
	name := c.QueryParam("name")
	cat := c.QueryParam("cat")

	return c.String(http.StatusOK, fmt.Sprintf("your name is %s\n and your type is %s", name, cat))
}


func youtubeData(c echo.Context) error {
	dataType := c.Param("dataType")
	fmt.Println(dataType)

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("your data is omade injaa :))))"))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string] string {
			"name" : "aaa",
		})
	}
	return c.JSON(http.StatusBadRequest, map[string] string {
		"error" : "no data type available in query string",
	})
}
type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func adding(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed to load data of body : %", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("failed to UnMarshal : %", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Print("This is your Cat %#d", cat)
	return c.String(http.StatusOK, "we got your cat")
}

func Newadd(c echo.Context) error {
	cat := Cat{}
	defer c.Request().Body.Close()

	buff := bytes.Buffer{}
	buff.ReadFrom(c.Request().Body)

	if err := json.Unmarshal(buff.Bytes(), &cat); err != nil {
		log.Printf("failed to UnMarshal : %", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is your Cat %v", cat)
	return c.String(http.StatusOK, "we got your cat")
}

func interfaciiing(c echo.Context) error {
	var i Myinterface

	i = myImplementation{}
	i.goingnow()
	i.asshole("asdasd", "asdajkshd")


	var g Myinterface

	g = NewMyinterface()
	g.goingnow()
	g.asshole("asdasd", "asdajkshd")


	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", homePage)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.GET("/channel", channeling)
	e.GET("/youtube", youtube)
	e.GET("/youtube/:dataType", youtubeData)
	e.POST("/youtube/add", adding)
	e.POST("/youtube/newadd", Newadd)
	e.GET("/interface", interfaciiing)


	e.Logger.Fatal(e.Start(":1323"))
}

type Myinterface interface{
	goingnow()
	asshole(string, string)
}

func NewMyinterface () Myinterface {
	return myImplementation{}
}

type myImplementation struct {
Aydin string
}

func (as myImplementation) goingnow() {
	fmt.Printf("salam")
}

func (as myImplementation) asshole(an, on string) {
	fmt.Printf("khobi ?")
}