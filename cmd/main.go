package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/fresanov/hello-api/handlers"
	"github.com/fresanov/hello-api/handlers/rest"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", rest.TranslateHandler).Methods("GET")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	adapter := gorillamux.NewV2(router)

	lambda.Start(adapter.ProxyWithContext)
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
