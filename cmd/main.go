package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"querybuilder/internal/config"
	"querybuilder/internal/storage"
	"querybuilder/services/employee"
	"querybuilder/services/manager"
	"querybuilder/services/product"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	cnf, err := config.Load()
	if err != nil {
		fmt.Println(err)
	}

	db, err := storage.NewMssqlStorage(cnf.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	storeProd := product.NewStore(db)
	handlerProd := product.NewHandler(storeProd)

	storeManager := manager.NewStore(db)
	handlerManager := manager.NewHandler(storeManager)

	storeEmp := employee.NewStore(db)
	handlerEmp := employee.NewHandler(storeEmp)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
			"http://192.168.102.217:4200",
			"http://192.168.102.217:4201",
			"http://qb.ibdarb.ru",
			"http://members.ibdarb.ru",
			"https://qb.ibdarb.ru",
			"https://members.ibdarb.ru",
			"https://core.ibdarb.ru",
			"https://192.168.102.217:4200",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           3600,
	}))
	r.Route("/products", func(r chi.Router) {
		r.Mount("/", product.Routes(handlerProd))
	})
	r.Route("/manager", func(r chi.Router) {
		r.Mount("/", manager.Routes(handlerManager))
	})
	r.Route("/employee", func(r chi.Router) {
		r.Mount("/", employee.Routes(handlerEmp))
	})
	log.Fatal(http.ListenAndServe(":8080", r))

}
