package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/remusa/devtube/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not found in environment")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL not found in environment")
	}
	log.Printf("Connected to database: %v", DB_URL)

	db, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/ready", handlerReadiness)
	mux.HandleFunc("GET /v1/err", handlerErr)

	corsMux := corsMiddleware(mux)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: corsMux,
	}

	log.Printf("Server started on port %v\n", PORT)
	log.Fatal(server.ListenAndServe())
}
