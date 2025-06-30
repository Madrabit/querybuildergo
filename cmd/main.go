package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	"querybuilder/internal/employee"
	"querybuilder/internal/manager"
	"querybuilder/internal/product"
	"querybuilder/internal/web"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cnf, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}
	db, err := database.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	server := web.NewServer()
	repProduct := product.NewRepository(db)
	svcProd := product.NewService(repProduct)
	controllerProduct := product.NewController(server, svcProd)
	controllerProduct.RegisterRoutes()
	repoManager := manager.NewRepository(db)
	serviceManager := manager.NewService(repoManager)
	controllerManager := manager.NewController(server, serviceManager)
	controllerManager.RegisterRoutes()
	repoEmployee := employee.NewRepository(db)
	generator := employee.NewGenerator()
	serviceEmployee := employee.NewService(repoEmployee, generator)
	controllerEmployee := employee.NewController(server, serviceEmployee)
	controllerEmployee.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", server.R))
}
