package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/fresanov/hello-api/handlers/rest"
)

var httpLambda *httpadapter.HandlerAdapter

func init() {
	http.HandleFunc("/hello", rest.TranslateHandler)

	httpLambda = httpadapter.New(http.DefaultServeMux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return httpLambda.ProxyWithContext(ctx, req)
}

func main() {

	/* addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux)) */

	lambda.Start(Handler)
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
