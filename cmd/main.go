package main

import (
	db "cms-user/database"
	"cms-user/internal/delivery/api"
	"cms-user/internal/repository/postgres"
	"cms-user/internal/usecase"
	"cms-user/routes"
	"fmt"
)

func main() {
	db := db.Init()

	repository := postgres.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	delivery := api.NewHandler(usecase)
	echo := routes.GetRoutes(delivery)

	host := fmt.Sprintf("%s:%s", "127.0.0.1", "8800")
	_ = echo.Start(host)
}
