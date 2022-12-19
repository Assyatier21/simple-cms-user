package main

import (
	db "cms/database"
	"cms/internal/delivery/api"
	"cms/internal/repository/postgres"
	"cms/routes"
	"fmt"
)

func main() {
	db := db.Init()

	repository := postgres.New(db)
	handler := api.New(repository)
	echo := routes.GetRoutes(handler)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
