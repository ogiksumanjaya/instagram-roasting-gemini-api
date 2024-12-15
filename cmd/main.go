package main

import (
	"context"
	"fmt"
	config "instagram-roasting"
	"instagram-roasting/core/module"
	"instagram-roasting/handler"
	"instagram-roasting/libs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	var (
		err error
		cfg *config.Config
	)
	if _, ok := os.LookupEnv("SERVER_REST_PORT"); !ok {
		cfg.Setup(".env")
	}

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	r := chi.NewRouter()

	// register repository
	scrappingRepo := libs.NewScrapingIGProfile()
	geminiAIRepo := libs.NewGeminiAI(cfg)

	// register use case
	roastingProfileIgUC := module.NewRoastingUC(scrappingRepo, geminiAIRepo)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handler.Healthness)
	r.Route("/", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			handler.NewRoastingProfileHandler(r, roastingProfileIgUC)
		})
	})

	srv := http.Server{
		Addr:    cfg.GetServerAddress(),
		Handler: r,
	}

	go func() {
		fmt.Println("starting server at", srv.Addr)

		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			fmt.Println("failed to start server", err)
		}
	}()

	<-signalChan

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err = srv.Shutdown(ctxWithTimeout)

	if err != nil {
		fmt.Println("failed to shutdown server", err)
	}

	fmt.Println("shutting down")
}
