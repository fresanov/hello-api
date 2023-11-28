package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fresanov/hello-api/handlers"
	"github.com/fresanov/hello-api/handlers/rest"
	"github.com/fresanov/hello-api/translation"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	router.HandleFunc("/hello", translateHandler.TranslateHandler).Methods("GET")
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
