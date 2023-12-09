package user_controller

import (
	"net/http"

	user_repositories "mehrang.ir/school/internal/modules/users/repositories"
	"mehrang.ir/school/utils"

	_ "github.com/mattn/go-sqlite3"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := user_repositories.GetUsers(1, 10)
	if err != nil {
		utils.ResponseToError(w, err, 400)
	}
	var statusCode = 200
	if users == nil {
		statusCode = 204
	}
	utils.WriteJson(w, statusCode, users)
}
