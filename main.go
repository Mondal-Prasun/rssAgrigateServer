package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	portNum := os.Getenv("PORT")

	if portNum == "" {
		log.Fatal("PORT is empty")
	}

	fmt.Println("PORT:", portNum)

	route := chi.NewRouter()

	route.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()

	v1Router.Get("/ready", handlerReadiness)
	v1Router.Get("/err", handlerError)

	route.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: route,
		Addr:    ":" + portNum,
	}

	log.Println("Server is running on: ", portNum)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
