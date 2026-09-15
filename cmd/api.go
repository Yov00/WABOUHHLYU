package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Yov00/ecomgov/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	// logger
	// db driver
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	conStr string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)                              // good for rate limiting
	r.Use(middleware.ClientIPFromHeader("CF-Connecting-IP")) // again for rate limiting and analytics/tracing
	// r.Use(middleware.ClientIPFromHeader("X-Real-IP")) // again for rate limiting and analytics/tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good!"))
	})

	productService := products.NewService()
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)

	// http.ListenAndServe(app.config.addr)
	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started at addr %s", app.config.addr)
	return srv.ListenAndServe()
}
