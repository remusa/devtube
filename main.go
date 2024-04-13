package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	// "github.com/remusa/devtube/internal/database"
)

// type apiConfig struct {
// 	DB *database.Queries
// }

func main() {
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not found in environment")
	}

	DATABASE_URL := os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		log.Fatal("DATABASE_URL not found in environment")
	}
	log.Printf("Connected to database: %v", DATABASE_URL)

	_, err := sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
	// dbQueries := database.New(db)

	// apiCfg := apiConfig{
	// 	DB: dbQueries,
	// }

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
