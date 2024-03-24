package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dayachettri/hotel-reservation/api"
	"github.com/dayachettri/hotel-reservation/db"
	"github.com/dayachettri/hotel-reservation/util"
	"github.com/labstack/echo/v4"
)

func main() {
	util.RequiredEnvVars()
	postgresStore := db.NewPostgresStore()
	err := postgresStore.Connect("DATABASE_URL")

	if err != nil {
		log.Fatal(err)
	}
	defer postgresStore.DB.Close()

	listenAddr := flag.String("listenAddr", ":1323", "The listen address of the API server")
	flag.Parse()

	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler

	apiv1 := e.Group("api/v1")

	userHandler := api.NewUserHandler(db.NewPostgresUserStore(postgresStore.DB))
	apiv1.GET("/user", userHandler.HandleGetUsers)
	apiv1.GET("/user/:id", userHandler.HandleGetUser)

	e.Logger.Fatal(e.Start(*listenAddr))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.JSON(code, map[string]string{"error": err.Error()})
}
