package bootstrap

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	coursesModule "mehrang.ir/school/internal/modules"
	user_controller "mehrang.ir/school/internal/modules/users/controller"
	"mehrang.ir/school/utils"
)

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := utils.WriteJson(w, http.StatusOK, struct {
			Message string
			Code    int
		}{
			Code:    http.StatusOK,
			Message: "App is running",
		})
		if err != nil {
			log.Fatalf("Error is accured %v", err)
		}
	})
	r.Get("/courses", coursesModule.List)
	r.Get("/users", user_controller.List)
	return r
}
