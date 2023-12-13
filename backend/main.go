package main

import (
	"log"

	"github.com/joho/godotenv"
	"mehrang.ir/school/pkg/bootstrap"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}
	bootstrap.Serve()
}
