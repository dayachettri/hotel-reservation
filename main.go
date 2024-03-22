package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dayachettri/hotel-reservation/api"
	"github.com/dayachettri/hotel-reservation/db"
	"github.com/dayachettri/hotel-reservation/util"
	"github.com/labstack/echo/v4"
)

func main() {
	util.RequiredEnvVars()
	postgresStore := db.NewPostgresStore()
	db, err := postgresStore.Connect("DATABASE_URL")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listenAddr := flag.String("listenAddr", ":1323", "The listen address of the API server")
	flag.Parse()

	e := echo.New()

	apiv1 := e.Group("api/v1")
	apiv1.GET("/user", api.HandleGetUsers)
	apiv1.GET("/user/:id", api.HandleGetUser)

	e.Logger.Fatal(e.Start(*listenAddr))
	fmt.Println("exited main")
}
