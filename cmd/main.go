package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"querybuilder/internal/config"
	"querybuilder/internal/database"
	employee2 "querybuilder/internal/employee"
	manager2 "querybuilder/internal/manager"
	product2 "querybuilder/internal/product"
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
	storeProd := product2.NewStore(db)
	handlerProd := product2.NewHandler(storeProd)

	storeManager := manager2.NewStore(db)
	handlerManager := manager2.NewHandler(storeManager)

	storeEmp := employee2.NewStore(db)
	handlerEmp := employee2.NewHandler(storeEmp)

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
		r.Mount("/", product2.Routes(handlerProd))
	})
	r.Route("/manager", func(r chi.Router) {
		r.Mount("/", manager2.Routes(handlerManager))
	})
	r.Route("/employee", func(r chi.Router) {
		r.Mount("/", employee2.Routes(handlerEmp))
	})
	log.Fatal(http.ListenAndServe(":8080", r))

}
