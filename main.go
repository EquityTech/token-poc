package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"github.com/ssb4/token-poc/api"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{"*"},
	})

	router := api.NewRouter()
	port := os.Getenv("PORT")
	// For local dev
	if port == "" {
		port = "4040"
	}

	handler := c.Handler(router)

	fmt.Println("Server is running.")
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
