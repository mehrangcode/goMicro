package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mehrangcode.ir/todoapp/cmd/initializers"
)

type Config struct{}

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
}
func main() {
	var PORT = os.Getenv("PORT")
	app := Config{}
	fmt.Println("PORT: ", PORT)
	log.Printf("\nApplication is up on ... %v", PORT)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
