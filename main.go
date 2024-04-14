package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	"github.com/remusa/devtube/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("root")
	})

	mux.HandleFunc("POST v1/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("GET v1/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	mux.HandleFunc("POST v1/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedsCreate))
	mux.HandleFunc("GET v1/feeds", apiCfg.handlerFeedsGet)

	mux.HandleFunc("GET v1/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	mux.HandleFunc("GET v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	mux.HandleFunc("POST v1/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsCreate))
	mux.HandleFunc("DELETE v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsDelete))

	mux.HandleFunc("GET v1/ready", handlerReadiness)
	mux.HandleFunc("GET v1/err", handlerErr)

	corsMux := corsMiddleware(mux)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	const collectionConcurrency = 10
	const collectionInterval = time.Minute
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	v := reflect.ValueOf(http.DefaultServeMux).Elem()
	fmt.Printf("routes: %v\n", v.FieldByName("mux121").FieldByName("m"))
	log.Printf("Serving on port: %s\n", port)
	err = server.ListenAndServe()
	log.Fatal(err)
}
