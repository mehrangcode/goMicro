package bootstrap

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "4000"),
		Handler: registerRoutes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
