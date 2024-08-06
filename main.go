package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mondal-Prasun/rssAgrigateServer/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")

	portNum := os.Getenv("PORT")

	if portNum == "" {
		log.Fatal("PORT is empty")
	}
	fmt.Println("PORT:", portNum)

	dbPortNum := os.Getenv("DB_URL")

	if dbPortNum == "" {
		log.Fatal("DB url is empty")
	}
	log.Println("DB_URL:", dbPortNum)

	dbConnection, err := sql.Open("postgres", dbPortNum)

	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	apiCfg := apiConfig{
		db: database.New(dbConnection),
	}

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
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/getUser", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	route.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: route,
		Addr:    ":" + portNum,
	}

	log.Println("Server is running on:", portNum)
	log.Println("Database is running on:", dbPortNum)
	erroR := srv.ListenAndServe()

	if erroR != nil {
		log.Fatal(err)
	}

}
