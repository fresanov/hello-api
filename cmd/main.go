package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/fresanov/hello-api/handlers/rest"
	"github.com/gorilla/mux"
)

func main() {

	/* addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) */

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found", r.RequestURI)
		http.Error(w, fmt.Sprintf("Not found: %s", r.RequestURI), http.StatusNotFound)
	})
	router.HandleFunc("/hello", rest.TranslateHandler).Methods("GET")
	adapter := gorillamux.NewV2(router)

	lambda.Start(adapter.ProxyWithContext)
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
