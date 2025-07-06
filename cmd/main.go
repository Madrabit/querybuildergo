package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	gracefulshutdown "github.com/quii/go-graceful-shutdown"
	"log"
	"net/http"
	"querybuilder/internal/common"
	"querybuilder/internal/database"
	"querybuilder/internal/employee"
	"querybuilder/internal/manager"
	"querybuilder/internal/product"
	"querybuilder/internal/validator"
	"querybuilder/internal/web"
)

// @title Query Builder API
// @version 1.0
// @description App for building queries

// @host localhost:8080
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cnf, err := common.Load()
	if err != nil {
		fmt.Println(err)
	}
	logger := common.NewLogger(cnf)
	defer func() { _ = logger.Sync() }()
	db, err := database.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	server := build(db, logger)
	httpServer := &http.Server{Addr: ":8080", Handler: server.R}
	ctx := context.Background()
	srv := gracefulshutdown.NewServer(httpServer)
	if err := srv.ListenAndServe(ctx); err != nil {
		logger.Fatal("didnt shutdown gracefully, some responses may have been lost")
	}
	logger.Info("shutdown gracefully! all responses were sent")
}

func build(db *sqlx.DB, logger *common.Logger) *web.Server {
	server := web.NewServer()
	vld := validator.New()
	repProduct := product.NewRepository(db)
	svcProd := product.NewService(repProduct, vld)
	controllerProduct := product.NewController(server, svcProd, logger)
	controllerProduct.RegisterRoutes()
	repoManager := manager.NewRepository(db)
	serviceManager := manager.NewService(repoManager, vld)
	controllerManager := manager.NewController(server, serviceManager, logger)
	controllerManager.RegisterRoutes()
	repoEmployee := employee.NewRepository(db)
	generator := employee.NewGenerator()
	serviceEmployee := employee.NewService(repoEmployee, generator, vld)
	controllerEmployee := employee.NewController(server, serviceEmployee, logger)
	controllerEmployee.RegisterRoutes()
	return server
}
