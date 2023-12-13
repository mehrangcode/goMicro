package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mehrang.ir/school/pkg/database"
)

func Serve() {
	err := database.InitDB()
	if err != nil {
		log.Panic(err)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: registerRoutes(),
	}

	fmt.Println("APP RUNING ON : http://localhost:" + os.Getenv("PORT"))
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
