package user_controller

import (
	"errors"
	"net/http"

	user_repositories "mehrang.ir/school/internal/modules/users/repositories"
	"mehrang.ir/school/utils"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

func List(w http.ResponseWriter, r *http.Request) {
	users, err := user_repositories.GetUsers(1, 10)
	if err != nil {
		utils.ResponseToError(w, err, 400)
		return
	}
	var statusCode = 200
	if users == nil {
		statusCode = 204
	}
	utils.WriteJson(w, statusCode, users)
}
func GetById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	users, err := user_repositories.GetUserById(userId)
	if err != nil {
		utils.ResponseToError(w, err, 400)
		return
	}
	var statusCode = 200
	if users.Id == "" {
		statusCode = 204
	}
	utils.WriteJson(w, statusCode, users)
}

func Create(w http.ResponseWriter, r *http.Request) {
	userId, err := user_repositories.CreateUser(w, r)
	if err != nil {
		utils.ResponseToError(w, err, 400)
		return
	}
	var statusCode = http.StatusCreated
	if userId == nil {
		utils.ResponseToError(w, errors.New("cant inser user"), 400)
		return
	}
	utils.WriteJson(w, statusCode, struct{ Id *int }{Id: userId})
}
func Update(w http.ResponseWriter, r *http.Request) {
	err := user_repositories.UpdateUser(w, r)
	if err != nil {
		utils.ResponseToError(w, err, 400)
		return
	}
	var statusCode = http.StatusOK
	utils.WriteJson(w, statusCode, struct{ Message string }{Message: "User Updated"})
}
func Delete(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userId")
	err := user_repositories.DeleteUser(userID)
	if err != nil {
		utils.ResponseToError(w, err, 400)
		return
	}
	var statusCode = http.StatusOK
	utils.WriteJson(w, statusCode, struct{ Message string }{Message: "User Deleted"})
}
