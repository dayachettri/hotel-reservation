package main

import (
	"flag"

	"github.com/dayachettri/hotel-reservation/api"
	"github.com/labstack/echo/v4"
)

func main() {
	listenAddr := flag.String("listenAddr", ":1323", "The listen address of the API server")
	flag.Parse()

	e := echo.New()

	apiv1 := e.Group("api/v1")
	apiv1.GET("/user", api.HandleGetUsers)
	apiv1.GET("/user/:id", api.HandleGetUser)

	e.Logger.Fatal(e.Start(*listenAddr))
}
