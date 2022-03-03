package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sub-rat/MyNewContactbook/internals/server"
)

func main() {
	fmt.Println("Starting ContactBook Api")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("DB_NAME"))
	srv := server.GetServer()
	srv.Run()
}
