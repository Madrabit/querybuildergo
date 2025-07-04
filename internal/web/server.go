package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Server struct {
	R *chi.Mux
}

//goland:noinspection HttpUrlsUsage,HttpUrlsUsage
func registerMiddleware(r *chi.Mux) {
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
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
}

func NewServer() *Server {
	r := chi.NewRouter()
	registerMiddleware(r)
	return &Server{
		R: r,
	}
}
