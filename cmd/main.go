package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fresanov/hello-api/handlers"
	"github.com/fresanov/hello-api/handlers/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", rest.TranslateHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	//adapter := gorillamux.NewV2(router)
	//lambda.Start(adapter.ProxyWithContext)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
