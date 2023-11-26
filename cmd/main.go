package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"), "error")
	if addr == ":" {
		addr = ":8080"
	}

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
