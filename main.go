package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

// Employee struct
type Employee struct {
	Name string `json:"name"`
}

var ctx = context.Background()

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
func getName(c echo.Context) error {
	clients := redis.NewClient(&redis.Options{
		Addr:     "redis707.oaf5fp.clustercfg.euc1.cache.amazonaws.com:6379",
		Password: "",
		DB:       0,
	})
	name := c.QueryParam("name")

	err := clients.Set(ctx, "name", name, 1).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := clients.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name:", val)
	return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", val))
}

func main() {
	var port string
	flag.StringVar(&port, "port", ":8080", "asdasd")
	flag.Parse()
	fmt.Println("Hello world!")
	e := echo.New()
	e.GET("/", hello)
	e.GET("/api/hello", getName)
	e.GET("/api/health", health)
	e.Logger.Fatal(e.Start(port))
}
