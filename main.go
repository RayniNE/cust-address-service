package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/raynine/customeraddresses/server"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	log.Println("-- Creating Server -- ")
	s := server.CreateServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(fmt.Sprintf(":%v", port), s)
}
