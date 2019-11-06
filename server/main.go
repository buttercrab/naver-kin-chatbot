package main

import (
	"github.com/buttercrab/naver-kin-chatbot/server/httpHandler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/api", httpHandler.ExampleHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
